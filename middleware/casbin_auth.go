package middleware

import (
	"context"
	_ "embed"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/util"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-mogu/hz-framework/global"
	gApp "github.com/go-mogu/hz-framework/global/app"
	"github.com/go-mogu/hz-framework/pkg/response"
	util2 "github.com/go-mogu/hz-framework/pkg/util"
	"github.com/go-mogu/hz-framework/pkg/util/gconv"
	"strings"
)

//go:embed rbac_model.conf
var rbacModelConf string

// CasbinAuth 用户权限验证
func CasbinAuth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		adapter, _ := gormAdapter.NewAdapterByDB(global.DB)
		rc, err := model.NewModelFromString(rbacModelConf)
		if err != nil {
			ctx.Abort()
			return
		}
		e, _ := casbin.NewEnforcer(rc, adapter)
		e.AddFunction("ParamsMatch", ParamsMatchFunc)
		e.AddFunction("ParamsActMatch", ParamsActMatchFunc)
		_ = e.LoadPolicy()

		//	获取当前请求的url
		obj := ctx.Request.RequestURI()
		act := ctx.Request.Method
		user, err := gApp.GetAdminInfo(c, ctx)
		if err != nil {
			response.UnauthorizedException(ctx, "权限异常")
			ctx.Abort()
			return
		}
		var flag = false
		for _, sub := range user.RoleIds {
			//	判断策略中是否存在
			subStr := gconv.String(sub)
			if ok, _ := e.Enforce(subStr, obj, act); ok {
				flag = true
				break
			}
		}
		if !flag {
			response.UnauthorizedException(ctx, "该用户无此权限")
			ctx.Abort()
			return
		}
		ctx.Next(c)
	}
}

// ParamsActMatchFunc 自定义规则函数
func ParamsActMatchFunc(args ...interface{}) (interface{}, error) {
	rAct := args[0].(string)
	pAct := args[1].(string)
	pActArr := strings.Split(pAct, ",")
	return util2.InAnySlice[string](pActArr, rAct), nil
}

// ParamsMatchFunc 自定义规则函数
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	key1 := strings.Split(name1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, name2), nil
}
