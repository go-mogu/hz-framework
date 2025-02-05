package base

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-mogu/hz-framework/pkg/response"
	"net/http"
)

type Controller struct{}

var Base = Controller{}

func (c *Controller) Index(context context.Context, ctx *app.RequestContext) {
	response.ResultJson(ctx, http.StatusOK, response.Success, "index", "")
}

func (c *Controller) Create(context context.Context, ctx *app.RequestContext) {
	response.ResultJson(ctx, http.StatusOK, response.Success, "create", "")
}

func (c *Controller) Delete(context context.Context, ctx *app.RequestContext) {
	response.ResultJson(ctx, http.StatusOK, response.Success, "delete", "")
}

func (c *Controller) Update(context context.Context, ctx *app.RequestContext) {
	response.ResultJson(ctx, http.StatusOK, response.Success, "update", "")
}

func (c *Controller) View(context context.Context, ctx *app.RequestContext) {
	response.ResultJson(ctx, http.StatusOK, response.Success, "view", "")
}
