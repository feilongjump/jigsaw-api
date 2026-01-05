package main

import (
	"github.com/feilongjump/jigsaw-api/api/router"
	"github.com/feilongjump/jigsaw-api/infrastructure/config"
	"github.com/feilongjump/jigsaw-api/infrastructure/db"
	"github.com/feilongjump/jigsaw-api/pkg/logger"
	"github.com/feilongjump/jigsaw-api/pkg/validator"

	carbonPkg "github.com/feilongjump/jigsaw-api/pkg/carbon"
)

func main() {

	config.Init()

	logger.Init()
	defer logger.Sync()

	carbonPkg.Init()

	validator.Init()

	db.Init()

	// 暂时先在所有环境下自动执行 AutoMigrate
	// 后续区分生产环境以及开发环境的执行方式
	db.AutoMigrate()
	// if !config.IsProd() {
	// 非生产环境下，自动执行 AutoMigrate
	// }

	r := router.Init()

	if err := r.Run(":" + config.Global.App.Port); err != nil {
		logger.Fatal("服务启动失败：" + err.Error())
	}
}
