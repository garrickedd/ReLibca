package model

import "github.com/garickedd/ReLibca/src/domain/model"

type UserRole struct {
	model.BaseModel
	UserId int `gorm:"not null"` // FK -> User
	RoleId int `gorm:"not null"` // FK -> Role

	// Relation
	User User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Role Role `gorm:"foreignKey:RoleId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}
