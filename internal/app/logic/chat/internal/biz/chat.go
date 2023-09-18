package biz

import (
	"context"
	"github.com/cd-home/Hissssss/api/pb/chat"
	"github.com/cd-home/Hissssss/internal/app/logic/chat/internal/adapter"
	"github.com/cd-home/Hissssss/internal/pkg/config/queue"
	"github.com/cd-home/Hissssss/internal/pkg/tool/snowid"
	"go.uber.org/zap"
)

type ChatBiz struct {
	logger *zap.Logger
	queue  queue.Config
	repo   adapter.ChatRepo
	mq     adapter.ChatMQ
	cache  adapter.ChatCache
	snow   *snowid.Snow
}

func NewChatBiz(logger *zap.Logger, queue queue.Config, repo adapter.ChatRepo, mq adapter.ChatMQ,
	cache adapter.ChatCache, snow *snowid.Snow) adapter.ChatBiz {
	return &ChatBiz{
		logger: logger.WithOptions(zap.Fields(zap.String("module", "chat biz"))),
		queue:  queue,
		repo:   repo,
		mq:     mq,
		cache:  cache,
		snow:   snow,
	}
}

func (cb *ChatBiz) Push(ctx context.Context, req *chat.SendMessageRequest) (msgId int64, err error) {
	// 获取用户双向有序消息ID, 目前消息id通过服务端生成, 后续考虑采用客户端生成方式
	msgId = cb.snow.ID()
	msg := &chat.MessageToMQ{
		MsgId: msgId,
		From:  req.From,
		To:    req.To,
		Body:  req.Body,
		Type:  req.Type,
		Sub:   req.Sub,
		Op:    req.Op,
	}
	return msgId, cb.mq.Push(ctx, msg)
}

func (cb *ChatBiz) PushRoom() error {
	return nil
}

func (cb *ChatBiz) BroadcastX() error {
	return nil
}
