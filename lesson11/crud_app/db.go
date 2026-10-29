package crud_app

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4$parseTime=True&log=local", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the database: %v", err)
	}

	if err := AutoMigrate(db); err != nil {
		return nil, fmt.Errorf("Failed to auto-migrate tables: %v", err)
	}

	return &DB{db}, nil

}
