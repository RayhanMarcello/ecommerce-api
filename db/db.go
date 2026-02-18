package db

import (
	"fmt"
	"golang-emarket/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	dsn := "rayhan:rayhan222@tcp(127.0.0.1:3306)/marketplace?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("cant connect to db")
	}
	db.AutoMigrate(models.Customers{}, models.OrderItems{}, models.Orders{}, models.Payments{}, models.Products{})
	return db
}
