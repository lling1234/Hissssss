package mq

import (
	"context"
	"encoding/json"
	"github.com/cd-home/Hissssss/api/pb/chat"
	"github.com/cd-home/Hissssss/internal/app/logic/chat/internal/adapter"
	"github.com/cd-home/Hissssss/internal/pkg/config/queue"
	"github.com/cd-home/Hissssss/internal/pkg/mq"
)

type RabbitMQ struct {
	Queue queue.Config
	mq    *mq.RabbitMQ
}

func NewRabbitMQ(queue queue.Config, mq *mq.RabbitMQ) adapter.ChatMQ {
	mq.SetUp(queue.Single, queue.Room, queue.Broadcast)
	r := &RabbitMQ{mq: mq, Queue: queue}
	return r
}

func (r *RabbitMQ) Push(ctx context.Context, msg *chat.MessageToMQ) error {
	data, _ := json.Marshal(msg)
	return r.mq.Client[r.Queue.Single].Push(ctx, data)
}

func (r *RabbitMQ) PushRoom()   {}
func (r *RabbitMQ) BroadcastX() {}
