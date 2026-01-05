package validator

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	trans ut.Translator
)

// Init 初始化校验器
func Init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册中文翻译器
		zhT := zh.New()
		uni := ut.New(zhT, zhT)
		trans, _ = uni.GetTranslator("zh")
		zh_translations.RegisterDefaultTranslations(v, trans)

		// 注册自定义 TagNameFunc，优先使用 label 标签作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			if name != "" {
				return name
			}
			name = fld.Tag.Get("json")
			if name != "" {
				return strings.Split(name, ",")[0]
			}
			name = fld.Tag.Get("form")
			if name != "" {
				return strings.Split(name, ",")[0]
			}
			return fld.Name
		})
	}
}

// Translate 翻译校验错误
func Translate(err error, obj any) map[string][]string {
	result := make(map[string][]string)

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return result
	}

	// 获取 obj 的类型对象
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for _, e := range errs {
		// 获取 StructField 名称
		fieldName := e.StructField()
		field, found := t.FieldByName(fieldName)
		if !found {
			continue
		}

		// 确定返回的 map key (json > form > fieldName)
		key := fieldName
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			key = strings.Split(jsonTag, ",")[0]
		} else {
			formTag := field.Tag.Get("form")
			if formTag != "" && formTag != "-" {
				key = strings.Split(formTag, ",")[0]
			}
		}

		// 确定错误信息
		var msg string

		// 尝试从 validator tag 获取自定义错误信息
		validatorTag := field.Tag.Get("validator")
		if validatorTag != "" {
			// 解析 validator tag: "required=页码不能为空,min=最小的页码为1"
			customMsgs := strings.Split(validatorTag, ",")
			for _, m := range customMsgs {
				kv := strings.SplitN(m, "=", 2)
				if len(kv) == 2 {
					k := strings.TrimSpace(kv[0])
					if k == e.Tag() {
						msg = kv[1]
						break
					}
				}
			}
		}

		// 如果没有自定义消息，使用翻译器
		if msg == "" {
			msg = e.Translate(trans)
		}

		// 添加到结果
		result[key] = append(result[key], msg)
	}

	return result
}
