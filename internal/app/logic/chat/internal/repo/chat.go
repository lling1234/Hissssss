package repo

import (
	"github.com/cd-home/Hissssss/internal/app/logic/chat/internal/adapter"
	"github.com/cd-home/Hissssss/internal/app/logic/chat/internal/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ChatRepo struct {
	cache *redis.Client
	db    *gorm.DB
}

func NewChatRepo(cache *redis.Client, db *gorm.DB) adapter.ChatRepo {
	return &ChatRepo{
		cache: cache,
		db:    db,
	}
}

func (c *ChatRepo) CreateAllMessage(msg *model.AllMessage) error {
	return nil
}
