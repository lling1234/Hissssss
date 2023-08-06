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
