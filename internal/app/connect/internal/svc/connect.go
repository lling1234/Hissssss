package svc

import (
	"context"
	"encoding/json"
	"github.com/cd-home/Hissssss/api/pb/common"
	"github.com/cd-home/Hissssss/api/pb/connect"
	conn "github.com/cd-home/Hissssss/internal/app/connect/internal/connect"
	"go.uber.org/zap"
	"sync"
)

type Message struct {
	connect.UnimplementedPushMessageServer
	logger  *zap.Logger
	Connect *conn.Websocket
	message *sync.Pool
}

type message struct {
	MsgId int64           `json:"msgId"`
	From  int64           `json:"from"`
	Body  string          `json:"body"`
	Op    common.OP       `json:"op"`
	Sub   common.Message  `json:"sub"`
	Type  common.PushType `json:"type"`
}

func NewMessage(logger *zap.Logger, connect *conn.Websocket) *Message {
	return &Message{
		logger:  logger.WithOptions(zap.Fields(zap.String("module", "connect service"))),
		Connect: connect,
		message: &sync.Pool{
			New: func() any {
				return &message{}
			},
		},
	}
}

func (m *Message) Push(ctx context.Context, req *connect.Message) (*connect.MessageReply, error) {
	b := m.Connect.BucketX(req.To)
	to := m.message.Get().(*message)
	to.From = req.From
	to.Body = req.Body
	to.Op = req.Op
	to.MsgId = req.MsgId
	to.Sub = req.Sub
	to.Type = req.Type
	data, _ := json.Marshal(to)
	switch req.Type {
	case common.PushType_Single:
		c := b.ChannelX(req.To)
		c.SendM(data)
	case common.PushType_Room:
		r := b.RoomX(req.To)
		r.Broadcast(data)
	case common.PushType_Broadcast:
		b.BroadcastToAllChannel(data)
	}
	to = &message{}
	m.message.Put(to)
	return &connect.MessageReply{}, nil
}
