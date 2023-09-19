package config

import (
	"github.com/cd-home/Hissssss/internal/pkg/cache"
	"github.com/cd-home/Hissssss/internal/pkg/config/node"
	"github.com/cd-home/Hissssss/internal/pkg/config/queue"
	"github.com/cd-home/Hissssss/internal/pkg/etcdv3"
	"github.com/cd-home/Hissssss/internal/pkg/logger"
	"github.com/cd-home/Hissssss/internal/pkg/mq"
	"github.com/cd-home/Hissssss/internal/pkg/xmongo"
)

type Config struct {
	Version string `yaml:"version"`
	Kind    string `yaml:"kind"`
	Spec    struct {
		Node     node.Config   `yaml:"node"`
		Logger   logger.Config `yaml:"logger"`
		Etcd     etcdv3.Config `yaml:"etcd"`
		Redis    cache.Config  `yaml:"redis"`
		RabbitMQ mq.Config     `yaml:"rabbitmq"`
		Queue    queue.Config  `yaml:"queue"`
		Mongo    xmongo.Config `yaml:"mongo"`
	} `yaml:"spec"`
}
