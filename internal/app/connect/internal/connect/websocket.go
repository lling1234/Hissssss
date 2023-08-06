package connect

import (
	"bytes"
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/cd-home/Hissssss/api/pb/account"
	"github.com/cd-home/Hissssss/internal/app/connect/config"
	uuid "github.com/satori/go.uuid"
	"github.com/zentures/cityhash"
	"go.uber.org/zap"
	"io"
	"net/http"
	"nhooyr.io/websocket"
	"strconv"
)

type Websocket struct {
	ID      string
	SnowID  *snowflake.Node
	config  config.Config
	Logger  *zap.Logger
	Buckets []*Bucket
	Account account.AccountClient
}

func New(config config.Config, logger *zap.Logger, account account.AccountClient) *Websocket {
	c := &Websocket{}
	c.ID = config.Spec.Node.ID
	node, _ := strconv.ParseInt(c.ID, 10, 64)
	c.SnowID, _ = snowflake.NewNode(node % 30)
	c.config = config
	c.Logger = logger.WithOptions(zap.Fields(zap.String("module", "websocket")))
	// TODO bucket size; channel room size to config
	buckets := make([]*Bucket, 4)
	for i := 0; i < 4; i++ {
		buckets[i] = NewBucket(Option{
			ChannelSize: 1000,
			RoomSize:    100,
		})
	}
	c.Buckets = buckets
	c.Account = account
	return c
}

func (w *Websocket) InitWebSocket(ctx context.Context) error {
	http.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(rw, r, &websocket.AcceptOptions{
			InsecureSkipVerify: true,
		})
		if err != nil {
			w.Logger.Error(err.Error())
			_, _ = io.Copy(rw, bytes.NewReader([]byte("upgrade webSocket error")))
			return
		}
		// TODO 临时测试采用uid, 后期采用auth_key认证获取uid
		q := r.URL.Query()
		key := q.Get("auth_key")
		uid, _ := strconv.ParseInt(key, 10, 64)
		// 加入连接管理
		ch := NewChannel(uid, conn, r.RemoteAddr, uuid.NewV4().String())
		b := w.BucketX(uid)
		b.JoinC(uid, ch)
		// 建立连接信息,主要是映射 uid:server_id;
		w.Connect(uid, w.ID)
		w.Logger.Info("[新建连接]", zap.Any("uid", ch.UID))
		go w.Read(context.Background(), ch)
		go w.Write(context.Background(), ch)
	})
	w.Logger.Info("websocket listen: " + w.config.Spec.Node.HTTP)
	return http.ListenAndServe(":"+w.config.Spec.Node.HTTP, nil)
}

// BucketX 获取用户所在的Bucket
func (w *Websocket) BucketX(uid int64) *Bucket {
	s := strconv.Itoa(int(uid))
	idx := cityhash.CityHash64([]byte(s), uint32(uint64(len(s)))) % uint64(len(w.Buckets))
	return w.Buckets[idx]
}

func (w *Websocket) Write(ctx context.Context, ch *Channel) {
	beat, stop := ch.HeartBeat()
	ch.Stop = stop
	go func() {
		for {
			select {
			case e := <-beat:
				if e != nil {
					w.BucketX(ch.UID).LeaveC(ch)
					w.DisConnect(ch.UID) // 断开连接逻辑(清除相关在线缓冲数据)
					return
				}
			}
		}
	}()
	for {
		select {
		case msg := <-ch.Send:
			//err := wsjson.Write(ctx, ch.Conn, msg) 此种方式msg是golang obj 即可
			err := ch.Conn.Write(ctx, websocket.MessageText, msg)
			if err != nil {
				w.Logger.Error(err.Error())
			}
		case <-ch.Ctx.Done():
			w.Logger.Info("[断开连接]", zap.Any("uid", ch.UID))
			close(ch.Stop)
			return
		}
	}
}

func (w *Websocket) Read(ctx context.Context, ch *Channel) {
	for i := 0; i < 2; i++ {
		go ch.Dispatch()
	}
	for {
		_, data, err := ch.Conn.Read(ctx)
		if err != nil {
			w.Logger.Warn("[数据异常]", zap.Any("uid", ch.UID))
			break
		}
		select {
		case ch.Receive <- data:
		}
	}
	ch.Cancel()
}

// Connect 连接信息
func (w *Websocket) Connect(uid int64, serverID string) {
	connectReply, err := w.Account.Connect(context.Background(), &account.ConnectRequest{
		Uid:      uid,
		ServerID: serverID,
	})
	if err != nil || connectReply.Code != 200 {
		w.Logger.Error(err.Error())
		return
	}
	return
}

// DisConnect 断开连接
func (w *Websocket) DisConnect(uid int64) {
	disConnectReply, err := w.Account.DisConnect(context.Background(), &account.DisConnectRequest{
		Uid: uid,
	})
	if err != nil || disConnectReply.Code != 200 {
		w.Logger.Error(err.Error())
		return
	}
	return
}

func (w *Websocket) Clear() {
	for i := 0; i < len(w.Buckets); i++ {
		b := w.Buckets[i]
		for uid, _ := range b.Channels {
			w.DisConnect(uid)
		}
	}
}
