package logger

import (
	"os"
	"path/filepath"
	"time"

	"github.com/feilongjump/jigsaw-api/infrastructure/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func Init() {
	loggerConfig := config.Global.Logger

	level, err := zapcore.ParseLevel(loggerConfig.Level)
	if err != nil {
		// 解析日志级别失败时，默认设置为 InfoLevel
		level = zapcore.InfoLevel
	}

	// 区分生产环境和开发环境
	var core zapcore.Core
	if config.IsProd() {
		// 生产环境下，将日志写入文件
		core = newProdCore(level, loggerConfig)
	} else {
		// 开发环境下，控制台打印
		core = newDevCore(level)
	}

	// 初始化 Logger
	Logger = zap.New(
		core,
		// 记录日志调用位置（文件、行号、函数名）
		zap.AddCaller(),
		// 跳过调用者栈帧，避免记录 logger 包内部调用
		zap.AddCallerSkip(1),
		// 错误级别以上添加堆栈信息
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
}

// newProdCore 生产环境下，将日志写入文件
func newProdCore(level zapcore.Level, loggerConfig config.LoggerConfig) zapcore.Core {

	// 日志文件名
	logname := time.Now().Format("2006-01-02.log")
	filename := filepath.Join(loggerConfig.FilePath, logname)

	// 日志文件切割配置
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    loggerConfig.MaxSize,
		MaxBackups: loggerConfig.MaxBackups,
		MaxAge:     loggerConfig.MaxAge,
		Compress:   loggerConfig.Compress,
	}

	// 生产环境下，使用生产模式的编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间 key
	encoderConfig.TimeKey = "time"
	// 自定义时间格式
	encoderConfig.EncodeTime = customTimeEncoder

	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		// 指定输出目标到文件
		zapcore.AddSync(lumberjackLogger),
		// 设置日志级别
		level,
	)
}

// newDevCore 开发环境下，控制台打印
func newDevCore(level zapcore.Level) zapcore.Core {

	// 开发环境下，使用开发模式的编码器配置
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	// 自定义时间格式
	encoderConfig.EncodeTime = customTimeEncoder
	// 打印带彩色的日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// 构建核心
	return zapcore.NewCore(
		// 自定义控制台编码器
		zapcore.NewConsoleEncoder(encoderConfig),
		// 指定输出目标到控制台
		zapcore.AddSync(os.Stdout),
		// 设置日志级别
		level,
	)
}

// 自定义时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// Sync 刷新日志缓冲区（程序退出时调用）
func Sync() error {
	return Logger.Sync()
}

// Debug 开发调试、临时日志、详细流程跟踪
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

// Info 服务启动、操作成功、重要状态变更
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

// Warn 缓存失败、参数无效、资源阈值预警
func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

// Error 业务操作失败、接口调用失败、数据处理错误
func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

// DPanic 程序逻辑错误、理论上不该发生的情况，在生产环境下仅记录错误，不崩溃
func DPanic(msg string, fields ...zap.Field) {
	Logger.DPanic(msg, fields...)
}

// Panic 用于记录当前 goroutine 无法恢复的严重错误，如数据库连接失败，服务无法提供功能，触发 Panic
func Panic(msg string, fields ...zap.Field) {
	Logger.Panic(msg, fields...)
}

// Fatal 最高级别的日志，用于服务启动阶段的致命错误（如监听端口失败、核心配置缺失）
func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}
