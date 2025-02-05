package routes

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/go-mogu/hz-framework/internal/controller/backend"
)

// InitBackendGroup 初始化后台接口路由
func InitBackendGroup(r *route.RouterGroup) (router route.IRoutes) {
	backendGroup := r.Group("backend")
	{
		backendGroup.POST("/user/create", backend.User.Create)
		backendGroup.GET("/user/view", backend.User.View)
		backendGroup.POST("/user/update", backend.User.Update)
		backendGroup.POST("/user/delete", backend.User.Delete)
	}
	return backendGroup
}
