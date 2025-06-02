package config

import "github.com/feilongjump/jigsaw-api/plugins/viper_lib"

type JWT struct {
	Secret  string
	Expires int64
}

func GetJWTConfig() *JWT {
	return &JWT{
		Secret:  viper_lib.GetConfig().GetString("jwt.secret"),
		Expires: viper_lib.GetConfig().GetInt64("jwt.expires"),
	}
}
