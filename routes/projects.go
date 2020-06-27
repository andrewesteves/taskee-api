package routes

import (
	"github.com/andrewesteves/taskee-api/services"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// Projects routes
func Projects(app *fiber.App, db *gorm.DB) {
	service := &services.ProjectService{
		DB: db,
	}
	projects := app.Group("/projects")
	projects.Get("/", service.Index)
	projects.Post("/", service.Store)
	projects.Put("/:id", service.Update)
	projects.Delete("/:id", service.Destroy)
}
