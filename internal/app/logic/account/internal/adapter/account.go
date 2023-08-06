package adapter

import (
	"context"
	"github.com/cd-home/Hissssss/api/pb/common"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/pkg/model"
)

type AccountBiz interface {
	SignUp(name string, pwd string, way common.SignupWay, platform common.Platform) error // 用户名密码注册
	SignIn(name string, pwd string) (string, error)                                       // 登录
	Authenticate(name string, pwd string) (uint, error)                                   // 认证
	SignToken(uid uint) (string, error)                                                   // 签发token
	SignOut() bool                                                                        // 登出
	PunchIn() bool                                                                        // 签到 daily attendance
	Connect(ctx context.Context, uid int64, serverID string) error                        // 连接
	DisConnect(ctx context.Context, uid int64) error                                      // 断开连接
}

type AccountRepo interface {
	Create(*model.User) error
	Retrieve(name string) (*model.User, bool, error)
}

type AccountCache interface {
	Connect(ctx context.Context, uid int64, serverID string) error // 连接
	DisConnect(ctx context.Context, uid int64) error               // 断开连接
}
