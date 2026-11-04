package todolist

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabse() *Database {
	return &Database{}
}

func (database *Database) InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "user01:password@tcp(localhost:3306)/mak?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&Task{})
	database.DB = db
	return db
}

func (database *Database) CloseDB() {
	database.DB.Close()
}
