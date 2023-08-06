package main

import (
	"github.com/cd-home/Hissssss/internal/app/connect/config"
	"github.com/cd-home/Hissssss/internal/app/connect/internal/connect"
	"github.com/cd-home/Hissssss/internal/app/connect/internal/svc"
	"github.com/cd-home/Hissssss/internal/app/connect/internal/transport/rpc"
	"github.com/cd-home/Hissssss/internal/app/connect/internal/transport/xclient"
	"github.com/cd-home/Hissssss/internal/pkg/etcdv3"
	"github.com/cd-home/Hissssss/internal/pkg/logger"
	"github.com/cd-home/Hissssss/internal/pkg/naming"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log"
)

var c config.Config

func main() {
	app := fx.Module(
		"connect",
		fx.Supply(c, c.Spec.Etcd, c.Spec.Logger),
		fx.Provide(logger.New),
		fx.Provide(etcdv3.New),
		fx.Provide(connect.New),
		fx.Provide(xclient.New),
		fx.Provide(svc.NewMessage),
		fx.Provide(naming.New),
		fx.Invoke(rpc.StartRpcServer),
		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)
	fx.New(app).Run()
}

func init() {
	var err error
	vp := viper.New()
	vp.SetConfigType("yaml")
	vp.SetConfigName("connect")
	vp.AddConfigPath("etc")
	vp.AddConfigPath("../etc")
	vp.SetEnvPrefix("CONNECT")
	vp.AutomaticEnv()
	if err = vp.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err = vp.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
}
