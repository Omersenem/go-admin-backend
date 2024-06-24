package database

import (
	"fmt"
	"github.com/your/repo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:rootroot@123@/go_admin"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to database")

	DB = database

	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
}
