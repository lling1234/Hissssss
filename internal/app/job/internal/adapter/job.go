package adapter

import (
	"context"
	"github.com/cd-home/Hissssss/api/pb/chat"
	"github.com/cd-home/Hissssss/internal/app/job/internal/model"
)

// ChatMQ 聊天专属MQ
type ChatMQ interface {
	Consume(ctx context.Context, ack chan struct{}, nack chan struct{}) (<-chan *chat.MessageToMQ, chan struct{})
	ConsumeRoom()
	ConsumeBroadcast()
}

type JobCache interface {
	// GetServerByUID 缓存获取用户映射服务器
	GetServerByUID(uid string) (string, error)
}

type JobRepo interface {
	// CreateOfflineMessage 离线消息存储
	CreateOfflineMessage(msg *model.OfflineMessage) error
}
