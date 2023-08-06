package repo

import (
	"errors"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/adapter"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/pkg/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AccountRepo struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewAccountRepo(logger *zap.Logger, db *gorm.DB) adapter.AccountRepo {
	return &AccountRepo{
		logger: logger.WithOptions(zap.Fields(zap.String("module", "account repo"))),
		db:     db,
	}
}

func (a *AccountRepo) Create(user *model.User) error {
	return a.db.Create(user).Error
}

func (a *AccountRepo) Retrieve(name string) (*model.User, bool, error) {
	var u model.User
	err := a.db.First(&u, "username = ?", name).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	return &u, true, nil
}
