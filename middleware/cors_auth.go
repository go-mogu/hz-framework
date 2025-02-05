package middleware

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
	"net/http"
)

// CorsAuth 跨域中间件
func CorsAuth() app.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch, http.MethodHead, http.MethodConnect, http.MethodOptions, http.MethodTrace}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "User-Agent", "Cookie", "Authorization", "X-Auth-Token", "X-Requested-With"}
	config.AllowOrigins = []string{"*"}
	config.MaxAge = 3628800
	config.AllowWildcard = true
	config.AllowBrowserExtensions = true
	config.AllowWebSockets = true
	return cors.New(config)
}
