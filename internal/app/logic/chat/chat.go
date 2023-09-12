package main

import (
	"github.com/cd-home/Hissssss/internal/app/logic/chat/config"
	"github.com/cd-home/Hissssss/internal/app/logic/chat/internal/biz"
	cacheBiz "github.com/cd-home/Hissssss/internal/app/logic/chat/internal/cache"
	chatMQ "github.com/cd-home/Hissssss/internal/app/logic/chat/internal/mq"
	"github.com/cd-home/Hissssss/internal/app/logic/chat/internal/repo"
	"github.com/cd-home/Hissssss/internal/app/logic/chat/internal/svc"
	"github.com/cd-home/Hissssss/internal/app/logic/chat/internal/transport/rpc"
	"github.com/cd-home/Hissssss/internal/pkg/cache"
	"github.com/cd-home/Hissssss/internal/pkg/etcdv3"
	"github.com/cd-home/Hissssss/internal/pkg/logger"
	"github.com/cd-home/Hissssss/internal/pkg/mq"
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
		"chat",
		fx.Supply(c, c.Spec.Etcd, c.Spec.Logger, c.Spec.Redis, c.Spec.RabbitMQ, c.Spec.Queue, c.Spec.Mysql),
		fx.Supply(snow),
		fx.Provide(logger.New),
		fx.Provide(cache.NewRedis),
		fx.Provide(etcdv3.New),
		fx.Provide(naming.New),
		fx.Provide(chatMQ.NewRabbitMQ),
		fx.Provide(mq.New),
		fx.Provide(xgorm.New),
		fx.Provide(repo.NewChatRepo),
		fx.Provide(cacheBiz.NewChatCache),
		fx.Provide(biz.NewChatBiz),
		fx.Provide(svc.NewChat),
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
	vp.SetConfigName("chat")
	vp.AddConfigPath("etc")
	vp.AddConfigPath("../etc")
	vp.SetEnvPrefix("CHAT")
	vp.AutomaticEnv()
	if err = vp.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err = vp.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
}
