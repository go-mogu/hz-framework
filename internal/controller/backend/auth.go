package backend

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-mogu/hz-framework/internal/controller/base"
	"github.com/go-mogu/hz-framework/internal/service/common"
	"github.com/go-mogu/hz-framework/pkg/response"
	"github.com/go-mogu/hz-framework/types/user"
)

type AuthController struct {
	base.Controller
}

var Auth = AuthController{}

// Login 用户登录
func (c *AuthController) Login(context context.Context, ctx *app.RequestContext) {
	var requestParams user.LoginRequest
	if err := ctx.BindAndValidate(&requestParams); err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	token, err := common.User.Login(requestParams)
	if err != nil {
		response.BadRequestException(ctx, "")
		return
	}
	response.SuccessJson(ctx, "", token)
}
