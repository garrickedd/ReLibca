package db

import (
	"fmt"
	"log"
	"os"

	"github.com/garrickedd/ReLibca/src/server/domain/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection() *gorm.DB {
	err := godotenv.Load("./.env")
	// err := godotenv.Load("../infrastructure/db/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}

	return db
}

func Migrate() {
	db := DbConnection()
	sqlDB, error := db.DB()

	if error != nil {
		log.Fatal(error)
	}

	defer sqlDB.Close()

	log.Println("Automigration working...")
	db.AutoMigrate(
		&entity.User{},
		&entity.Product{},
		&entity.Promotion{},
		&entity.Book{},
		&entity.Category{},
		&entity.Order{},
		&entity.OrderItem{},
		&entity.Payment{},
		&entity.Property{},
	)
}
