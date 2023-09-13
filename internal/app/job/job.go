package main

import (
	"github.com/cd-home/Hissssss/internal/app/job/config"
	"github.com/cd-home/Hissssss/internal/app/job/internal/biz"
	cacheBiz "github.com/cd-home/Hissssss/internal/app/job/internal/cache"
	"github.com/cd-home/Hissssss/internal/app/job/internal/connect"
	"github.com/cd-home/Hissssss/internal/app/job/internal/job"
	chatMQ "github.com/cd-home/Hissssss/internal/app/job/internal/mq"
	"github.com/cd-home/Hissssss/internal/pkg/cache"
	"github.com/cd-home/Hissssss/internal/pkg/etcdv3"
	"github.com/cd-home/Hissssss/internal/pkg/logger"
	"github.com/cd-home/Hissssss/internal/pkg/mq"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log"
)

var c config.Config

func main() {
	app := fx.Module(
		"job",
		fx.Supply(c, c.Spec.Etcd, c.Spec.Logger, c.Spec.Node, c.Spec.Queue, c.Spec.RabbitMQ, c.Spec.Redis),
		fx.Provide(logger.New),
		fx.Provide(etcdv3.New),
		fx.Provide(mq.New),
		fx.Provide(cache.New),
		fx.Provide(chatMQ.NewRabbitMQ),
		fx.Provide(cacheBiz.NewJobCache),
		fx.Provide(biz.NewJobBiz),
		fx.Provide(connect.New),
		fx.Invoke(job.New),
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
	vp.SetConfigName("job")
	vp.AddConfigPath("etc")
	vp.AddConfigPath("../etc")
	vp.SetEnvPrefix("JOB")
	vp.AutomaticEnv()
	if err = vp.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err = vp.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
}
