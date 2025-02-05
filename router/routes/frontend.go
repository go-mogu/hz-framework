package routes

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/go-mogu/hz-framework/internal/controller/frontend"
)

// InitFrontendGroup 初始化前台接口路由
func InitFrontendGroup(r *route.RouterGroup) (router route.IRoutes) {
	frontendGroup := r.Group("api")
	{
		frontendGroup.POST("/user/create", frontend.User.Create)
		frontendGroup.GET("/user/view", frontend.User.View)
		frontendGroup.POST("/user/update", frontend.User.Update)
		frontendGroup.POST("/user/delete", frontend.User.Delete)
	}
	return frontendGroup
}
