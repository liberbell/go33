package scraper

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err := sql.Open("mysql", "user01@tcp(127.0.0.1:3306)/msk")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
		CREATE TABLE IF NOT EXISTS blogs (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		date VARCHAR(255) NOT NULL,
		author VARCHAR(255)
		);
	`
}
