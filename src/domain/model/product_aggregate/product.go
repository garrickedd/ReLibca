package model

import "github.com/garickedd/ReLibca/src/domain/model"

type Product struct {
	model.BaseModel
	Name        string `gorm:"type:text;not null"`
	Description string `gorm:"type:text;null"`
	Status      bool   `gorm:"type:text;null"`
	TypeId      int    `gorm:"not null"` // FK -> ProductTypes
	PropertyId  int    `gorm:"not null"` // FK -> ProductProperties

	// Relation
	Type     ProductType     `gorm:"foreignKey:TypeId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Property ProductProperty `gorm:"foreignKey:PropertyId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}
