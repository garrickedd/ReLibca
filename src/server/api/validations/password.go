package validations

import (
	"github.com/garrickedd/ReLibca/src/server/common"
	"github.com/go-playground/validator/v10"
)

func PasswordValidation(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		fld.Param()
		return false
	}

	return common.CheckPassword(value)
}
