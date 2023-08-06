package mq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Config struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func NewRabbitMQ(config Config) *amqp.Connection {
	endpoint := fmt.Sprintf("amqp://%s:%s@%s", config.User, config.Password, config.Addr)
	conn, err := amqp.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
