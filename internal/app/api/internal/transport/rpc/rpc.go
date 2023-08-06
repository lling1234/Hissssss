package rpc

import (
	"github.com/cd-home/Hissssss/api/pb/api"
	"github.com/cd-home/Hissssss/internal/app/api/config"
	svc "github.com/cd-home/Hissssss/internal/app/api/internal/svc"
	"github.com/cd-home/Hissssss/internal/app/api/internal/transport/http"
	"github.com/cd-home/Hissssss/internal/pkg/naming"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"runtime"
)

type GrpcServer struct {
	fx.In
	Naming *naming.Naming
	Logger *zap.Logger
	Config config.Config
	Svc    *svc.Api
}

func StartRpcServer(gs GrpcServer) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	gs.Logger = gs.Logger.WithOptions(zap.Fields(zap.String("module", "api grpc server")))
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
	api.RegisterApiServer(grpcServer, gs.Svc)
	gs.Logger.Info("serving grpc-server on " + gs.Config.Spec.Node.Addr)
	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			gs.Logger.Error(err.Error())
			return
		}
	}()
	http.StartHttpServer(gs.Config, gs.Logger)
}
