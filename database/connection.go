package database

import (
	"os"
	"practicalblog/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db := os.Getenv("DB")

	connection, err := gorm.Open(mysql.Open(db), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.Article{})
	connection.AutoMigrate(&models.Comment{})

}
