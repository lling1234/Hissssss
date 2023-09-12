package job

import (
	"context"
	"github.com/cd-home/Hissssss/api/pb/chat"
	"github.com/cd-home/Hissssss/api/pb/common"
	connectx "github.com/cd-home/Hissssss/api/pb/connect"
	"github.com/cd-home/Hissssss/internal/app/job/config"
	"github.com/cd-home/Hissssss/internal/app/job/internal/adapter"
	"github.com/cd-home/Hissssss/internal/app/job/internal/biz"
	"github.com/cd-home/Hissssss/internal/app/job/internal/connect"
	"github.com/cd-home/Hissssss/internal/pkg/key"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.uber.org/zap"
	"sync"
	"time"
)

type Job struct {
	lock       sync.RWMutex
	logger     *zap.Logger
	connect    map[string]*connect.Connect // 连接层对象 server_id : client
	config     config.Config
	etcdClient *clientv3.Client
	manager    endpoints.Manager
	biz        *biz.JobBiz
	mq         adapter.ChatMQ
	cache      adapter.JobCache
}

func New(c config.Config, logger *zap.Logger, etcdClient *clientv3.Client,
	biz *biz.JobBiz, mq adapter.ChatMQ, cache adapter.JobCache) {
	j := &Job{
		logger:     logger.WithOptions(zap.Fields(zap.String("module", "job server"))),
		config:     c,
		connect:    make(map[string]*connect.Connect),
		etcdClient: etcdClient,
		biz:        biz,
		mq:         mq,
		cache:      cache,
	}
	// 监听链接层
	go j.watchConnect("connect")
	// 消费消息
	// 需要等待一下, 需要先监听到Connect层客户端
	time.Sleep(time.Second)
	go j.consumeSingle()
}

func (j *Job) watchConnect(service string) {
	etcdManager, err := endpoints.NewManager(j.etcdClient, service)
	if err != nil {
		j.logger.Error(err.Error())
		return
	}
	// map[connect/127.0.0.1:7071:{127.0.0.1:7071 3001}]
	j.manager = etcdManager
	j.fresh()
	// TODO ticker 时间 配置化
	tk := time.NewTicker(time.Second * 5)
	defer tk.Stop()
	for {
		select {
		case <-tk.C:
			j.fresh()
		}
	}
}

func (j *Job) fresh() {
	key2EndpointMap, err := j.manager.List(context.Background())
	if err != nil {
		j.logger.Error(err.Error())
		return
	}
	kem := make(map[string]string, len(key2EndpointMap))
	for _, endpoint := range key2EndpointMap {
		kem[endpoint.Metadata.(string)] = endpoint.Addr
	}
	j.lock.RLock()
	// 添加新的映射
	for k, addr := range kem {
		if _, ok := j.connect[k]; !ok {
			cc := connect.New(k)
			cc.MakeClient(addr)
			j.connect[k] = cc
		}
	}
	// 删除不存活的映射
	for k, _ := range j.connect {
		if _, ok := kem[k]; !ok {
			delete(j.connect, k)
		}
	}
	j.lock.RUnlock()
}

func (j *Job) consumeSingle() {
	j.logger.Info("consumeSingle")
	ack := make(chan struct{}, 1)
	nack := make(chan struct{}, 1)
	ch, stop := j.mq.Consume(context.Background(), ack, nack)
	for msg := range ch {
		if msg == nil {
			close(stop)
			time.Sleep(time.Second * 20) // TODO 间隔多长时间重启消费, 取决于rabbitmq服务重启的时间
			j.consumeSingle()
			return
		}
		err := j.Push(msg)
		if err != nil {
			j.logger.Error(err.Error())
			nack <- struct{}{}
		} else {
			// other biz 处理完成, 确认消息
			ack <- struct{}{}
		}
	}
}

func (j *Job) Push(msg *chat.MessageToMQ) error {
	// 选择 server_id : connect 进行消息发送
	// 获取用户所在的server_id
	msg.Server, _ = j.cache.GetServerByUID(key.GenUidMappingServer(msg.To))
	if len(msg.Server) == 0 {
		// 写入离线消息表
		return nil
	}
	c, ok := j.connect[msg.Server]
	if ok {
		push := &connectx.Message{
			MsgId:  msg.MsgId,
			Server: msg.Server,
			From:   msg.From,
			To:     msg.To,
			Body:   msg.Body,
			Type:   msg.Type,
			Op:     common.OP_NOTIFY, // 服务端通知
		}
		switch msg.Type {
		case common.PushType_Single:
			c.Push(push)
		case common.PushType_Room:
			c.PushRoom(push)
		case common.PushType_Broadcast:
			c.Broadcast(push)
		default:
			j.logger.Warn("no support this type: " + msg.Type.String())
		}
	} else {
		j.logger.Warn("not found connect service: ", zap.String("service", msg.Server))
	}
	return nil
}
