package main

import (
	"fmt"
	"os"

	"github.com/andrewesteves/taskee-api/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Connect to the database
func Connect() *gorm.DB {
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open("mysql", conn)

	if err != nil {
		panic(err)
	}

	return db
}

// Migrate tables
func Migrate(db *gorm.DB) {
	db.CreateTable(&entities.User{})
	db.CreateTable(&entities.Project{})
	db.CreateTable(&entities.Task{})
}
