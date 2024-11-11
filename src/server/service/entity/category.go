package entity

import _ "gorm.io/gorm"

type Category struct {
	CategoryId string `gorm:"primaryKey;not null;type:varchar(20)"`
	Name       string `gorm:"type:text;not null"`
}
