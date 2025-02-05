package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-mogu/hz-framework/global"
	"github.com/go-mogu/hz-framework/pkg/auth"
	"github.com/go-mogu/hz-framework/pkg/response"
	"strings"
)

// LoginAuth 登录中间件
func LoginAuth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		token := string(ctx.GetHeader(global.Cfg.Jwt.TokenKey))
		if token == "" {
			response.UnauthorizedException(ctx, "")
			ctx.Abort()
			return
		}
		b := strings.Contains(token, "Bearer")
		if !b {
			response.UnauthorizedException(ctx, "")
			ctx.Abort()
			return
		}
		token = strings.TrimSpace(strings.TrimLeft(token, "Bearer"))
		if _, err := auth.ParseJwtToken(token, global.Cfg.Jwt.Secret); err != nil {
			response.UnauthorizedException(ctx, err.Error())
			ctx.Abort()
			return
		}
		ctx.Next(c)
	}
}
