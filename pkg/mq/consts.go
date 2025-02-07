package mq

import (
	"github.com/pkg/errors"
	"time"
)

const (
	// 连接失败后重新连接服务器时间间隔
	reconnectDelay = 5 * time.Second

	// 建立通道时出现通道异常时间间隔
	reInitDelay = 2 * time.Second

	// 重新发送消息时，服务器没有确认时间间隔
	resendDelay = 5 * time.Second
)

var (
	// 交换机连接方式
	exchangeTypeList = []string{"topic", "direct", "fanout", "headers"}

	errNotConnected  = errors.New("not connected to a server")
	errAlreadyClosed = errors.New("already closed: not connected to the server")
	errShutdown      = errors.New("session is shutting down")
	errFailedToPush  = errors.New("failed to push: not connected")
)
