package zap

import (
	"github.com/go-mogu/hz-framework/global"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var FileRotateLogs = new(fileRotateLogs)

type fileRotateLogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
// Author [SliverHorn](https://github.com/SliverHorn)
func (r *fileRotateLogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotateLogs.New(
		path.Join(global.Cfg.Zap.Director, "%Y-%m-%d", level+".log"),
		rotateLogs.WithClock(rotateLogs.Local),
		rotateLogs.WithMaxAge(time.Duration(global.Cfg.Zap.MaxAge)*24*time.Hour), // 日志留存时间
		rotateLogs.WithRotationTime(time.Hour*24),
	)
	if global.Cfg.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
