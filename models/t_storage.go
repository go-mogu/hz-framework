package models

import "time"

var TStorageTbName = "t_storage"

// TStorage 存储信息表
type TStorage struct {
	Uid            string    `gorm:"primaryKey;column:uid;type:varchar(32);NOT NULL;comment:唯一uid" json:"uid"`                               // 唯一uid
	AdminUid       string    `gorm:"column:admin_uid;type:varchar(32);NOT NULL;comment:管理员uid" json:"admin_uid"`                             // 管理员uid
	StorageSize    int64     `gorm:"column:storage_size;type:bigint(20);default:0;NOT NULL;comment:网盘容量大小" json:"storage_size"`              // 网盘容量大小
	Status         int8      `gorm:"column:status;type:tinyint(1) unsigned;default:1;NOT NULL;comment:状态" json:"status"`                     // 状态
	CreateTime     time.Time `gorm:"column:create_time;type:timestamp;default:0000-00-00 00:00:00;NOT NULL;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime     time.Time `gorm:"column:update_time;type:timestamp;default:0000-00-00 00:00:00;NOT NULL;comment:更新时间" json:"update_time"` // 更新时间
	MaxStorageSize *int64    `gorm:"column:max_storage_size;type:bigint(20);NULL;comment:最大容量大小" json:"max_storage_size"`                    // 最大容量大小
}
