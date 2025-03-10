package producer

import (
	"github.com/go-mogu/hz-framework/pkg/mq"
	"time"
)

type (
	BaseConfig struct {
		Amqp     *mq.RabbitMQ
		Data     []byte
		CallBack Fn
	}
	Fn func(mq *mq.RabbitMQ, data []byte) error
)

// New 实例化
func New(mq *mq.RabbitMQ, data []byte, f Fn) *BaseConfig {
	return &BaseConfig{
		Amqp:     mq,
		Data:     data,
		CallBack: f,
	}
}

// Publish 消费者
func (c *BaseConfig) Publish() error {
	// 防止消费过快 加一秒时间
	time.Sleep(time.Second * 1)
	if err := c.CallBack(c.Amqp, c.Data); err != nil {
		return err
	}
	return nil
}
