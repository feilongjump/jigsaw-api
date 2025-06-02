package cmd

import (
	"github.com/feilongjump/jigsaw-api/config"
	"github.com/feilongjump/jigsaw-api/plugins/validator"
	"github.com/feilongjump/jigsaw-api/routes"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动拼图 API 服务",
	Long:  "启动拼图 API 服务",
	Run:   runServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func runServer(cmd *cobra.Command, args []string) {
	appConfig := config.GetAppConfig()

	r := routes.SetupRouter()

	// 注册自定义验证器
	validator.RegisterValidator()

	r.SetTrustedProxies([]string{appConfig.Host})
	r.Run(":" + appConfig.Port)
}
