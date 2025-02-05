package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/requestid"
	"net/http"
)

type JsonResponse struct {
	Code      string      `json:"code"`
	RequestId string      `json:"requestId"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

// ResultJson 基础返回
func ResultJson(ctx *app.RequestContext, status int, resultCode Code, message string, data interface{}) {
	if message == "" {
		message = CodeMap[resultCode]
	}
	if status == 0 {
		status = http.StatusOK
	}
	ctx.JSON(status, JsonResponse{
		Code:      string(resultCode),
		Message:   message,
		RequestId: requestid.Get(ctx),
		Data:      data,
	})
}

// SuccessJson 成功返回
func SuccessJson(ctx *app.RequestContext, message string, data interface{}) {
	if message == "" {
		message = Success.Msg()
	}
	ResultJson(ctx, http.StatusOK, Success, message, data)
}

// FailedJson 失败返回
func FailedJson(ctx *app.RequestContext, message string, data interface{}) {
	if message == "" {
		message = Success.Msg()
	}
	ResultJson(ctx, http.StatusOK, Failed, message, data)
}

// BadRequestException 400错误
func BadRequestException(ctx *app.RequestContext, message string) {
	if message == "" {
		message = CodeMap[RequestParamErr]
	}
	ResultJson(ctx, http.StatusBadRequest, RequestParamErr, message, nil)
}

// UnauthorizedException 401错误
func UnauthorizedException(ctx *app.RequestContext, message string) {
	if message == "" {
		message = CodeMap[UnAuthed]
	}
	ResultJson(ctx, http.StatusUnauthorized, UnAuthed, message, nil)
}

// ForbiddenException 403错误
func ForbiddenException(ctx *app.RequestContext, message string) {
	if message == "" {
		message = CodeMap[Failed]
	}
	ResultJson(ctx, http.StatusForbidden, Failed, message, nil)
}

// NotFoundException 404错误
func NotFoundException(ctx *app.RequestContext, message string) {
	if message == "" {
		message = CodeMap[RequestMethodErr]
	}
	ResultJson(ctx, http.StatusNotFound, RequestMethodErr, message, nil)
}

// InternalServerException 500错误
func InternalServerException(ctx *app.RequestContext, message string) {
	if message == "" {
		message = CodeMap[InternalErr]
	}
	ResultJson(ctx, http.StatusInternalServerError, InternalErr, message, nil)
}
