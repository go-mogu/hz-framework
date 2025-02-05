package backend

import (
	"fmt"
	"github.com/go-mogu/hz-framework/global"
	"github.com/go-mogu/hz-framework/models"
	goRabbitmq "github.com/go-mogu/hz-framework/pkg/mq"
	"github.com/go-mogu/hz-framework/pkg/paginator"
	"github.com/go-mogu/hz-framework/pkg/util/gconv"
	"github.com/go-mogu/hz-framework/types/user"
	"sync"
)

type UserService struct{}

var User = &UserService{}

// GetIndex 获取列表
func (s *UserService) GetIndex(requestParams user.IndexRequest) (interface{}, error) {
	multiFields := []paginator.SelectTableField{
		{Model: models.SysUser{}, Table: models.SysUserTbName, Field: []string{"password", "salt", "_omit"}},
		{Model: models.SysUserInfo{}, Table: models.SysUserInfoTbName, Field: []string{"id", "user_id", "role_ids"}},
	}
	pagination, err := paginator.NewBuilder[user.UserList]().
		WithDB(global.DB).
		WithModel(models.SysUser{}).
		//WithFields(models.SysUser{}, models.SysUserTbName, []string{"password", "salt", "_omit"}).
		//WithFields(models.SysUserInfo{}, models.SysUserInfoTbName, []string{"id", "user_id", "role_ids"}).
		WithMultiFields(multiFields).
		WithJoins("left", []paginator.OnJoins{{
			LeftTableField:  paginator.JoinTableField{Table: models.SysUserTbName, Field: "id"},
			RightTableField: paginator.JoinTableField{Table: models.SysUserInfoTbName, Field: "user_id"},
		}}).
		Pagination(requestParams.Page, global.Cfg.Server.DefaultPageSize)
	return pagination, err
}

// GetList 获取列表
func (s *UserService) GetList(requestParams user.IndexRequest) (interface{}, error) {
	pagination, err := paginator.NewBuilder[user.UserList]().
		WithDB(global.DB).
		WithModel(models.SysUser{}).
		WithPreload("UserInfo").
		Pagination(requestParams.Page, global.Cfg.Server.DefaultPageSize)
	return pagination, err
}

// AmqpConsumerHandler 处理消费者方法
func (s *UserService) AmqpConsumerHandler(mq *goRabbitmq.RabbitMQ, data map[string]interface{}) error {
	var wg sync.WaitGroup
	chErrors := make(chan error)
	consumerNum := gconv.Int(data["consumerNum"])
	wg.Add(consumerNum)
	for i := 0; i < consumerNum; i++ {
		fmt.Printf("正在开启消费者：第 %d 个\n", i+1)
		go func() {
			defer wg.Done()
			deliveries, err := mq.Consume()
			if err != nil {
				chErrors <- err
			}
			for d := range deliveries {
				// 消费者逻辑 to do
				fmt.Printf("got %dbyte delivery: [%v] %s %q\n", len(d.Body), d.DeliveryTag, d.Exchange, d.Body)
				d.Ack(false)
			}
		}()
	}
	select {
	case err := <-chErrors:
		close(chErrors)
		fmt.Printf("Consumer failed: %s\n", err)
		return err
	}
	wg.Wait()
	return nil
}

// AmqpProducerHandler 处理生产者方法
func (s *UserService) AmqpProducerHandler(mq *goRabbitmq.RabbitMQ, data []byte) error {
	if err := mq.Push(data); err != nil {
		fmt.Println("Push failed: " + err.Error())
		return err
	}
	fmt.Println("Push succeeded!", string(data))
	return nil
}
