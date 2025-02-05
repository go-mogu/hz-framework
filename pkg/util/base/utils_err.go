package utils

import "github.com/cloudwego/hertz/pkg/common/hlog"

func ErrIsNil(err error, msg ...string) {
	if !IsNil(err) {
		hlog.Error(err.Error())
		if len(msg) > 0 {
			panic(msg[0])
		} else {
			panic(err.Error())
		}
	}
}

func ValueIsNil(value interface{}, msg string) {
	if IsNil(value) {
		panic(msg)
	}
}
