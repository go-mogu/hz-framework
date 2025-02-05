package models

import "time"

var TFileSortTbName = "t_file_sort"

// TFileSort 文件分类表
type TFileSort struct {
	Uid         string    `gorm:"primaryKey;column:uid;type:varchar(36);NOT NULL;comment:唯一uid" json:"uid"`                               // 唯一uid
	ProjectName *string   `gorm:"column:project_name;type:varchar(255);NULL;comment:项目名" json:"project_name"`                             // 项目名
	SortName    *string   `gorm:"column:sort_name;type:varchar(255);NULL;comment:分类名" json:"sort_name"`                                   // 分类名
	Url         *string   `gorm:"column:url;type:varchar(255);NULL;comment:分类路径" json:"url"`                                              // 分类路径
	Status      int8      `gorm:"column:status;type:tinyint(3) unsigned;default:1;NOT NULL;comment:状态" json:"status"`                     // 状态
	CreateTime  time.Time `gorm:"column:create_time;type:timestamp;default:0000-00-00 00:00:00;NOT NULL;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;type:timestamp;default:0000-00-00 00:00:00;NOT NULL;comment:更新时间" json:"update_time"` // 更新时间
}
