package mq

import (
	"context"
	"errors"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"sync"
	"time"
)

type RabbitMQ struct {
	Logger          *zap.Logger
	config          Config
	Ready           bool
	Close           chan struct{}
	NotifyConnClose chan *amqp.Error
	NotifyChanClose chan *amqp.Error
	Conn            *amqp.Connection
	Ok              chan struct{}
	Client          map[string]*Client
	Queue           []string
	Retry           bool
	First           bool
	Once            *sync.Once
}

func New(config Config, logger *zap.Logger) *RabbitMQ {
	endpoint := fmt.Sprintf("amqp://%s:%s@%s", config.User, config.Password, config.Addr)
	r := &RabbitMQ{Logger: logger, config: config, Ok: make(chan struct{}, 1), First: true, Once: &sync.Once{}}
	go r.reconnect(endpoint)
	return r
}

type Client struct {
	Logger          *zap.Logger
	Ready           bool
	Queue           string
	Channel         *amqp.Channel
	Conn            *amqp.Connection
	NotifyConfirm   chan amqp.Confirmation
	NotifyChanClose chan *amqp.Error
	NotifyConnClose chan *amqp.Error
	Close           chan struct{}
}

func (r *RabbitMQ) SetUp(queue ...string) {
	r.Client = make(map[string]*Client)
	r.Close = make(chan struct{}, len(queue))
	r.NotifyConnClose = make(chan *amqp.Error, len(queue))
	r.NotifyChanClose = make(chan *amqp.Error, len(queue))
	r.Queue = queue
	r.First = false
	<-r.Ok
	for i := 0; i < len(queue); i++ {
		c := &Client{
			Logger: r.Logger.WithOptions(zap.Fields(zap.String("module", queue[i]))),
			Queue:  queue[i],
		}
		c.Close = make(chan struct{}, 1)
		go c.handleReInit(r.Conn, r.Close, r.NotifyConnClose, r.NotifyChanClose)
		r.Client[queue[i]] = c
	}
	return
}

func (r *RabbitMQ) reconnect(endpoint string) {
	for {
		r.Ready = false
		// 第一阶段 Connection
		_, err := r.connect(endpoint)
		if err != nil {
			r.Logger.Warn("failed to connect, retrying...")
			r.Once.Do(func() {
				r.Retry = !r.First
			})
			select {
			case <-r.Close:
				return
			case <-time.After(time.Second * 3):
			}
			continue
		}
		r.Ok <- struct{}{}
		if r.Retry {
			r.SetUp(r.Queue...)
		}
		// 第二阶段 Channel
		select {
		case <-r.Close:
		case <-r.NotifyConnClose:
		case <-r.NotifyChanClose:
		}
	}
}

func (r *RabbitMQ) connect(endpoint string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	r.Conn = conn
	r.NotifyConnClose = make(chan *amqp.Error, 1)
	r.Conn.NotifyClose(r.NotifyConnClose)
	r.Ready = true
	return conn, nil
}

func (c *Client) handleReInit(conn *amqp.Connection, close chan struct{}, notify, chanClose chan *amqp.Error) {
	for {
		c.Ready = false
		err := c.init(conn)
		if err != nil {
			c.Logger.Warn("failed to initialize channel. retrying...")
			select {
			case c := <-c.Close:
				close <- c
				return
			case n := <-c.NotifyConnClose:
				notify <- n
				return
			case <-time.After(time.Second * 3):
			}
			continue
		}
		select {
		case k := <-c.Close:
			close <- k
			c.Ready = false
			return
		case n := <-c.NotifyConnClose:
			notify <- n
			c.Ready = false
			return
		case m := <-c.NotifyChanClose:
			chanClose <- m
			c.Ready = false
			return
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

func (c *Client) Consume() (<-chan amqp.Delivery, error) {
	if !c.Ready {
		return nil, errors.New("failed to push: not connected")
	}
	// 应该综合带宽、每条消息的数据报大小、消费者线程处理的速率等等角度去考虑
	_ = c.Channel.Qos(30, 0, false)
	ch, err := c.Channel.Consume(c.Queue, "", false, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return ch, nil
}
