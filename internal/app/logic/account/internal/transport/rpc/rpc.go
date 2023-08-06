package rpc

import (
	"github.com/cd-home/Hissssss/api/pb/account"
	"github.com/cd-home/Hissssss/internal/app/logic/account/config"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/svc"
	"github.com/cd-home/Hissssss/internal/pkg/naming"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"runtime"
)

type GrpcServer struct {
	fx.In
	Naming *naming.Naming
	Logger *zap.Logger
	Config config.Config
	Svc    *svc.Account
}

func StartRpcServer(gs GrpcServer) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	gs.Logger = gs.Logger.WithOptions(zap.Fields(zap.String("module", "account rpc server")))
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
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		// TODO  read from config and consider value should be
		//grpc.KeepaliveParams(keepalive.ServerParameters{}),
	)
	account.RegisterAccountServer(grpcServer, gs.Svc)
	gs.Logger.Info("serving grpc-server on " + gs.Config.Spec.Node.Addr)
	err = grpcServer.Serve(listener)
	if err != nil {
		gs.Logger.Error(err.Error())
		return
	}
}
