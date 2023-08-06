package config

import (
	"github.com/cd-home/Hissssss/internal/pkg/cache"
	"github.com/cd-home/Hissssss/internal/pkg/etcdv3"
	"github.com/cd-home/Hissssss/internal/pkg/logger"
	"github.com/cd-home/Hissssss/internal/pkg/mq"
	"github.com/cd-home/Hissssss/internal/pkg/xgorm"
)

type Config struct {
	Version string `yaml:"version"`
	Kind    string `yaml:"kind"`
	Spec    Spec   `yaml:"spec"`
}

type Spec struct {
	Node     Node          `yaml:"service"`
	Logger   logger.Config `yaml:"logger"`
	Etcd     etcdv3.Config `yaml:"etcd"`
	Redis    cache.Config  `yaml:"redis"`
	Mysql    xgorm.Config  `yaml:"mysql"`
	RabbitMQ mq.Config     `yaml:"rabbitmq"`
	Queue    Queue         `yaml:"queue"`
}

type Node struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
	TTL  int64  `yaml:"ttl"`
	HTTP string `yaml:"http"`
}

type Queue struct {
	Single    string `yaml:"single"`
	Room      string `yaml:"room"`
	Broadcast string `yaml:"broadcast"`
}
