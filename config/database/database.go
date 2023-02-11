package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// INITDB : initialize Database Connection

var DB *gorm.DB

func InitDB() {
	var err error

	dsn := "root:@tcp(127.0.0.1:3306)/learn-go?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed connect to database")
	}

	fmt.Println("Successfully connected to database")
}
