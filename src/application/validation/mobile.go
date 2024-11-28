package validation

import (
	"github.com/garickedd/ReLibca/src/shared/common"
	"github.com/go-playground/validator/v10"
)

func VietnameseMobileNumberValidator(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	return common.VietnameseMobileNumberValidate(value)
}
