package entities

import "github.com/jinzhu/gorm"

// Project entity
type Project struct {
	gorm.Model
	UserID      uint
	User        User   `json:"user"`
	Description string `json:"description"`
}
