package connect

import (
	"context"
	"errors"
	"nhooyr.io/websocket"
	"time"
)

type Channel struct {
	UUID     string
	Ctx      context.Context
	Cancel   context.CancelFunc
	UID      int64
	CTime    int64
	IP       string
	PongTime int64
	Conn     *websocket.Conn
	Send     chan []byte
	Receive  chan []byte
	Next     *Channel
	Prev     *Channel
	Room     *Room
	Stop     chan struct{}
	Beat     time.Duration
}

func NewChannel(uid int64, conn *websocket.Conn, ip, uuid string) *Channel {
	ctx, cancel := context.WithCancel(context.Background())
	c := &Channel{
		UUID:    uuid,
		Ctx:     ctx,
		Cancel:  cancel,
		UID:     uid,
		IP:      ip,
		CTime:   time.Now().Unix(),
		Conn:    conn,
		Send:    make(chan []byte), // TODO 消息通道size
		Receive: make(chan []byte), // TODO 消息通道size
		Beat:    time.Second * 2,   // TODO beat时间配置化
	}
	return c
}

func (c *Channel) SendM(msg []byte) {
	c.Send <- msg
}

func (c *Channel) HeartBeat() (<-chan error, chan struct{}) {
	ch := make(chan error, 1)
	tk := time.NewTicker(c.Beat)
	stop := make(chan struct{})
	go func() {
		defer tk.Stop()
		for {
			select {
			case <-stop:
				ch <- errors.New("go away")
				return
			case t := <-tk.C:
				if err := c.Conn.Ping(context.Background()); err != nil {
					ch <- err
					return
				}
				c.PongTime = t.Unix()
			}
		}
	}()
	return ch, stop
}

func (c *Channel) Dispatch() {
	for {
		select {
		case <-c.Ctx.Done():
			return
		case _ = <-c.Receive: // 目前通过api服务投递信息
		}
	}
}
