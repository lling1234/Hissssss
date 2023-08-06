package mq

import (
	"context"
	"encoding/json"
	"github.com/cd-home/Hissssss/api/pb/chat"
	"github.com/cd-home/Hissssss/internal/app/job/config"
	"github.com/cd-home/Hissssss/internal/app/job/internal/adapter"
	"github.com/cd-home/Hissssss/internal/pkg/mq"
	"go.uber.org/zap"
)

type RabbitMQ struct {
	logger *zap.Logger
	queue  config.Queue
	mq     *mq.RabbitMQ
}

func NewRabbitMQ(queue config.Queue, logger *zap.Logger, mq *mq.RabbitMQ) adapter.ChatMQ {
	mq.SetUp(queue.Single, queue.Room, queue.Broadcast)
	return &RabbitMQ{
		queue:  queue,
		logger: logger.WithOptions(zap.Fields(zap.String("module", "mq"))),
		mq:     mq,
	}
}

func (r *RabbitMQ) Consume(ctx context.Context, ack chan struct{}, nack chan struct{}) (<-chan *chat.MessageToMQ, chan struct{}) {
	msg := make(chan *chat.MessageToMQ, 1)
	stop := make(chan struct{}, 1)
	go func() {
		ch, err := r.mq.Client[r.queue.Single].Consume()
		if err != nil {
			r.logger.Error(err.Error())
		}
		for {
			select {
			case c := <-ch:
				var m *chat.MessageToMQ
				_ = json.Unmarshal(c.Body, &m)
				msg <- m
				select {
				case <-ack:
					err := c.Ack(false)
					if err != nil {
						r.logger.Error(err.Error())
					}
				case <-nack:
					//err := c.Reject(true)      // 发送给其他的消费者 ｜ 还是直接丢弃
					err := c.Nack(false, true)
					if err != nil {
						r.logger.Error(err.Error())
					}
				}
			}
		}
	}()
	return msg, stop
}

func (r *RabbitMQ) ConsumeRoom() {

}

func (r *RabbitMQ) ConsumeBroadcast() {

}
