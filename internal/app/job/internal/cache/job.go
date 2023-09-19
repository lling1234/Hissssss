package cache

import (
	"context"
	"github.com/cd-home/Hissssss/internal/app/job/internal/adapter"
	"github.com/cd-home/Hissssss/internal/pkg/cache"
	"go.uber.org/zap"
)

type JobCache struct {
	logger *zap.Logger
	cache  *cache.Cache
}

func NewJobCache(logger *zap.Logger, cache *cache.Cache) adapter.JobCache {
	return &JobCache{logger: logger, cache: cache}
}

func (j *JobCache) GetServerByUID(uid string) (string, error) {
	return j.cache.Get(context.Background(), uid)
}
