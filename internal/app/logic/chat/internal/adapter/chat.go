package adapter

import (
	"context"
	"github.com/cd-home/Hissssss/api/pb/chat"
	"github.com/cd-home/Hissssss/internal/app/logic/chat/internal/model"
)

type ChatBiz interface {
	Push(ctx context.Context, req *chat.SendMessageRequest) (int64, error)
	PushRoom() error
	BroadcastX() error
}

type ChatRepo interface {
	// CreateAllMessage 全消息存储Mongo
	CreateAllMessage(msg *model.AllMessage) error
}

type ChatCache interface {
	// GetServerByUID 缓存获取用户映射服务器
	GetServerByUID(key string) (string, error)
	// GetUserMessageID 获取单聊用户的消息序号
	GetUserMessageID(key string) (int64, error)
}

// ChatMQ 聊天专属MQ
type ChatMQ interface {
	Push(ctx context.Context, msg *chat.MessageToMQ) error
	PushRoom()
	BroadcastX()
}
