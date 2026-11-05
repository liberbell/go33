package todolist

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (database *Database) InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "user01:password@tcp(localhost:3306)/mak?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect database")
	}

	db.AutoMigrate(&Task{})
	database.DB = db
	return db
}

func (database *Database) CloseDB() {
	database.DB.Close()
}

func (database *Database) CreateTask(title string, description string, completed bool) Task {
	task := Task{Title: title, Description: description, Completed: completed}
	database.DB.Create(&task)
	return task
}

func (database *Database) UpdateTask(task Task) Task {
	database.DB.Save(&task)
	return task
}

func (database *Database) DeleteTask(id uint) bool {
	if err := database.DB.Delete(&Task{}, id).Error; err != nil {
		return false
	}
	return true
}

func (database *Database) GetTasks() []Task {
	var tasks []Task
	database.DB.Find(&tasks)
	return tasks
}
