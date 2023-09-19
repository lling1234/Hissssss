package svc

import (
	"context"
	"github.com/cd-home/Hissssss/api/pb/account"
	"github.com/cd-home/Hissssss/api/pb/api"
	"github.com/cd-home/Hissssss/api/pb/chat"
	"github.com/cd-home/Hissssss/api/pb/common"
	"github.com/cd-home/Hissssss/internal/pkg/code"
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
		codex, message := code.From(err)
		a.logger.Error("signup err", zap.Int64("code", codex), zap.String("message", message))
		return &api.SignUpReply{Code: codex, Message: message}, nil
	}
	return &api.SignUpReply{Code: reply.Code, Message: reply.Message}, nil
}

func (a *Api) SignIn(ctx context.Context, req *api.SignInRequest) (*api.SignInReply, error) {
	reply, err := a.account.SignIn(ctx, &account.SignInRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		codex, message := code.From(err)
		a.logger.Error("signin err", zap.Int64("code", codex), zap.String("message", message))
		return &api.SignInReply{Code: codex, Message: message}, nil
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
		Op:   common.OP_REQ,
		Sub:  common.Message_User,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "发送消息失败")
	}
	return &api.SendMessageReplyAck{
		Code:  ack.Code,
		Msg:   ack.Msg,
		MsgId: ack.MsgId,
		Op:    ack.Op, // TODO 普通API调用不需要 回复 OP
	}, nil
}

func (a *Api) AckMessage(ctx context.Context, req *api.AckMessageReqeust) (*api.AckMessageReplyAck, error) {
	_, err := a.chat.Push(ctx, &chat.SendMessageRequest{
		From: 1000,
		To:   req.From,
		Type: common.PushType_Single,
		Op:   common.OP_ACK,
		Sub:  common.Message_System,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "发送消息失败")
	}
	return &api.AckMessageReplyAck{
		From:  req.From,
		To:    req.To,
		MsgId: req.MsgId,
		Op:    common.OP_ACK,
		Sub:   common.Message_System,
	}, nil
}
