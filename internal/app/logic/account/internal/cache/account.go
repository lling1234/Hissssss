package cache

import (
	"context"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/adapter"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/pkg/key"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type AccountCache struct {
	logger *zap.Logger
	cache  *redis.Client
}

func NewAccountCache(logger *zap.Logger, cache *redis.Client) adapter.AccountCache {
	return &AccountCache{
		logger: logger,
		cache:  cache,
	}
}

func (a *AccountCache) Connect(ctx context.Context, uid int64, serverID string) error {
	// redis version <= 6 to Zero expiration means the key has no expiration time.
	return a.cache.Set(ctx, key.GenUidMappingServer(uid), serverID, 0).Err()
}

func (a *AccountCache) DisConnect(ctx context.Context, uid int64) error {
	return a.cache.Del(ctx, key.GenUidMappingServer(uid)).Err()
}
