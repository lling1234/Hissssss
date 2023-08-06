package mq

import (
	"context"
	"errors"
	"fmt"
	"github.com/cd-home/Hissssss/internal/pkg/mq"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"time"
)

type RabbitMQ struct {
	Logger *zap.Logger
	config mq.Config
}

func New(config mq.Config, logger *zap.Logger) *RabbitMQ {
	return &RabbitMQ{Logger: logger, config: config}
}

type Client struct {
	Logger          *zap.Logger
	Ready           bool
	Queue           string
	Channel         *amqp.Channel
	Conn            *amqp.Connection
	Close           chan struct{}
	NotifyConfirm   chan amqp.Confirmation
	NotifyConnClose chan *amqp.Error
	NotifyChanClose chan *amqp.Error
}

func (r *RabbitMQ) NewClient(queue string) *Client {
	c := &Client{
		Close:  make(chan struct{}, 1),
		Logger: r.Logger.WithOptions(zap.Fields(zap.String("module", queue))),
		Queue:  queue,
	}
	endpoint := fmt.Sprintf("amqp://%s:%s@%s", r.config.User, r.config.Password, r.config.Addr)
	go c.reconnect(endpoint)
	return c
}

func (c *Client) reconnect(endpoint string) {
	for {
		c.Ready = false
		// 第一阶段 Connection
		conn, err := c.connect(endpoint)
		if err != nil {
			c.Logger.Warn("failed to connect, retrying...")
			select {
			case <-c.Close:
				return
			case <-time.After(time.Second * 3):
			}
			continue
		}
		// 第二阶段 Channel
		if done := c.handleReInit(conn); done {
			break
		}
	}
}

func (c *Client) connect(endpoint string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	c.Conn = conn
	c.NotifyConnClose = make(chan *amqp.Error, 1)
	c.Conn.NotifyClose(c.NotifyConnClose)
	return conn, nil
}

func (c *Client) handleReInit(conn *amqp.Connection) bool {
	for {
		c.Ready = false
		err := c.init(conn)
		if err != nil {
			c.Logger.Warn("failed to initialize channel. retrying...")
			select {
			case <-c.Close:
				return true
			case <-c.NotifyConnClose:
				return false
			case <-time.After(time.Second * 3):
			}
			continue
		}
		select {
		case <-c.Close:
			return true
		case <-c.NotifyConnClose:
			return false
		case <-c.NotifyChanClose:
			c.Logger.Warn("channel closed. re-running init...")
		}
	}
}

func (c *Client) init(conn *amqp.Connection) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	err = ch.Confirm(false)
	if err != nil {
		return err
	}
	_, err = ch.QueueDeclare(c.Queue, true, false, false, false, nil)
	if err != nil {
		return err
	}
	c.Channel = ch
	c.NotifyConfirm = make(chan amqp.Confirmation, 1)
	c.NotifyChanClose = make(chan *amqp.Error, 1)
	c.Channel.NotifyClose(c.NotifyChanClose)
	c.Channel.NotifyPublish(c.NotifyConfirm)
	c.Ready = true
	return nil
}

func (c *Client) Push(ctx context.Context, body []byte) error {
	if !c.Ready {
		return errors.New("failed to push: not connected")
	}
	var err error
	var retries = 3
	for i := 0; i < retries; i++ {
		err = c.Channel.PublishWithContext(
			ctx, "", c.Queue, false, false, amqp.Publishing{
				ContentType:  "text/plain",
				DeliveryMode: amqp.Persistent,
				Body:         body,
			})
		if err != nil {
			continue
		}
		if confirm := <-c.NotifyConfirm; confirm.Ack {
			c.Logger.Debug("Push confirmed:", zap.Any("confirm.deliveryTag", confirm.DeliveryTag))
			return nil
		}
	}
	return err
}
