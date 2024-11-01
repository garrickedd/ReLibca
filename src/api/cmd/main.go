package cmd

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)
func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	// db, err := sql.Open("postgres", dbURL)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}