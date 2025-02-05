package models

import "time"

var TNetworkDiskTbName = "t_network_disk"

// TNetworkDisk 网盘文件表
type TNetworkDisk struct {
	Uid         string    `gorm:"primaryKey;column:uid;type:varchar(32);NOT NULL;comment:唯一uid" json:"uid"`                               // 唯一uid
	AdminUid    string    `gorm:"column:admin_uid;type:varchar(32);NOT NULL;comment:管理员uid" json:"admin_uid"`                             // 管理员uid
	ExtendName  *string   `gorm:"column:extend_name;type:varchar(255);NULL;comment:扩展名" json:"extend_name"`                               // 扩展名
	FileName    *string   `gorm:"column:file_name;type:varchar(255);NULL;comment:文件名" json:"file_name"`                                   // 文件名
	FilePath    *string   `gorm:"column:file_path;type:varchar(255);NULL;comment:文件路径" json:"file_path"`                                  // 文件路径
	FileSize    int64     `gorm:"column:file_size;type:bigint(20);NOT NULL;comment:文件大小" json:"file_size"`                                // 文件大小
	IsDir       int       `gorm:"column:is_dir;type:int(11);NOT NULL;comment:是否目录" json:"is_dir"`                                         // 是否目录
	Status      int8      `gorm:"column:status;type:tinyint(1) unsigned;default:1;NOT NULL;comment:状态" json:"status"`                     // 状态
	CreateTime  time.Time `gorm:"column:create_time;type:timestamp;default:0000-00-00 00:00:00;NOT NULL;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;type:timestamp;default:0000-00-00 00:00:00;NOT NULL;comment:更新时间" json:"update_time"` // 更新时间
	LocalUrl    *string   `gorm:"column:local_url;type:varchar(255);NULL;comment:本地文件URL" json:"local_url"`                               // 本地文件URL
	QiNiuUrl    *string   `gorm:"column:qi_niu_url;type:varchar(255);NULL;comment:七牛云URL" json:"qi_niu_url"`                              // 七牛云URL
	FileOldName *string   `gorm:"column:file_old_name;type:varchar(255);NULL;comment:上传前文件名" json:"file_old_name"`                        // 上传前文件名
	MinioUrl    *string   `gorm:"column:minio_url;type:varchar(255);NULL;comment:Minio文件URL" json:"minio_url"`                            // Minio文件URL
}
