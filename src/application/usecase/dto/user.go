package dto

import modeluser "github.com/garickedd/ReLibca/src/domain/model/user_aggregate"

type TokenDetail struct {
	AccessToken            string
	RefreshToken           string
	AccessTokenExpireTime  int64
	RefreshTokenExpireTime int64
}

type RegisterUserByUsername struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

func ToUserModel(from RegisterUserByUsername) modeluser.User {
	return modeluser.User{Username: from.Username,
		FirstName: from.FirstName,
		LastName:  from.LastName,
		Email:     from.Email,
	}
}

type RegisterLoginByMobile struct {
	MobileNumber string
	Otp          string
}

type LoginByUsername struct {
	Username string
	Password string
}
