package main

import (
	"github.com/cd-home/Hissssss/internal/app/logic/account/config"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/biz"
	cacheBiz "github.com/cd-home/Hissssss/internal/app/logic/account/internal/cache"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/repo"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/svc"
	"github.com/cd-home/Hissssss/internal/app/logic/account/internal/transport/rpc"
	"github.com/cd-home/Hissssss/internal/pkg/cache"
	"github.com/cd-home/Hissssss/internal/pkg/etcdv3"
	"github.com/cd-home/Hissssss/internal/pkg/logger"
	"github.com/cd-home/Hissssss/internal/pkg/naming"
	"github.com/cd-home/Hissssss/internal/pkg/tool/snowid"
	"github.com/cd-home/Hissssss/internal/pkg/xgorm"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log"
	"strconv"
)

var c config.Config

func main() {
	id, _ := strconv.ParseInt(c.Spec.Node.ID, 10, 64)
	snow, err := snowid.New(id%30, id%10)
	if err != nil {
		log.Println(err)
		return
	}
	app := fx.Module(
		"Account",
		fx.Supply(c, c.Spec.Node, c.Spec.Etcd, c.Spec.Logger, c.Spec.Redis, c.Spec.Mysql, c.Spec.Jwt),
		fx.Supply(snow),
		fx.Provide(logger.New),
		fx.Provide(etcdv3.New),
		fx.Provide(xgorm.New),
		fx.Provide(cache.New),
		fx.Provide(repo.NewAccountRepo),
		fx.Provide(cacheBiz.NewAccountCache),
		fx.Provide(biz.NewAccountBiz),
		fx.Provide(svc.NewAccount),
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
	vp.SetConfigName("account")
	vp.AddConfigPath("etc")
	vp.AddConfigPath("../etc")
	vp.SetEnvPrefix("ACCOUNT")
	vp.AutomaticEnv()
	if err = vp.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err = vp.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
}
