package db

import (
	"github.com/garrickedd/ReLibca/src/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:admin@localhost:5432/relibcadb"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	DB = db
}