package routes

import (
	"github.com/andrewesteves/taskee-api/services"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// Tasks routes
func Tasks(app *fiber.App, db *gorm.DB) {
	service := &services.TaskService{
		DB: db,
	}
	tasks := app.Group("/tasks")
	tasks.Get("/", service.Index)
	tasks.Post("/", service.Store)
}
