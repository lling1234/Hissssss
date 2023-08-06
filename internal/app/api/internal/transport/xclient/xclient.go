package xclient

import (
	"fmt"
	"github.com/cd-home/Hissssss/api/pb/account"
	"github.com/cd-home/Hissssss/api/pb/chat"
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
	Chat    chat.ChatClient       // logic层 chat服务
}

func New(logger *zap.Logger, naming *naming.Naming) RpcClient {
	ac, builder := naming.Discovery("account")
	accountConn, _ := grpc.Dial(
		ac,
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(builder),
	)
	ca, builder := naming.Discovery("chat")
	chatConn, _ := grpc.Dial(
		ca,
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(builder),
	)
	clients := RpcClient{
		Account: account.NewAccountClient(accountConn),
		Chat:    chat.NewChatClient(chatConn),
	}
	return clients
}
