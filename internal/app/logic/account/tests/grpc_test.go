package tests

import (
	"context"
	"fmt"
	"github.com/cd-home/Hissssss/api/pb/account"
	"github.com/cd-home/Hissssss/internal/pkg/etcdv3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestGrpcServer(t *testing.T) {
	client := etcdv3.New(etcdv3.Config{Addr: []string{"10.211.55.18:2379"}})
	etcdResolverBuilder, err := resolver.NewBuilder(client)
	if err != nil {
		t.Log(err)
		return
	}
	target := fmt.Sprintf("etcd:///%s", "account")
	conn, err := grpc.Dial(
		target,
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(etcdResolverBuilder),
	)
	if err != nil {
		t.Log(err)
		return
	}
	grpcClient := account.NewAccountClient(conn)
	resp, err := grpcClient.SignUp(context.Background(), &account.SignUpRequest{
		Username: "1",
		Password: "2",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
