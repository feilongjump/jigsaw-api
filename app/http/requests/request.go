package requests

import (
	"github.com/feilongjump/jigsaw-api/app/http/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type getErrMsgFunc func(string) string

// ValidateJSON 验证器
func ValidateJSON(ctx *gin.Context, obj any, getErrMessage getErrMsgFunc) bool {
	if err := ctx.ShouldBindJSON(obj); err != nil {

		// 使用反射获取结构体类型
		objT := reflect.TypeOf(obj)
		// 检查是否为指针类型
		if objT.Kind() == reflect.Pointer {
			// 获取指针指向的实际类型
			objT = objT.Elem()
		}

		// 我更希望能有多个错误，而不是一个错误
		// 但作者回复了，好像不会去进行支持
		// https://github.com/go-playground/validator/issues/639
		errors := make(map[string]string)

		for _, v := range err.(validator.ValidationErrors) {
			fieldName := v.Field()
			// 获取错误信息
			msg := getErrMessage(fieldName + "." + v.Tag())
			if msg == "" {
				msg = v.Error()
			}

			// 获取字段名称
			if field, ok := objT.FieldByName(fieldName); ok {
				fieldName = field.Tag.Get("json")
			}

			errors[fieldName] = msg
		}

		responses.ValidatorError(ctx, errors)

		return false
	}

	return true
}
