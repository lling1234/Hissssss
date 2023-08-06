package svc

import (
	"context"
	"github.com/cd-home/Hissssss/api/pb/account"
	"github.com/cd-home/Hissssss/api/pb/common"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/adapter"
	"go.uber.org/zap"
)

type Account struct {
	account.UnimplementedAccountServer
	logger *zap.Logger
	biz    adapter.AccountBiz
}

func NewAccount(logger *zap.Logger, biz adapter.AccountBiz) *Account {
	return &Account{
		logger: logger.WithOptions(zap.Fields(zap.String("module", "account service"))),
		biz:    biz,
	}
}

// SignUp 目前默认Web端用户注册
func (a *Account) SignUp(ctx context.Context, req *account.SignUpRequest) (*account.SignUpReply, error) {
	if err := a.biz.SignUp(req.Username, req.Password, common.SignupWay_Username, common.Platform_Web); err != nil {
		a.logger.Error("[signup]: ", zap.Error(err))
		return nil, err
	}
	return &account.SignUpReply{
		Code:    200,
		Message: "OK",
	}, nil
}

// SignIn 登录
func (a *Account) SignIn(ctx context.Context, req *account.SignInRequest) (*account.SignInReply, error) {
	token, err := a.biz.SignIn(req.Username, req.Password)
	if err != nil {
		a.logger.Warn(err.Error())
		return &account.SignInReply{Code: 500, Message: err.Error()}, nil
	}
	return &account.SignInReply{Code: 200, Message: "success", Token: token}, nil
}

// Connect 连接信息
func (a *Account) Connect(ctx context.Context, req *account.ConnectRequest) (*account.ConnectReply, error) {
	a.logger.Debug("[Connect]: connect call account", zap.Any("uid", req.Uid))
	err := a.biz.Connect(ctx, req.Uid, req.ServerID)
	if err != nil {
		a.logger.Error(err.Error())
		// TODO 优化code, message管理
		return &account.ConnectReply{
			Code:    -1,
			Message: "保持连接信息失败",
		}, err
	}
	return &account.ConnectReply{Code: 200, Message: "success"}, nil
}

// DisConnect 清除连接信息
func (a *Account) DisConnect(ctx context.Context, req *account.DisConnectRequest) (*account.DisConnectReply, error) {
	a.logger.Debug("[DisConnect]: connect call account", zap.Any("uid", req.Uid))
	err := a.biz.DisConnect(ctx, req.Uid)
	if err != nil {
		return &account.DisConnectReply{
			Code:    -1,
			Message: "清除连接失败",
		}, err
	}
	return &account.DisConnectReply{
		Code:    200,
		Message: "success",
	}, nil
}
