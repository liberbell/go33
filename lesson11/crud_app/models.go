package crud_app

import (
	"gorm.io/gorm"
)

type User struct {
	ID    int    `gorm:"primarykey;autoIncrement"`
	Name  string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:varchar(100);unique;not null"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
