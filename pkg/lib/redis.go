package lib

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Addr     string
	Password string
	DbNum    int
}

// NewRedis redis连接
func NewRedis(config RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DbNum,
	})
	err := client.Ping(context.Background()).Err()
	return client, err
}
