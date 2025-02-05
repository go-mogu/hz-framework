package admin

import "github.com/go-mogu/hz-framework/models"

type BaseAdmin models.SysAdmin

type Admin struct {
	BaseAdmin
	RoleIds []uint64 `gorm:"-" json:"role_ids"`
}
