package response

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-mogu/hz-framework/pkg/util/gconv"
)

func RecoveryHandler(c context.Context, ctx *app.RequestContext, err interface{}, stack []byte) {
	if ctx.GetResponse().BodyBuffer().Len() > 0 && err == nil {
		return
	}
	if err != nil {
		msg := gconv.String(err)
		FailedJson(ctx, msg, string(stack))
	}
}
