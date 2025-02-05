package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/hertz-contrib/requestid"
	"strings"
	"time"
)

type CustomResponseWriter struct {
	network.Writer
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.WriteString(s)
}

// RequestIdAuth requestId中间件
func RequestIdAuth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 开始时间
		reqStartTime := time.Now()
		writer := CustomResponseWriter{
			body:   bytes.NewBufferString(""),
			Writer: ctx.GetWriter(),
		}
		// 获取请求参数
		reqBody := getRequestParams(c, ctx)
		// 处理请求
		ctx.Next(c)
		fields := map[string]interface{}{
			"req_body":     reqBody,
			"req_host":     ctx.Request.Host,
			"req_method":   ctx.Request.Method,
			"req_clientIp": ctx.ClientIP(),
			"req_id":       requestid.Get(ctx),
			"req_uri":      ctx.Request.RequestURI,
			"res_time":     fmt.Sprintf("%s", time.Now().Sub(reqStartTime)), // 响应时间
		}

		if ctx.Response.StatusCode() != 200 {
			responseData := writer.body.String()
			fields["req_header"] = ctx.Request.Header
			fields["data"] = responseData
			// 记录request日志
			hlog.Warn(fields)
		} else {
			hlog.Info(fields)
		}
	}
}

// getRequestParams 获取请求参数（GET,POST,PUT,DELETE）等
func getRequestParams(c context.Context, ctx *app.RequestContext) string {
	if string(ctx.Request.Method()) == "GET" {
		var params []string
		ctx.QueryArgs().VisitAll(func(key, value []byte) {
			params = append(params, string(key)+"="+string(value[0]))
		})
		return strings.Join(params, "&")
	}
	if string(ctx.ContentType()) != "application/json" {
		return ctx.PostArgs().String()
	}
	rawData := ctx.GetRawData()
	//读取后，重新赋值 c.Request.Body ，以供后续的其他操作
	//ctx.Request.SetBody(rawData)
	var m map[string]string
	var params []string
	// 反序列化
	err := json.Unmarshal(rawData, &m)
	if err != nil {
		return ""
	}
	for key, value := range m {
		params = append(params, key+"="+value)
	}
	return strings.Join(params, "&")
}
