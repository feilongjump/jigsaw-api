package db

import (
	"fmt"
	"net/url"
	"time"

	"github.com/feilongjump/jigsaw-api/infrastructure/config"
	"github.com/feilongjump/jigsaw-api/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var globalDB *gorm.DB

func Init() {
	databaseConfig := config.Global.Database

	// 配置 GORM 日志级别
	var gormLogLevel gormlogger.LogLevel
	if config.IsProd() {
		gormLogLevel = gormlogger.Warn
	} else {
		gormLogLevel = gormlogger.Info
	}

	timeZone := config.Global.App.TimeZone
	// 对时区值进行 URL 编码
	timeZone = url.QueryEscape(timeZone)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Database,
		databaseConfig.Charset,
		timeZone,
	)
	mysqlConfig := mysql.Config{
		DSN: dsn,
		// 禁止启用datetime精度
		DisableDatetimePrecision: true,
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		// 禁止启用外键约束功能
		DisableForeignKeyConstraintWhenMigrating: true,
		// 配置 GORM 日志级别
		Logger: gormlogger.Default.LogMode(gormLogLevel),
	})
	if err != nil {
		logger.Fatal("数据库连接失败: " + err.Error())
	}
	globalDB = db

	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("获取数据库连接失败: " + err.Error())
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sqlDB.SetMaxIdleConns(100)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(25)
	// SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(5) * time.Minute)

	logger.Info("数据库连接成功")
}

// Get 获取数据库连接
func Get() *gorm.DB {
	return globalDB
}
