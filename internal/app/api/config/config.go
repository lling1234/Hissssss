package config

import (
	"github.com/cd-home/Hissssss/internal/pkg/etcdv3"
	"github.com/cd-home/Hissssss/internal/pkg/logger"
)

type Config struct {
	Version string `yaml:"version"`
	Kind    string `yaml:"kind"`
	Spec    Spec   `yaml:"spec"`
}

type Spec struct {
	Logger logger.Config `yaml:"logger"`
	Node   Node          `yaml:"node"`
	Etcd   etcdv3.Config `yaml:"etcd"`
}

type Node struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
	TTL  int64  `yaml:"ttl"`
	HTTP string `yaml:"http"`
}
