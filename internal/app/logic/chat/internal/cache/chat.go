package cache

import (
	"context"
	"github.com/cd-home/Hissssss/internal/app/logic/chat/internal/adapter"
	"github.com/cd-home/Hissssss/internal/pkg/cache"
	"go.uber.org/zap"
)

type ChatCache struct {
	logger *zap.Logger
	cache  *cache.Cache
}

func NewChatCache(logger *zap.Logger, cache *cache.Cache) adapter.ChatCache {
	return &ChatCache{
		logger: logger,
		cache:  cache,
	}
}

func (c *ChatCache) GetServerByUID(key string) (string, error) {
	return c.cache.Get(context.Background(), key)
}

func (c *ChatCache) GetUserMessageID(key string) (int64, error) {
	return 0, nil
}
