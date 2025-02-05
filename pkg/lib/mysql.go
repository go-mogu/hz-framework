package lib

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-mogu/hz-framework/pkg/consts/EStatus"
	"github.com/go-mogu/hz-framework/pkg/consts/SQLConf"
	utils "github.com/go-mogu/hz-framework/pkg/util/base"
	"github.com/yitter/idgenerator-go/idgen"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type DatabaseConfig struct {
	Host          string
	Port          string
	User          string
	Pass          string
	DbName        string
	Prefix        string
	MaxIdleConnes int
	MaxOpenConnes int
	MaxLifeTime   int // 分钟
	LogLevel      logger.LogLevel
}

// NewMysql 数据库连接
func NewMysql(config DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Pass, config.Host, config.Port, config.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Prefix,
			SingularTable: true, // 是否设置单数表名，设置为 是
		},
		Logger: logger.Default.LogMode(config.LogLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the database, please check the MySQL configuration information first,the error details are:" + err.Error())
	}
	// GORM 使用 database/sql 维护连接池
	sqlDB, _ := db.DB()
	// SetMaxIdleConnes 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(config.MaxIdleConnes)
	// SetMaxOpenConnes 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.MaxOpenConnes)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(config.MaxLifeTime))
	err = db.Callback().Create().Before("gorm:create").Register("gorm:create_meta_object_handler", CreateMetaObjectHandler)
	if err != nil {
		return nil, err
	}
	err = db.Callback().Update().Before("gorm:update").Register("gorm:update_meta_object_handler", UpdateMetaObjectHandler)
	return db, err
}

// NewNullString 空字符串
func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// NewNullInt64 空数值
func NewNullInt64(s int64) sql.NullInt64 {
	if s == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: s,
		Valid: true,
	}
}

func CreateMetaObjectHandler(db *gorm.DB) {
	fields := db.Statement.Schema.Fields
	now := time.Now()
	ctx := context.Background()
	for _, field := range fields {
		//值为空时填充默认值
		if _, isZero := field.ValueOf(ctx, db.Statement.ReflectValue); isZero {
			switch field.DBName {
			case SQLConf.UID:
				err := field.Set(ctx, db.Statement.ReflectValue, idgen.NextId())
				utils.ErrIsNil(err)
			case SQLConf.CREATE_TIME:
				err := field.Set(ctx, db.Statement.ReflectValue, now)
				utils.ErrIsNil(err)
			case SQLConf.UPDATE_TIME:
				err := field.Set(ctx, db.Statement.ReflectValue, time.Now())
				utils.ErrIsNil(err)
			case SQLConf.STATUS:
				err := field.Set(ctx, db.Statement.ReflectValue, EStatus.ENABLE)
				utils.ErrIsNil(err)
			}
		}

	}
}
func UpdateMetaObjectHandler(db *gorm.DB) {
	fields := db.Statement.Schema.Fields
	ctx := context.Background()
	for _, field := range fields {
		//值为空时填充默认值
		if _, isZero := field.ValueOf(ctx, db.Statement.ReflectValue); isZero {
			switch field.DBName {
			case SQLConf.UPDATE_TIME:
				err := field.Set(ctx, db.Statement.ReflectValue, time.Now())
				utils.ErrIsNil(err)

			}
		}

	}
}
