package rpc

import (
	"context"
	"github.com/cd-home/Hissssss/api/pb/connect"
	"github.com/cd-home/Hissssss/internal/app/connect/config"
	conn "github.com/cd-home/Hissssss/internal/app/connect/internal/connect"
	"github.com/cd-home/Hissssss/internal/app/connect/internal/svc"
	"github.com/cd-home/Hissssss/internal/pkg/naming"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"runtime"
)

type GrpcServer struct {
	fx.In
	Naming  *naming.Naming
	Logger  *zap.Logger
	Config  config.Config
	Svc     *svc.Message    // connect 提供的grpc服务
	Connect *conn.Websocket // Websocket
}

func StartRpcServer(lifecycle fx.Lifecycle, gs GrpcServer) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	gs.Logger = gs.Logger.WithOptions(zap.Fields(zap.String("module", "connect grpc server")))
	err := gs.Naming.Register(
		gs.Config.Spec.Node.Name,
		gs.Config.Spec.Node.Addr,
		gs.Config.Spec.Node.ID,
		gs.Config.Spec.Node.TTL,
	)
	if err != nil {
		gs.Logger.Error(err.Error())
		return
	}
	listener, err := net.Listen("tcp", gs.Config.Spec.Node.Addr)
	if err != nil {
		gs.Logger.Error(err.Error())
		return
	}
	grpcServer := grpc.NewServer()
	connect.RegisterPushMessageServer(grpcServer, gs.Svc)
	gs.Logger.Info("serving grpc-server on " + gs.Config.Spec.Node.Addr)
	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		err = gs.Connect.InitWebSocket(context.Background())
		if err != nil {
			panic(err)
		}
	}()
	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			gs.Connect.Clear()
			return nil
		},
	})
}
