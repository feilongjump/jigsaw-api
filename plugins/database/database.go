package database

import (
	"github.com/feilongjump/jigsaw-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func Initialize() {
	db := Connect()

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// 设置空闲连接池中的最大连接数。
	sqlDB.SetMaxIdleConns(10)
	// 设置数据库的最大打开连接数。
	sqlDB.SetMaxOpenConns(100)
	// 设置可以重复使用连接的最长时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func Connect() *gorm.DB {
	var databaseConfig = config.GetDatabaseConfig()

	var err error

	if databaseConfig.Driver == "mysql" {
		DB, err = gorm.Open(mysql.Open(config.GetMySQLDSN()), &gorm.Config{
			// 禁止启用外键约束功能
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	}

	if err != nil {
		panic(err)
	}

	return DB
}
