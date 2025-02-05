package frontend

import (
	"github.com/go-mogu/hz-framework/internal/controller/base"
)

type UserController struct {
	*base.Controller
}

var User = &UserController{}
