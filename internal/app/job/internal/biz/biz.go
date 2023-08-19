package biz

import (
	"github.com/cd-home/Hissssss/internal/app/job/internal/adapter"
	"github.com/cd-home/Hissssss/internal/pkg/config/queue"
	"go.uber.org/zap"
)

type JobBiz struct {
	logger *zap.Logger
	queue  queue.Config
	mq     adapter.ChatMQ
}

func NewJobBiz(logger *zap.Logger, queue queue.Config, mq adapter.ChatMQ) *JobBiz {
	j := &JobBiz{
		logger: logger,
		queue:  queue,
		mq:     mq,
	}
	return j
}

// ConsumeSingle 消费单聊
func (j *JobBiz) ConsumeSingle() {
	// 可能有其他的业务
}

// ConsumeRoom 消费房间
func (j *JobBiz) ConsumeRoom() {
}

// ConsumeBroadcast 消费广播
func (j *JobBiz) ConsumeBroadcast() {
}
