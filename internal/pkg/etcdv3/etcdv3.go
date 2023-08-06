package etcdv3

import (
	"go.etcd.io/etcd/client/v3"
	"time"
)

type Config struct {
	Addr []string `yaml:"addr"`
}

func New(config Config) *clientv3.Client {
	etcdClient, _ := clientv3.New(clientv3.Config{
		Endpoints:   config.Addr,
		DialTimeout: time.Second * 3,
	})
	return etcdClient
}
