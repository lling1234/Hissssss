package config

import (
	"github.com/cd-home/Hissssss/internal/pkg/config/node"
	"github.com/cd-home/Hissssss/internal/pkg/etcdv3"
	"github.com/cd-home/Hissssss/internal/pkg/logger"
)

type Config struct {
	Version string `yaml:"version"`
	Kind    string `yaml:"kind"`
	Spec    struct {
		Node   node.Config   `yaml:"node"`
		Logger logger.Config `yaml:"logger"`
		Etcd   etcdv3.Config `yaml:"etcd"`
	} `yaml:"spec"`
}
