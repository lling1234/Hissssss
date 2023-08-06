package svc

import (
	"context"
	"github.com/cd-home/Hissssss/api/pb/account"
	"github.com/cd-home/Hissssss/api/pb/api"
	"github.com/cd-home/Hissssss/api/pb/chat"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Api struct {
	api.UnimplementedApiServer
	logger  *zap.Logger
	account account.AccountClient
	chat    chat.ChatClient
}

func New(logger *zap.Logger, account account.AccountClient, chat chat.ChatClient) *Api {
	return &Api{
		logger:  logger.WithOptions(zap.Fields(zap.String("module", "api service"))),
		account: account,
		chat:    chat,
	}
}

func (a *Api) SignUp(ctx context.Context, req *api.SignUpRequest) (*api.SignUpReply, error) {
	reply, err := a.account.SignUp(ctx, &account.SignUpRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		a.logger.Error("signup err", zap.Error(err))
		return &api.SignUpReply{Code: 500, Message: "注册失败, 请重试."}, nil
	}
	return &api.SignUpReply{Code: reply.Code, Message: reply.Message}, nil
}

func (a *Api) SignIn(ctx context.Context, req *api.SignInRequest) (*api.SignInReply, error) {
	reply, err := a.account.SignIn(ctx, &account.SignInRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return &api.SignInReply{Code: 500, Message: err.Error()}, nil
	}
	return &api.SignInReply{Token: reply.Token, Message: reply.Message, Code: reply.Code}, nil
}

func (a *Api) SendMessage(ctx context.Context, req *api.SendMessageRequest) (*api.SendMessageReplyAck, error) {
	ack, err := a.chat.Push(ctx, &chat.SendMessageRequest{
		From: req.From,
		To:   req.To,
		Room: req.Room,
		Body: req.Body,
		Type: req.Type,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "发送消息失败")
	}
	return &api.SendMessageReplyAck{
		Code:  ack.Code,
		Msg:   ack.Msg,
		MsgId: ack.MsgId,
		Op:    ack.Op,
	}, nil
}
