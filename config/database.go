package config

import (
	"fmt"
	"github.com/feilongjump/jigsaw-api/plugins/viper_lib"
)

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Charset  string
}

func GetDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Driver:   viper_lib.GetConfig().GetString("database.driver"),
		Host:     viper_lib.GetConfig().GetString("database.host"),
		Port:     viper_lib.GetConfig().GetString("database.port"),
		Database: viper_lib.GetConfig().GetString("database.database"),
		Username: viper_lib.GetConfig().GetString("database.username"),
		Password: viper_lib.GetConfig().GetString("database.password"),
		Charset:  viper_lib.GetConfig().GetString("database.charset"),
	}
}

func GetMySQLDSN() string {
	mysqlC := GetDatabaseConfig()

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		mysqlC.Username,
		mysqlC.Password,
		mysqlC.Host,
		mysqlC.Port,
		mysqlC.Database,
		mysqlC.Charset,
	)
}
