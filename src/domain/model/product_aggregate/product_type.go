package model

import "github.com/garickedd/ReLibca/src/domain/model"

type ProductType struct {
	model.BaseModel
	Type string `gorm:"type:string;size:10;not null"`
}
