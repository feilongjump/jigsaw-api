package viper_lib

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
	"os"
)

var viperConfig *viper.Viper

func init() {
	viperConfig = viper.New()

	viperConfig.AddConfigPath(".")
	viperConfig.SetConfigType("toml")
}

func LoadConfig(env string) {
	// 加载 .jigsaw.toml 文件
	envPath := ".jigsaw"
	if len(env) > 0 {
		envPath = envPath + "." + env
	}
	envPath = envPath + ".toml"

	// 检查目标配置文件是否存在
	// 当文件不存在时，自动生成配置文件
	if _, err := os.Stat(envPath); err != nil {
		copyExampleConfigToTarget(envPath)
	}

	viperConfig.SetConfigName(envPath)
	if err := viperConfig.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件出错")
		panic(err)
	}

	viperConfig.WatchConfig()
}

// 复制 .jigsaw.example.toml 文件到目标文件
// 不希望破坏 example 文件格式，所以没有使用 viper 库的生成配置文件方式
func copyExampleConfigToTarget(envPath string) {
	// 先读取 .jigsaw.example.toml 文件
	sourceFile := ".jigsaw.example.toml"
	source, err := os.Open(sourceFile)
	if err != nil {
		fmt.Printf("打开源文件 %s 出错\n", sourceFile)
		panic(err)
	}
	defer source.Close()

	// 创建配置文件
	target, err := os.Create(envPath)
	if err != nil {
		fmt.Printf("创建配置文件 %s 出错\n", envPath)
		panic(err)
	}
	defer target.Close()

	if _, err := io.Copy(target, source); err != nil {
		fmt.Println("生成配置信息错误")
		panic(err)
	}
}

func GetConfig() *viper.Viper {
	return viperConfig
}
