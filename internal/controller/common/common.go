package common

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-mogu/hz-framework/internal/controller/base"
	"github.com/go-mogu/hz-framework/internal/service/common"
	"github.com/go-mogu/hz-framework/pkg/response"
	common2 "github.com/go-mogu/hz-framework/types/common"
)

type CommonController struct {
	*base.Controller
}

var Common = &CommonController{}

// Ping 心跳
func (c *CommonController) Ping(context context.Context, ctx *app.RequestContext) {
	response.SuccessJson(ctx, "Pong!", "")
}

// Routes 获取所有路由
func (c *CommonController) Routes(context context.Context, ctx *app.RequestContext) {
	var requestParams common2.RouteRequest
	if err := ctx.BindAndValidate(&requestParams); err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	casbinRuleList, err := common.Common.AddRoutes(requestParams)
	if err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	response.SuccessJson(ctx, "", casbinRuleList)
}
