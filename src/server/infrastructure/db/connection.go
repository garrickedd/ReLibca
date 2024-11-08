package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection() *gorm.DB {
	// err := godotenv.Load("../../.env")

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	Host 		:= os.Getenv("POSTGRES_HOST")
	Port 		:= os.Getenv("POSTGRES_PORT")
	User 		:= os.Getenv("POSTGRES_USER")
	Password 	:= os.Getenv("POSTGRES_PASSWORD")
	Dbname 		:= os.Getenv("POSTGRES_DB")

	dsn := "host=" + Host + "user=" + User + "password=" + Password + "dbname=" + Dbname + "port=" + Port + "sslmode=disable"

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

	)
}