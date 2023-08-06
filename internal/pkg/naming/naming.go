package naming

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"go.uber.org/zap"
	gResolver "google.golang.org/grpc/resolver"
)

type Naming struct {
	ctx        context.Context
	cancel     context.CancelFunc
	etcdClient *clientv3.Client
	Manager    endpoints.Manager
	logger     *zap.Logger
}

func New(etcdClient *clientv3.Client, logger *zap.Logger) *Naming {
	return &Naming{etcdClient: etcdClient, logger: logger}
}

func (n *Naming) Register(service string, addr string, metadata string, ttl int64) error {
	n.ctx, n.cancel = context.WithCancel(context.Background())
	etcdManager, err := endpoints.NewManager(n.etcdClient, service)
	if err != nil {
		n.logger.Error(err.Error())
		return err
	}
	n.Manager = etcdManager
	lease, err := n.etcdClient.Grant(context.Background(), ttl)
	if err != nil {
		n.logger.Error(err.Error())
		return err
	}
	key := fmt.Sprintf("%s/%s", service, addr)
	// metadata目前先存server_id
	endpoint := endpoints.Endpoint{Addr: addr, Metadata: metadata}
	err = etcdManager.AddEndpoint(context.Background(), key, endpoint, clientv3.WithLease(lease.ID))
	if err != nil {
		n.logger.Error(err.Error())
		return err
	}
	alive, err := n.etcdClient.KeepAlive(n.ctx, lease.ID)
	if err != nil {
		n.logger.Error(err.Error())
		return err
	}
	go func() {
		for {
			select {
			case <-alive:
			case <-n.ctx.Done():
				return
			}
		}
	}()
	return nil
}

func (n *Naming) Discovery(service string) (string, gResolver.Builder) {
	etcdResolverBuilder, err := resolver.NewBuilder(n.etcdClient)
	if err != nil {
		n.logger.Error(err.Error())
	}
	target := fmt.Sprintf("etcd:///%s", service)
	return target, etcdResolverBuilder
}
