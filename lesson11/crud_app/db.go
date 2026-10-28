package crud_app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DB struct {
	*gorm.DB
}

func NewDatabases() (*DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loding .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
}
