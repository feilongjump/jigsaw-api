package cmd

import (
	"fmt"
	"github.com/feilongjump/jigsaw-api/plugins/database"
	"github.com/feilongjump/jigsaw-api/plugins/logger"
	"github.com/feilongjump/jigsaw-api/plugins/viper_lib"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:  "jigsaw",
	Long: "生活就像拼图，每一块都有意义。",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 优先运行
		// 初始化配置
		viper_lib.LoadConfig(env)
		// 初始化日志
		logger.Initialize()
		// 初始化数据库连接
		database.Initialize()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// 数据库自动迁移
		runMigration(cmd, args)
		// 运行 web 服务
		runServer(cmd, args)
	},
}

var env string

func init() {
	rootCmd.PersistentFlags().StringVarP(&env, "env", "e", "", "加载环境变量文件，默认加载：.jigsaw.local.toml")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
