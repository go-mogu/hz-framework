package common

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-mogu/hz-framework/global"
	"github.com/go-mogu/hz-framework/internal/controller/base"
	"github.com/go-mogu/hz-framework/pkg/auth"
	"github.com/go-mogu/hz-framework/pkg/response"
	"strings"
)

type TokenController struct {
	base.Controller
}

var Token = TokenController{}

// Create 生成token
func (c *TokenController) Create(context context.Context, ctx *app.RequestContext) {
	token, err := auth.GenerateJwtToken(global.Cfg.Jwt.Secret, global.Cfg.Jwt.TokenExpire, map[string]interface{}{"id": 1}, global.Cfg.Jwt.TokenIssuer)
	if err != nil {
		response.UnauthorizedException(ctx, err.Error())
		return
	}
	response.SuccessJson(ctx, "", token)
}

// View token解析
func (c *TokenController) View(context context.Context, ctx *app.RequestContext) {
	token := string(ctx.GetHeader(global.Cfg.Jwt.TokenKey))
	if token == "" {
		response.UnauthorizedException(ctx, "")
		return
	}
	flag := strings.Contains(token, "Bearer")
	if !flag {
		response.UnauthorizedException(ctx, "")
		return
	}
	token = strings.TrimSpace(strings.TrimLeft(token, "Bearer"))
	jwtTokenArr, err := auth.ParseJwtToken(token, global.Cfg.Jwt.Secret)
	if err != nil {
		response.UnauthorizedException(ctx, "")
		return
	}
	response.SuccessJson(ctx, "", jwtTokenArr)
}
