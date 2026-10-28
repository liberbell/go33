package crud_app

import (
	"log"

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
}
