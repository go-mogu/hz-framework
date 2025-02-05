package models

import "time"

var TFileTbName = "t_file"

// TFile 文件表
type TFile struct {
	Uid             string    `gorm:"primaryKey;column:uid;type:varchar(36);NOT NULL;comment:唯一uid" json:"uid"`                               // 唯一uid
	FileOldName     *string   `gorm:"column:file_old_name;type:varchar(255);NULL;comment:旧文件名" json:"file_old_name"`                          // 旧文件名
	PicName         *string   `gorm:"column:pic_name;type:varchar(255);NULL;comment:文件名" json:"pic_name"`                                     // 文件名
	PicUrl          *string   `gorm:"column:pic_url;type:varchar(255);NULL;comment:文件地址" json:"pic_url"`                                      // 文件地址
	PicExpandedName *string   `gorm:"column:pic_expanded_name;type:varchar(255);NULL;comment:文件扩展名" json:"pic_expanded_name"`                 // 文件扩展名
	FileSize        *int      `gorm:"column:file_size;type:int(11);NULL;comment:文件大小" json:"file_size"`                                       // 文件大小
	FileSortUid     *string   `gorm:"column:file_sort_uid;type:varchar(36);NULL;comment:文件分类uid" json:"file_sort_uid"`                        // 文件分类uid
	AdminUid        *string   `gorm:"column:admin_uid;type:varchar(36);NULL;comment:管理员uid" json:"admin_uid"`                                 // 管理员uid
	UserUid         *string   `gorm:"column:user_uid;type:varchar(36);NULL;comment:用户uid" json:"user_uid"`                                    // 用户uid
	Status          int8      `gorm:"column:status;type:tinyint(3) unsigned;default:1;NOT NULL;comment:状态" json:"status"`                     // 状态
	CreateTime      time.Time `gorm:"column:create_time;type:timestamp;default:0000-00-00 00:00:00;NOT NULL;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime      time.Time `gorm:"column:update_time;type:timestamp;default:0000-00-00 00:00:00;NOT NULL;comment:更新时间" json:"update_time"` // 更新时间
	QiNiuUrl        *string   `gorm:"column:qi_niu_url;type:varchar(255);NULL;comment:七牛云地址" json:"qi_niu_url"`                               // 七牛云地址
	MinioUrl        *string   `gorm:"column:minio_url;type:varchar(255);NULL;comment:Minio文件URL" json:"minio_url"`                            // Minio文件URL
	FileMd5         *string   `gorm:"column:file_md5;type:varchar(255);NULL;comment:文件md5值" json:"file_md5"`                                  // 文件md5值
}
