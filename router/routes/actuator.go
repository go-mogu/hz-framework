package routes

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/go-mogu/hz-framework/internal/controller/actuator"
)

func InitActuatorGroup(r *route.RouterGroup) {
	actuatorGroup := r.Group("actuator")
	{
		actuatorGroup.GET("/ping", actuator.Actuator.Ping)
	}
}
