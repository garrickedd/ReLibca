package validations

import (
	"github.com/garrickedd/ReLibca/src/server/common"
	"github.com/go-playground/validator/v10"
)

func VietnameseMobileNumberValidator(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	return common.VietnameseMobileNumberValidate(value)
}
