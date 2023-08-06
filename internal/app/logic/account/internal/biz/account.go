package biz

import (
	"context"
	"errors"
	"github.com/cd-home/Hissssss/api/pb/common"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/adapter"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/pkg/model"
	"github.com/cd-home/Hissssss/internal/pkg/jwt"
	"github.com/cd-home/Hissssss/internal/pkg/tool/bcrypt/pwd"
	"go.uber.org/zap"
)

type AccountBiz struct {
	logger *zap.Logger
	repo   adapter.AccountRepo
	cache  adapter.AccountCache
	jwt    jwt.Config
}

func NewAccountBiz(logger *zap.Logger, repo adapter.AccountRepo, cache adapter.AccountCache, jwt jwt.Config) adapter.AccountBiz {
	return &AccountBiz{
		logger: logger.WithOptions(zap.Fields(zap.String("module", "account biz"))),
		repo:   repo,
		cache:  cache,
		jwt:    jwt,
	}
}

func (ab *AccountBiz) SignUp(name string, password string, way common.SignupWay, platform common.Platform) error {
	ab.logger.Debug("[signup]: ", zap.String("name", name))
	encrypt, _ := pwd.Encrypt(password)
	return ab.repo.Create(&model.User{
		UserName:  name,
		Password:  encrypt,
		SignupWay: way,
	})
}

func (ab *AccountBiz) SignIn(name string, pwd string) (string, error) {
	uid, err := ab.Authenticate(name, pwd)
	if err != nil {
		return "", err
	}
	return ab.SignToken(uid)
}

func (ab *AccountBiz) Authenticate(name string, p string) (uint, error) {
	user, ok, err := ab.repo.Retrieve(name)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, errors.New("用户不存在")
	}
	if !pwd.Verify(p, user.Password) {
		return 0, errors.New("密码验证不通过")
	}
	return user.ID, nil
}

func (ab *AccountBiz) SignToken(uid uint) (string, error) {
	return jwt.SignJwtToken(uid, ab.jwt)
}

func (ab *AccountBiz) SignOut() bool {
	//TODO implement me
	panic("implement me")
}

func (ab *AccountBiz) PunchIn() bool {
	//TODO implement me
	panic("implement me")
}

func (ab *AccountBiz) Connect(ctx context.Context, uid int64, serverID string) error {
	return ab.cache.Connect(ctx, uid, serverID)
}

func (ab *AccountBiz) DisConnect(ctx context.Context, uid int64) error {
	return ab.cache.DisConnect(ctx, uid)
}
