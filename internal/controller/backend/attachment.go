package backend

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-mogu/hz-framework/global"
	"github.com/go-mogu/hz-framework/internal/controller/base"
	"github.com/go-mogu/hz-framework/pkg/response"
	"github.com/go-mogu/hz-framework/pkg/util"
	"github.com/go-mogu/hz-framework/types/attachment"
)

type AttachmentController struct {
	base.Controller
}

var Attachment = AttachmentController{}

// Upload 上传图片
func (c *AttachmentController) Upload(context context.Context, ctx *app.RequestContext) {
	var requestParams attachment.UploadRequest
	if err := ctx.BindAndValidate(requestParams); err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	file, err := util.UploadFile(global.Cfg.Server.FileUploadPath+requestParams.FilePath, ctx)
	if err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	response.SuccessJson(ctx, "", file)
}
