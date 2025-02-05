package client

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/go-mogu/hz-framework/global"
	utils "github.com/go-mogu/hz-framework/pkg/util/base"
	"github.com/go-mogu/mogu-registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"net/http"
	"net/url"
)

var (
	scheme     = "http"
	BaseClient *client.Client
)

func InitClient() {
	nacosCli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &global.Cfg.Nacos.Client,
			ServerConfigs: global.Cfg.Nacos.Server,
		})
	utils.ErrIsNil(err)
	r := nacos.NewNacosResolver(nacosCli)
	baseClient, err := client.NewClient()
	utils.ErrIsNil(err)
	baseClient.Use(sd.Discovery(r))
	BaseClient = baseClient
}

func Get(ctx context.Context, serviceName, url string, data url.Values) (result []byte, err error) {
	if data != nil {
		url = url + "?" + data.Encode()
	}
	url = fmt.Sprintf("%s://%s%s", scheme, serviceName, url)
	_, result, err = BaseClient.Get(ctx, nil, url, config.WithSD(true))
	if err != nil {
		return nil, err
	}
	return
}

func Post(ctx context.Context, serviceName, url string, query url.Values, param map[string]string, body []byte) (result []byte, err error) {
	if query != nil {
		url = url + "?" + query.Encode()
	}
	url = fmt.Sprintf("%s://%s%s", scheme, serviceName, url)
	request := protocol.AcquireRequest()
	response := protocol.AcquireResponse()
	defer func(response *protocol.Response) {
		err = response.CloseBodyStream()
		if err != nil {
			hlog.Error(err)
		}
	}(response)
	request.SetOptions(config.WithSD(true))
	request.SetMethod(http.MethodPost)
	request.SetRequestURI(url)
	request.ParseURI()
	if param != nil {
		request.SetFormData(param)
	}
	if body != nil {
		request.SetBody(body)
	}

	err = BaseClient.Do(ctx, request, response)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}
	result = response.Body()
	return
}
