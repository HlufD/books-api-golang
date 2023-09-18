package config

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var ConnectToDb = func() {
	dsn := "root:root@1318@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Faild to connect to Database", err)
	}
	db = conn
	fmt.Println("Database connected")
}

var GetDb = func() *gorm.DB {
	return db
}
