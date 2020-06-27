package main

import (
	"fmt"

	"github.com/andrewesteves/taskee-api/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	host = "172.17.0.2"
	user = "root"
	pass = "4321"
	name = "taskee"
	port = 3306
)

// Connect to the database
func Connect() *gorm.DB {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, name)
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
}
