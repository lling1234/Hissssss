package config

import (
	"github.com/cd-home/Hissssss/internal/pkg/cache"
	"github.com/cd-home/Hissssss/internal/pkg/config/node"
	"github.com/cd-home/Hissssss/internal/pkg/etcdv3"
	"github.com/cd-home/Hissssss/internal/pkg/jwt"
	"github.com/cd-home/Hissssss/internal/pkg/logger"
	"github.com/cd-home/Hissssss/internal/pkg/xgorm"
	"github.com/cd-home/Hissssss/internal/pkg/xmongo"
)

type Config struct {
	Version string `yaml:"version"`
	Kind    string `yaml:"kind"`
	Spec    struct {
		Node   node.Config   `yaml:"node"`
		Logger logger.Config `yaml:"logger"`
		Etcd   etcdv3.Config `yaml:"etcd"`
		Redis  cache.Config  `yaml:"redis"`
		Mysql  xgorm.Config  `yaml:"mysql"`
		Jwt    jwt.Config    `yaml:"jwt"`
		Mongo  xmongo.Config `yaml:"mongo"`
	} `yaml:"spec"`
}
