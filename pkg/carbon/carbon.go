package carbon

import (
	"github.com/dromara/carbon/v2"
	"github.com/feilongjump/jigsaw-api/infrastructure/config"
)

func Init() {
	appConfig := config.Global.App

	carbon.SetTimezone(appConfig.TimeZone)
	carbon.SetLocale(appConfig.Locale)
}
