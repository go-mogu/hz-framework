package backend

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-mogu/hz-framework/internal/controller/base"
	"github.com/go-mogu/hz-framework/internal/service/backend"
	"github.com/go-mogu/hz-framework/pkg/response"
	"github.com/go-mogu/hz-framework/types/user"
)

type UserController struct {
	base.Controller
}

var User = UserController{}

// Index 获取列表
func (c *UserController) Index(context context.Context, ctx *app.RequestContext) {
	var requestParams user.IndexRequest
	if err := ctx.BindAndValidate(&requestParams); err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	list, err := backend.User.GetIndex(requestParams)
	if err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	response.SuccessJson(ctx, "", list)
}

// List 获取列表
func (c *UserController) List(context context.Context, ctx *app.RequestContext) {
	var requestParams user.IndexRequest
	if err := ctx.BindAndValidate(&requestParams); err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	list, err := backend.User.GetList(requestParams)
	if err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	response.SuccessJson(ctx, "", list)
}
