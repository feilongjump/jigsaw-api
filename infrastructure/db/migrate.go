package db

import (
	"github.com/feilongjump/jigsaw-api/domain/entity"
	"github.com/feilongjump/jigsaw-api/pkg/logger"
)

func AutoMigrate() {
	logger.Info("开始执行数据库自动迁移...")

	err := globalDB.AutoMigrate(
		&entity.User{},
		&entity.Note{},
		&entity.File{},
		&entity.UserWallet{},
		&entity.LedgerCategory{},
		&entity.LedgerRecord{},
	)
	if err != nil {
		logger.Fatal("数据库自动迁移失败: " + err.Error())
	}

	if err := Seed(); err != nil {
		logger.Fatal("系统分类初始化失败: " + err.Error())
	}

	logger.Info("数据库自动迁移成功...")
}
