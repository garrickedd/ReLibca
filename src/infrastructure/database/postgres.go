package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Postgres() (*sqlx.DB, error) {
	err := godotenv.Load("/mnt/d/Engineer/Server-adv/Project/ReLibca/.env")
	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbName, port)

	// fmt.Println(config)
	// // return sqlx.Open("postgres", config)
	return sqlx.Connect("postgres", config)

	// // Kết nối PostgreSQL
	// db, err := sqlx.Connect("postgres", config)
	// if err != nil {
	// 	return nil, err
	// }

	// // Kiểm tra kết nối
	// err = db.Ping()
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println("Kết nối PostgreSQL thành công!")
	// return db, nil
}
