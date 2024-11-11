package entity

import _ "gorm.io/gorm"

type Category struct {
	CategoryId string `json:"category_id" gorm:"primaryKey;not null;type:varchar(20)"`
	Name       string `json:"name" gorm:"type:text;not null"`
}
