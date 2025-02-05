package global

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/go-mogu/hz-framework/config"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB      // Mysql数据库
	Redis  *redis.Client // redis连接池
	Router *server.Hertz // 路由
	Cfg    *config.Conf  // yaml配置
	Viper  *viper.Viper
)
