package logger

import (
	"github.com/feilongjump/jigsaw-api/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

var logger *zap.Logger

func Initialize() {

	encoder := getEncoder()

	// 获取日志级别
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(config.GetLogConfig().Level)); err != nil {
		panic(err)
	}

	core := zapcore.NewCore(encoder, getLogWriter(), logLevel)

	logger = zap.New(
		core,
		zap.AddCaller(),                   // 调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1),              // 封装了一层，调用文件去除一层(runtime.Caller(1))
		zap.AddStacktrace(zap.ErrorLevel), // Error 时才会显示 stacktrace
	)
}

// 获取日志编码器
func getEncoder() zapcore.Encoder {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder

	// 根据环境选择日志输出
	if config.IsLocal() {
		// 本地环境，高亮输出到控制台
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 自定义时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// 获取日志写入器
func getLogWriter() zapcore.WriteSyncer {
	logConfig := config.GetLogConfig()

	logname := time.Now().Format("2006-01-02.log")
	filename := strings.ReplaceAll(logConfig.Filename, "logs.log", logname)

	// 滚动日志
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    logConfig.MaxSize,
		MaxBackups: logConfig.MaxBackups,
		MaxAge:     logConfig.MaxAge,
		Compress:   logConfig.Compress,
	}

	if config.IsLocal() {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}

	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger))
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}
