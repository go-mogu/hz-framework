package main

import (
	"fmt"
	"github.com/go-mogu/hz-framework/bootstrap"
	"github.com/go-mogu/hz-framework/config"
	"github.com/go-mogu/hz-framework/global"
	"github.com/go-mogu/hz-framework/internal/amqp/producer"
	"github.com/go-mogu/hz-framework/internal/service/backend"
	"github.com/go-mogu/hz-framework/pkg/mq"
	"github.com/go-mogu/hz-framework/pkg/util"
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
	amqp := mq.New(amqpConfig, "test", "", "", 0, 1, true)
	fmt.Println("启动生产者...")
	// 定时器 1s 执行一次生产者
	util.NewTicker(1, func() error {
		data := []byte("{\"hello\":\"world " + time.Now().Format("2006-01-02 15:04:05") + "\"}")
		if err := producer.New(amqp, data, backend.User.AmqpProducerHandler).Publish(); err != nil {
			return err
		}
		return nil
	}).Start()
}
