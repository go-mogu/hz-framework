package bootstrap

import (
	"bytes"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-mogu/hz-framework/config"
	"github.com/go-mogu/hz-framework/global"
	"github.com/go-mogu/hz-framework/pkg/lib"
	"github.com/go-mogu/hz-framework/pkg/util"
	"github.com/go-mogu/hz-framework/pkg/zap"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/yitter/idgenerator-go/idgen"
)

// 定义服务列表
const (
	LoggerService = `Logger`
	MysqlService  = `Mysql`
	RedisService  = `Redis`
	IdGenerator   = `IdGenerator`
)

type bootServiceMap map[string]func() error

// BootedService 已经加载的服务
var (
	BootedService []string
	err           error
	// serviceMap 程序启动时需要自动加载的服务
	serviceMap = bootServiceMap{
		LoggerService: bootLogger,
		MysqlService:  bootMysql,
		RedisService:  bootRedis,
		IdGenerator:   bootIdGenerator,
	}
)

// BootService 加载服务
func BootService(services ...string) {
	// 初始化配置
	if err = bootConfig(); err != nil {
		panic("初始化config配置失败：" + err.Error())
	}
	if len(services) == 0 {
		services = serviceMap.keys()
	}
	BootedService = make([]string, 0)
	for k, val := range serviceMap {
		if util.InAnySlice[string](services, k) {
			if err := val(); err != nil {
				panic("程序服务启动失败:" + err.Error())
			}
			BootedService = append(BootedService, k)
		}
	}
}

// bootConfig 载入配置
func bootConfig() error {
	global.Cfg, global.Viper, err = config.InitConfig()
	if err == nil && global.Cfg.Nacos.Watch {
		err = ListenConfig()
	}
	return err
}

func ListenConfig() error {
	// 创建动态配置客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &global.Cfg.Nacos.Client,
			ServerConfigs: global.Cfg.Nacos.Server,
		},
	)
	global.Cfg.Nacos.Config.OnChange = func(namespace, group, dataId, data string) {
		hlog.Debug("group:" + group + ", dataId:" + dataId + ", configure to change!")
		v := global.Viper
		err = v.ReadConfig(bytes.NewBuffer([]byte(data)))
		if err != nil {
			hlog.Error(err)
			return
		}
		if err := v.Unmarshal(&global.Cfg); err != nil {
			hlog.Error(err)
			return
		}
	}
	err = configClient.ListenConfig(global.Cfg.Nacos.Config)
	return err
}

// bootLogger 将配置载入日志服务
func bootLogger() error {
	logger := zap.NewLogger(global.Cfg.Zap.Director)
	defer logger.Sync()
	hlog.SetLogger(logger)
	hlog.Infof("程序载入Logger服务成功 [ 日志路径：%s ]", global.Cfg.Zap.Director)
	return err
}

// bootMysql 装配数据库连接
func bootMysql() error {
	if global.DB != nil {
		return nil
	}
	dbConfig := lib.DatabaseConfig{
		Host:          global.Cfg.Mysql[0].Host,
		Port:          global.Cfg.Mysql[0].Port,
		User:          global.Cfg.Mysql[0].User,
		Pass:          global.Cfg.Mysql[0].Password,
		DbName:        global.Cfg.Mysql[0].DbName,
		Prefix:        global.Cfg.Mysql[0].Prefix,
		MaxIdleConnes: global.Cfg.Mysql[0].MaxIdleConns,
		MaxOpenConnes: global.Cfg.Mysql[0].MaxOpenConns,
		MaxLifeTime:   global.Cfg.Mysql[0].MaxLifeTime,
	}
	global.DB, err = lib.NewMysql(dbConfig)
	if err == nil {
		hlog.Info("程序载入MySQL服务成功")
	}
	return err
}

// bootRedis 装配redis服务
func bootRedis() error {
	redisConfig := lib.RedisConfig{
		Addr:     fmt.Sprintf("%s:%s", global.Cfg.Redis.Host, global.Cfg.Redis.Port),
		Password: global.Cfg.Redis.Password,
		DbNum:    global.Cfg.Redis.DbNum,
	}
	global.Redis, err = lib.NewRedis(redisConfig)
	if err == nil {
		hlog.Info("程序载入Redis服务成功")
	}
	return err
}

// bootIdGenerator 装配雪花id生成器
func bootIdGenerator() error {
	// 创建 IdGeneratorOptions 对象，请在构造函数中输入 WorkerId：
	var options = idgen.NewIdGeneratorOptions(1)
	// options.WorkerIdBitLength = 10 // WorkerIdBitLength 默认值6，支持的 WorkerId 最大值为2^6-1，若 WorkerId 超过64，可设置更大的 WorkerIdBitLength
	// ...... 其它参数设置参考 IdGeneratorOptions 定义，一般来说，只要再设置 WorkerIdBitLength （决定 WorkerId 的最大值）。

	// 保存参数（必须的操作，否则以上设置都不能生效）：
	idgen.SetIdGenerator(options)
	return nil
}

// keys 获取BootServiceMap中所有键值
func (m bootServiceMap) keys() []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
