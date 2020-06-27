package routes

import (
	"github.com/andrewesteves/taskee-api/services"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// Users routes
func Users(app *fiber.App, db *gorm.DB) {
	service := &services.User{
		DB: db,
	}
	users := app.Group("/users")
	users.Post("/register", service.Register)
}
