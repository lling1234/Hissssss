package biz

import (
	"context"
	"errors"
	"github.com/cd-home/Hissssss/api/pb/common"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/adapter"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/pkg/model"
	"github.com/cd-home/Hissssss/internal/pkg/code"
	"github.com/cd-home/Hissssss/internal/pkg/jwt"
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
	_, ok, err := ab.repo.Retrieve(name)
	if err != nil {
		return code.Rsp(code.InternalError)
	}
	if ok {
		return code.Rsp(code.NameAlreadyExists)
	}
	if ab.repo.Create(&model.User{
		Name:      name,
		Password:  password,
		SignupWay: way,
	}); err != nil {
		return code.Rsp(code.InternalError)
	}
	return nil
}

func (ab *AccountBiz) SignIn(name string, pwd string) (string, error) {
	uid, err := ab.Authenticate(name, pwd)
	if err != nil {
		return "", err
	}
	return ab.SignToken(uid)
}

func (ab *AccountBiz) Authenticate(name string, password string) (int64, error) {
	doc, ok, err := ab.repo.Retrieve(name)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, errors.New("用户不存在")
	}
	if !doc.Verify(password) {
		return 0, errors.New("密码验证不通过")
	}
	return doc.ID, nil
}

func (ab *AccountBiz) SignToken(uid int64) (string, error) {
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
