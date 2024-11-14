package validations

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func VietnameseMobileNumberValidator(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	// Updated regular expression for Vietnamese mobile numbers based on the JavaScript pattern
	res, err := regexp.MatchString(`^(03|05|07|08|09|01[2689])[0-9]{8}$`, value)
	if err != nil {
		log.Print(err.Error())
	}
	return res
}
