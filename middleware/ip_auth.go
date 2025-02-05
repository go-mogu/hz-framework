package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-mogu/hz-framework/config"
	"net/http"
)

// IpAuth 白名单验证
func IpAuth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		clientIp := ctx.ClientIP()
		flag := false
		for _, value := range config.Whitelist {
			if value == "*" || clientIp == value {
				flag = true
				break
			}
		}
		if !flag {
			ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("%s 不在ip白名单中", clientIp))
			ctx.Abort()
			return
		}
		ctx.Next(c)
	}
}
