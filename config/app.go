package config

import "github.com/feilongjump/jigsaw-api/plugins/viper_lib"

type App struct {
	Name    string
	Env     string
	Debug   bool
	GinMode string
	Host    string
	Port    string
}

func GetAppConfig() *App {
	return &App{
		Name:    viper_lib.GetConfig().GetString("app.name"),
		Env:     viper_lib.GetConfig().GetString("app.env"),
		Debug:   viper_lib.GetConfig().GetBool("app.debug"),
		GinMode: viper_lib.GetConfig().GetString("app.gin_mode"),
		Host:    viper_lib.GetConfig().GetString("app.host"),
		Port:    viper_lib.GetConfig().GetString("app.port"),
	}
}

func IsLocal() bool {
	return GetAppConfig().Env == "local"
}

func IsProduction() bool {
	return GetAppConfig().Env == "production"
}
