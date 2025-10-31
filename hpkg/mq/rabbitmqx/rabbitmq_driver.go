package rabbitmqx

import (
	"context"
	"encoding/json"
	configs "github.com/jhinih/hin/hconfig"
	"github.com/jhinih/hin/hglobal"
	"github.com/jhinih/hin/hlog/zlog"
	"github.com/jhinih/hin/hpkg/mq"
	"github.com/jhinih/hin/hutils/workerpoolUtils"
	"github.com/streadway/amqp"
	"strconv"
	"sync"
	"sync/atomic"
)

type RabbitMQ struct {
	conn  *amqp.Connection
	pool  []*amqp.Channel
	index uint32
	mu    sync.Mutex
}

func NewRabbitMQ() mq.MQ {
	return &RabbitMQ{}
}

const (
	core     = 4  // 最小并发
	capacity = 32 // 最大并发
)

func (r *RabbitMQ) InitMQ(config configs.Config) error {
	conn, err := amqp.Dial(config.MQ.RabbitMQ.Dsn)
	if err != nil {
		return err
	}
	r.conn = conn
	PoolSize, _ := strconv.ParseInt(config.MQ.RabbitMQ.ChannelPoolSize, 10, 64)
	poolSize := int(PoolSize)
	if poolSize <= 0 {
		poolSize = 8
	}

	r.pool = make([]*amqp.Channel, poolSize)
	for i := 0; i < poolSize; i++ {
		ch, err := conn.Channel()
		if err != nil {
			return err
		}
		r.pool[i] = ch
	}
	hglobal.RabbitMQ = r

	zlog.Infof("RabbitMQ 初始化成功，channel 池大小：%d", poolSize)
	return r.initProducer()
}

// 初始化生产者
func (r *RabbitMQ) initProducer() error {
	infra := []struct {
		Exchange string
		Type     string
		Queue    string
		Key      string
	}{
		//交换器，队列
		{"upload", "direct", "upload_queue", "upload_key"},
		{"email", "direct", "email_queue", "email_key"},
		{"order", "topic", "order_event", "order.created"},
	}

	ch := r.getChannel()
	for _, v := range infra {
		if err := ch.ExchangeDeclare(v.Exchange, v.Type, true, false, false, false, nil); err != nil {
			return err
		}
		if _, err := ch.QueueDeclare(v.Queue, true, false, false, false, nil); err != nil {
			return err
		}
		if err := ch.QueueBind(v.Queue, v.Key, v.Exchange, false, nil); err != nil {
			return err
		}
	}
	return nil
}
func (r *RabbitMQ) Push(exchange, key string, task interface{}) error {
	ch := r.getChannel()
	body, _ := json.Marshal(task)
	return ch.Publish(
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}

func (r *RabbitMQ) Consume(queue string, handler func(msg []byte) error) error {
	ch := r.getChannel()
	if err := ch.Qos(capacity, 0, false); err != nil {
		return err
	}
	deliveries, err := ch.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	wp := workerpoolUtils.New(core, capacity)
	defer wp.Stop()

	// 优雅退出
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-hglobal.CtxDone() // 程序退出时调用
		cancel()
		ch.Close()
	}()

	for d := range deliveries {
		msg := d
		wp.Submit(func() {
			if err := handler(msg.Body); err != nil {
				zlog.Errorf("消费失败: %v", err)
				_ = msg.Nack(false, true) // 重入队
				return
			}
			_ = msg.Ack(false)
		})
	}
	<-hglobal.CtxDone() //  阻塞在这里
	return nil
}

func (r *RabbitMQ) Close() error {
	// 1. 关闭所有 channel
	for _, ch := range r.pool {
		if ch != nil {
			_ = ch.Close()
		}
	}
	// 2. 关闭连接
	if r.conn != nil {
		return r.conn.Close()
	}
	return nil
}
func (r *RabbitMQ) getChannel() *amqp.Channel {
	idx := atomic.AddUint32(&r.index, 1)
	return r.pool[idx%uint32(len(r.pool))]
}
