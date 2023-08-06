package xclient

import (
	"fmt"
	"github.com/cd-home/Hissssss/api/pb/account"
	"github.com/cd-home/Hissssss/internal/pkg/naming"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

// RpcClient fx.Out 多对象返回注入
type RpcClient struct {
	fx.Out
	Account account.AccountClient // logic层 account服务
}

func New(logger *zap.Logger, naming *naming.Naming) RpcClient {
	target, builder := naming.Discovery("account")
	conn, err := grpc.Dial(
		target,
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(builder),
	)
	if err != nil {
		logger.Error(err.Error())
		return RpcClient{}
	}
	clients := RpcClient{
		Account: account.NewAccountClient(conn),
	}
	return clients
}
