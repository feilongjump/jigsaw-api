package validator

import (
	"reflect"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

// Init 初始化验证器和翻译器
func Init() {
	// 注册中文翻译器
	zh := zh.New()
	uni = ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")

	// 获取 gin 的校验器引擎
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		validate = v
		// 注册中文翻译
		zh_translations.RegisterDefaultTranslations(v, trans)

		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			label := fld.Tag.Get("label")
			if label != "" {
				return label
			}
			return fld.Name
		})
	}
}

// Translate 翻译错误信息
func Translate(err error, payload interface{}) map[string][]string {
	result := make(map[string][]string)

	errors, ok := err.(validator.ValidationErrors)
	if !ok {
		return result
	}

	for _, e := range errors {
		fieldName := resolveJSONField(payload, e.StructField())
		result[fieldName] = append(result[fieldName], e.Translate(trans))
	}
	return result
}

func resolveJSONField(payload interface{}, structField string) string {
	if payload == nil {
		return lowerFirst(structField)
	}
	structType := reflect.TypeOf(payload)
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}
	if structType.Kind() != reflect.Struct {
		return lowerFirst(structField)
	}
	field, ok := structType.FieldByName(structField)
	if !ok {
		return lowerFirst(structField)
	}
	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		return lowerFirst(structField)
	}
	name := strings.Split(jsonTag, ",")[0]
	if name == "" || name == "-" {
		return lowerFirst(structField)
	}
	return name
}

func lowerFirst(value string) string {
	if value == "" {
		return value
	}
	runes := []rune(value)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}
