package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Logger   LoggerConfig   `mapstructure:"logger"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

type AppConfig struct {
	Name     string `mapstructure:"name"`
	Port     string `mapstructure:"port"`
	Env      string `mapstructure:"env"`
	TimeZone string `mapstructure:"time_zone"`
	Locale   string `mapstructure:"locale"`
}

type LoggerConfig struct {
	// debug < info < warn < error < dpanic < panic < fatal
	Level string `mapstructure:"level"`
	// 日志文件名
	FilePath string `mapstructure:"file_path"`
	// 日志文件最大大小，单位 MB
	MaxSize int `mapstructure:"max_size"`
	// 最多保存日志文件的数量
	MaxBackups int `mapstructure:"max_backups"`
	// 日志文件保留的最大天数，0 表示不删
	MaxAge int `mapstructure:"max_age"`
	// 是否压缩旧日志文件
	Compress bool `mapstructure:"compress"`
}

type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Charset  string `mapstructure:"charset"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	TTL    int    `mapstructure:"ttl"` // Token 有效期（秒）
}

var Global Config

func Init() {
	viperLib := viper.New()

	// 单独设置配置文件的路径、文件名、文件类型，以便后续可根据需要读取其他配置文件
	configName := ".jigsaw"
	viperLib.AddConfigPath(".")
	viperLib.SetConfigName(configName)
	viperLib.SetConfigType("toml")

	if err := viperLib.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 开启自动读取系统环境变量（优先级最高）
	viperLib.AutomaticEnv()

	if err := viperLib.Unmarshal(&Global); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 监听配置文件变化
	viperLib.WatchConfig()
}

// IsProd 是否是生产环境
func IsProd() bool {
	return Global.App.Env == "prod"
}
