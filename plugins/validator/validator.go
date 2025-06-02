package validator

import (
	"github.com/feilongjump/jigsaw-api/plugins/database"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"strings"
)

// RegisterValidator 注册自定义验证器
func RegisterValidator() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("c_unique_db", CUniqueDb)
	}
}

// CUniqueDb 验证在数据库中是否唯一
func CUniqueDb(fl validator.FieldLevel) bool {
	if fl.Param() == "" {
		// 参数为空，不验证
		return false
	}

	param := strings.Split(fl.Param(), ":")

	if len(param) != 2 {
		// 参数格式错误，不验证
		return false
	}

	table := param[0]
	field := param[1]

	var count int64
	database.DB.Table(table).
		Where(field, fl.Field().String()).
		Count(&count)

	return count == 0
}
