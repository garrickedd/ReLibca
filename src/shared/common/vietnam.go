package common

import (
	"log"
	"regexp"
)

const vietnameseMobileNumberPattern string = `^(03|05|07|08|09|01[2689])[0-9]{8}$`

func VietnameseMobileNumberValidate(mobileNumber string) bool {
	res, err := regexp.MatchString(vietnameseMobileNumberPattern, mobileNumber)
	if err != nil {
		log.Print(err.Error())
	}

	return res
}
