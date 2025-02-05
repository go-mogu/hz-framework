package router

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/go-mogu/hz-framework/global"
	"github.com/go-mogu/hz-framework/middleware"
	baseClient "github.com/go-mogu/hz-framework/pkg/client"
	"github.com/go-mogu/hz-framework/pkg/response"
	"github.com/go-mogu/hz-framework/pkg/util"
	"github.com/go-mogu/hz-framework/router/routes"
	"github.com/go-mogu/mogu-registry/nacos"
	"github.com/hertz-contrib/requestid"
	"github.com/hertz-contrib/swagger"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	swaggerFiles "github.com/swaggo/files"
)

func Register(port string) *server.Hertz {
	//获取本机ip
	addr := util.GetIpAddr()
	//nacos服务发现客户端
	nacosCli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &global.Cfg.Nacos.Client,
			ServerConfigs: global.Cfg.Nacos.Server,
		})
	if err != nil {
		panic(err)
	}
	if global.Cfg.Server.Port == "" {
		global.Cfg.Server.Port = port
	}
	addr = addr + ":" + port
	//注册服务
	r := nacos.NewNacosRegistry(nacosCli)
	bindConfig := binding.NewBindConfig()
	bindConfig.LooseZeroMode = true
	h := server.New(
		server.WithHostPorts("0.0.0.0"+":"+port),
		server.WithBindConfig(bindConfig),
		server.WithRegistry(r, &registry.Info{
			ServiceName: global.Cfg.Server.Name,
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      1,
			Tags:        global.Cfg.Nacos.Discovery.Metadata,
		}),
	)
	url := swagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", port)) // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
	// recovery
	h.Use(recovery.Recovery(recovery.WithRecoveryHandler(response.RecoveryHandler)))
	// cors
	h.Use(middleware.CorsAuth())
	// header add X-Request-Id
	h.Use(requestid.New())
	h.Use(middleware.RequestIdAuth())
	// 404 not found
	h.NoRoute(func(c context.Context, ctx *app.RequestContext) {
		path := ctx.Request.URI().Path()
		method := ctx.Request.Method()
		response.NotFoundException(ctx, fmt.Sprintf("%s %s not found", method, path))
	})
	//健康检查路由
	routes.InitActuatorGroup(h.Group("/"))
	// 路由分组
	var (
		publicMiddleware = []app.HandlerFunc{
			middleware.IpAuth(),
		}
		commonGroup = h.Group("/", publicMiddleware...)
		authGroup   = h.Group("/", append(publicMiddleware, middleware.LoginAuth(), middleware.CasbinAuth())...)
	)
	// 公用组
	routes.InitCommonGroup(commonGroup)
	// 后台组
	routes.InitBackendGroup(authGroup)
	// 前台组
	routes.InitFrontendGroup(authGroup)
	// 初始化客户端
	baseClient.InitClient()
	// 赋给全局
	global.Router = h
	return h
}
