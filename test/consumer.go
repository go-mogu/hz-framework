package main

import (
	"github.com/go-mogu/hz-framework/bootstrap"
	"github.com/go-mogu/hz-framework/config"
	"github.com/go-mogu/hz-framework/global"
	"github.com/go-mogu/hz-framework/internal/amqp/consumer"
	"github.com/go-mogu/hz-framework/internal/service/backend"
	"github.com/go-mogu/hz-framework/pkg/mq"
	"time"
)

func main() {
	config.ConfEnv = "dev"
	bootstrap.BootService(bootstrap.LoggerService)
	amqpConfig := &mq.Config{
		User:     global.Cfg.Amqp.User,
		Password: global.Cfg.Amqp.Password,
		Host:     global.Cfg.Amqp.Host,
		Port:     global.Cfg.Amqp.Port,
		Vhost:    global.Cfg.Amqp.Vhost,
	}
	// 实例化amqp
	amqp := mq.New(amqpConfig, "mogu.email", "exchange.direct", "mogu.email", 1, 1, true)
	time.Sleep(time.Second * 1)
	// 启动3个消费者
	data := map[string]interface{}{
		"consumerNum": 1,
	}
	if err := consumer.New(amqp, data, backend.User.AmqpConsumerHandler).Consumer(); err != nil {
	}
}
