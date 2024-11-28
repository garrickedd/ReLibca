package model

import "github.com/garickedd/ReLibca/src/domain/model"

type BookCategoryRelation struct {
	model.BaseModel

	BookId     int `gorm:"not null"` // FK -> Book
	CategoryId int `gorm:"not null"` // FK -> BookCategory

	// Relation
	Book     Book         `gorm:"foreignKey:BookId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Category BookCategory `gorm:"foreignKey:CategoryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}
