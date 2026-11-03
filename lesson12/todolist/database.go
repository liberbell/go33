package todolist

import "github.com/jinzhu/gorm"

type Database struct {
	DB *gorm.DB
}

func NewDatabse() *Database {
	return &Database{}
}

func (database *Database) InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "user01:password@tcp(localhost:3306)/mak?charset=utf8")
}
