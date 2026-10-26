package scraper

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err := sql.Open("mysql", "user01@tcp(127.0.0.1:3306)")
	if err != nil {
		log.Fatal(err)
	}
}
