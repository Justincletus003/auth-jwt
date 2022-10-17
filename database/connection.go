package database

import (
	"github.com/Justincletus003/auth-system/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:Pass#123@tcp(localhost:3306)/auth_db"), &gorm.Config{})
	if err != nil {
		panic("Could not connected to database")
	}
	DB = connection
	connection.AutoMigrate(&models.User{})
}