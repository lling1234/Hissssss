package cache

import (
	"context"
	"github.com/cd-home/Hissssss/internal/app/job/internal/adapter"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type JobCache struct {
	logger *zap.Logger
	cache  *redis.Client
}

func NewJobCache(logger *zap.Logger, cache *redis.Client) adapter.JobCache {
	return &JobCache{logger: logger, cache: cache}
}

func (j *JobCache) GetServerByUID(uid string) (string, error) {
	return j.cache.Get(context.Background(), uid).Result()
}
