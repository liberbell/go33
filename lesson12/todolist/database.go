package todolist

import "github.com/jinzhu/gorm"

type Database struct {
	DB *gorm.DB
}

func NewDatabse() *Database {
	return &Database{}
}

func (database *Database) InitDB() *gorm.DB {
	db, err := gorm.Open("mysql")
}
