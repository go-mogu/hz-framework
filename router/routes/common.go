package routes

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/go-mogu/hz-framework/internal/controller/backend"
	"github.com/go-mogu/hz-framework/internal/controller/common"
)

func InitCommonGroup(r *route.RouterGroup) (router route.IRoutes) {
	commonGroup := r.Group("")
	{
		// ping
		commonGroup.GET("/ping", common.Common.Ping)
		// 默认给超级管理员角色
		commonGroup.GET("/routes", common.Common.Routes)

		// 生成token
		commonGroup.GET("/token/create", common.Token.Create)
		// 解析token
		commonGroup.POST("/token/view", common.Token.View)
		// 登录
		commonGroup.POST("/user/login", backend.Auth.Login)
		// 获取用户列表
		commonGroup.GET("/user/index", backend.User.Index)
		// 获取用户列表
		commonGroup.GET("/user/list", backend.User.List)
		// 上传附件
		commonGroup.POST("/attachment/upload", backend.Attachment.Upload)
	}
	return commonGroup
}
