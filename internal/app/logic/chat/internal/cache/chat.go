package cache

import (
	"context"
	"github.com/cd-home/Hissssss/internal/app/logic/chat/internal/adapter"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type ChatCache struct {
	logger *zap.Logger
	cache  *redis.Client
}

func NewChatCache(logger *zap.Logger, cache *redis.Client) adapter.ChatCache {
	return &ChatCache{
		logger: logger,
		cache:  cache,
	}
}

func (c *ChatCache) GetServerByUID(key string) (string, error) {
	return c.cache.Get(context.Background(), key).Result()
}

func (c *ChatCache) GetUserMessageID(key string) (int64, error) {
	return c.cache.Incr(context.Background(), key).Result()
}
