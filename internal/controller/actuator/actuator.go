package actuator

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-mogu/hz-framework/pkg/response"
)

type cActuatorController struct {
}

var Actuator = &cActuatorController{}

// Ping 心跳
func (c *cActuatorController) Ping(context context.Context, ctx *app.RequestContext) {
	hlog.Info(string(ctx.Response.Header.Header()))
	response.SuccessJson(ctx, "Pong!", "")
}
