package repo

import (
	"github.com/cd-home/Hissssss/internal/app/logic/chat/internal/adapter"
	"github.com/cd-home/Hissssss/internal/app/logic/chat/internal/model"
	"gorm.io/gorm"
)

type ChatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) adapter.ChatRepo {
	return &ChatRepo{
		db: db,
	}
}

func (c *ChatRepo) CreateAllMessage(msg *model.AllMessage) error {
	return nil
}
