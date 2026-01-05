package db

import (
	"github.com/feilongjump/jigsaw-api/domain/entity"
	"github.com/feilongjump/jigsaw-api/pkg/logger"
)

func AutoMigrate() {
	logger.Info("开始执行数据库自动迁移...")

	err := globalDB.AutoMigrate(
		&entity.Note{},
	)
	if err != nil {
		logger.Fatal("数据库自动迁移失败: " + err.Error())
	}

	logger.Info("数据库自动迁移成功...")
}
