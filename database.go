package main

import (
	"fmt"

	"github.com/andrewesteves/taskee-api/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Connect to the database
func Connect() *gorm.DB {
	env, err := EnvConfig{}.Vars()
	if err != nil {
		panic(err.Error())
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", env.DB.User, env.DB.Pass, env.DB.Host, env.DB.Port, env.DB.Name)
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
