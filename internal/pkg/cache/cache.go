package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type Config struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Cache struct {
	ctx    context.Context
	client *redis.Client
}

func New(config Config) *Cache {
	return &Cache{
		ctx:    context.TODO(),
		client: NewRedis(config),
	}
}

func NewRedis(config Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         config.Addr,
		Password:     config.Password,
		DB:           config.DB,
		PoolSize:     10,
		MinIdleConns: 5,
		MaxIdleConns: 10,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		PoolTimeout:  2 * time.Second,
	})
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Redis Connection Successful: " + pong)
	return rdb
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *Cache) Set(ctx context.Context, key string, value any, expire time.Duration) error {
	return c.client.Set(ctx, key, value, expire).Err()
}

func (c *Cache) SetNX(ctx context.Context, key string, value any, expire time.Duration) error {
	return c.client.SetNX(ctx, key, value, expire).Err()
}

func (c *Cache) SetXX(ctx context.Context, key string, value any, expire time.Duration) error {
	return c.client.SetXX(ctx, key, value, expire).Err()
}

func (c *Cache) SetEx(ctx context.Context, key string, value any, expire time.Duration) error {
	return c.client.SetEx(ctx, key, value, expire).Err()
}

func (c *Cache) Del(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}
