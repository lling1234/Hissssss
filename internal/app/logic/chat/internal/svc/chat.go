package svc

import (
	"context"
	"github.com/cd-home/Hissssss/api/pb/chat"
	"github.com/cd-home/Hissssss/api/pb/common"
	"github.com/cd-home/Hissssss/internal/app/logic/chat/internal/adapter"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Chat struct {
	chat.UnimplementedChatServer
	logger *zap.Logger
	biz    adapter.ChatBiz
}

func NewChat(logger *zap.Logger, biz adapter.ChatBiz) *Chat {
	return &Chat{
		logger: logger.WithOptions(zap.Fields(zap.String("module", "chat svc"))),
		biz:    biz,
	}
}

// Push 发送单聊消息
func (c *Chat) Push(ctx context.Context, req *chat.SendMessageRequest) (*chat.SendMessageReply, error) {
	msgId, err := c.biz.Push(ctx, req)
	if err != nil {
		c.logger.Error(err.Error())
		return nil, status.Errorf(codes.Unimplemented, "发送错误")
	}
	return &chat.SendMessageReply{
		Code:  200,
		Msg:   "OK",
		MsgId: msgId,
		Op:    common.OP_ACK,
	}, nil
}

func (c *Chat) PushRoom(ctx context.Context, req *chat.SendMessageRequest) (*chat.SendMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushRoom not implemented")
}
func (c *Chat) Broadcast(ctx context.Context, req *chat.SendMessageRequest) (*chat.SendMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
