package main

import (
	"github.com/cd-home/Hissssss/internal/app/api/config"
	"github.com/cd-home/Hissssss/internal/app/api/internal/svc"
	"github.com/cd-home/Hissssss/internal/app/api/internal/transport/rpc"
	"github.com/cd-home/Hissssss/internal/app/api/internal/transport/xclient"
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
		"api",
		fx.Supply(c, c.Spec.Logger, c.Spec.Etcd),
		fx.Provide(logger.New),
		fx.Provide(etcdv3.New),
		fx.Provide(naming.New),
		fx.Provide(svc.New),
		fx.Provide(xclient.New),
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
	vp.SetConfigName("api")
	vp.AddConfigPath("etc")
	vp.AddConfigPath("../etc")
	vp.SetEnvPrefix("API")
	vp.AutomaticEnv()
	if err = vp.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err = vp.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
}
