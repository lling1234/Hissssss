package xmongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type Config struct {
	Host string   `yaml:"host"`
	Port string   `yaml:"port"`
	User string   `yaml:"user"`
	Pwd  string   `yaml:"pwd"`
	DB   []string `yaml:"db"`
}

type XMongo struct {
	ctx      context.Context
	client   *mongo.Client
	database map[string]*mongo.Database
}

func New(config Config) *XMongo {
	dns := fmt.Sprintf("mongodb://%s:%s@%s:%s/?connect=direct&authSource=admin", config.User, config.Pwd, config.Host, config.Port)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dns))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	mgo := &XMongo{
		ctx:      context.Background(),
		client:   client,
		database: make(map[string]*mongo.Database, len(config.DB)),
	}
	for i := 0; i < len(config.DB); i++ {
		mgo.database[config.DB[i]] = client.Database(config.DB[i])
	}
	return mgo
}
