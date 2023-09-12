package connect

import (
	"context"
	"github.com/cd-home/Hissssss/api/pb/connect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync/atomic"
)

type Connect struct {
	ID             string
	ctx            context.Context
	cancel         context.CancelFunc
	client         connect.PushMessageClient
	PushCh         []chan *connect.Message // 单聊
	RoomCh         []chan *connect.Message // 群聊
	BroadcastCh    []chan *connect.Message // 广播
	Routines       uint64
	PushIndex      uint64
	RoomIndex      uint64
	BroadcastIndex uint64
}

func New(id string) *Connect {
	c := &Connect{}
	// TODO 配置化
	c.Routines = 4
	c.ID = id
	c.ctx, c.cancel = context.WithCancel(context.Background())
	// TODO 加快处理流程
	var i uint64
	c.PushCh = make([]chan *connect.Message, c.Routines)
	c.RoomCh = make([]chan *connect.Message, c.Routines)
	c.BroadcastCh = make([]chan *connect.Message, c.Routines)
	for i = 0; i < c.Routines; i++ {
		c.PushCh[i] = make(chan *connect.Message, 1)
		c.RoomCh[i] = make(chan *connect.Message, 1)
		c.BroadcastCh[i] = make(chan *connect.Message, 1)
		go c.Dispatch(c.PushCh[i], c.RoomCh[i], c.BroadcastCh[i])
	}
	return c
}

func (c *Connect) MakeClient(addr string) *Connect {
	conn, _ := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	client := connect.NewPushMessageClient(conn)
	c.client = client
	return c
}

func (c *Connect) Push(msg *connect.Message) {
	c.PushCh[atomic.AddUint64(&c.PushIndex, 1)%c.Routines] <- msg
}

func (c *Connect) PushRoom(msg *connect.Message) {
	c.RoomCh[atomic.AddUint64(&c.RoomIndex, 1)%c.Routines] <- msg
}

func (c *Connect) Broadcast(msg *connect.Message) {
	c.BroadcastCh[atomic.AddUint64(&c.BroadcastIndex, 1)%c.Routines] <- msg
}

func (c *Connect) Dispatch(pushCh chan *connect.Message, roomCh chan *connect.Message, broadcastCh chan *connect.Message) {
	for {
		select {
		case singleMsg := <-pushCh:
			_, err := c.client.Push(context.Background(), &connect.Message{
				MsgId:  singleMsg.MsgId,
				Server: singleMsg.Server,
				Room:   singleMsg.Room,
				From:   singleMsg.From,
				To:     singleMsg.To,
				Body:   singleMsg.Body,
				Type:   singleMsg.Type,
			})
			if err != nil {
				return
			}
		case roomMsg := <-roomCh:
			log.Println(roomMsg)
		case broadcastMsg := <-broadcastCh:
			log.Println(broadcastMsg)
		case <-c.ctx.Done():
			return
		}
	}
}
