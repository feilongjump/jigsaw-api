package validator

import (
	"reflect"

	"github.com/feilongjump/jigsaw-api/domain/entity"
	"github.com/go-playground/validator/v10"
)

// registerWalletTypeValidation 注册钱包类型验证
func registerWalletTypeValidation(v *validator.Validate) {
	v.RegisterValidation("wallet_type", func(fl validator.FieldLevel) bool {
		field := fl.Field()
		var value uint64
		switch field.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value = field.Uint()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value = uint64(field.Int())
		default:
			return false
		}
		return value == uint64(entity.UserWalletTypeCash) ||
			value == uint64(entity.UserWalletTypeBankCard) ||
			value == uint64(entity.UserWalletTypeWeChat) ||
			value == uint64(entity.UserWalletTypeAlipay) ||
			value == uint64(entity.UserWalletTypeCreditCard) ||
			value == uint64(entity.UserWalletTypeStoredValue) ||
			value == uint64(entity.UserWalletTypeInvestment) ||
			value == uint64(entity.UserWalletTypeMargin)
	})
}
