package model

import "time"

type User struct {
	Id_user    string     `db:"id_user" form:"id_user" json:"id_user" valid:"-" `
	First_name string     `db:"first_name" form:"first_name" json:"first_name" valid:"-" `
	Last_name  string     `db:"last_name" form:"last_name" json:"last_name" valid:"-" `
	Username   string     `db:"username" form:"username" json:"username" valid:"-" `
	Email      string     `db:"email" form:"email" json:"email" valid:"-" `
	Pass       string     `db:"pass" form:"pass" json:"pass" valid:"stringlength(8|20)~create password from 8-20 character"`
	Phone      string     `db:"phone" form:"phone" json:"phone" valid:"-" `
	Role       int        `db:"role" form:"role" json:"role" valid:"-" `
	Created_at *time.Time `db:"created_at" json:"created_at" valid:"-" `
	Updated_at *time.Time `db:"updated_at" json:"updated_at" valid:"-" `
}
