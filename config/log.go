package config

import "github.com/feilongjump/jigsaw-api/plugins/viper_lib"

type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

type Log struct {
	Level      LogLevel
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func GetLogConfig() *Log {
	return &Log{
		// 日志级别
		Level: LogLevel(viper_lib.GetConfig().GetString("log.level")),
		// 日志文件名
		Filename: viper_lib.GetConfig().GetString("log.filename"),
		// 单个文件最大 xMB
		MaxSize: viper_lib.GetConfig().GetInt("log.max_size"),
		// 多于 x 个日志文件后，清理较旧的日志
		MaxBackups: viper_lib.GetConfig().GetInt("log.max_backups"),
		// x 天一切割
		MaxAge: viper_lib.GetConfig().GetInt("log.max_age"),
		// 是否压缩
		Compress: viper_lib.GetConfig().GetBool("log.compress"),
	}
}
