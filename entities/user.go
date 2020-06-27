package entities

import "github.com/jinzhu/gorm"

// User entity
type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Token    string    `json:"token"`
	Projects []Project `json:"projects"`
}
