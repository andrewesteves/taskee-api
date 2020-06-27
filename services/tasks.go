package services

import (
	"github.com/andrewesteves/taskee-api/entities"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// TaskService type
type TaskService struct {
	DB *gorm.DB
}

// Index list resources
func (p TaskService) Index(ctx *fiber.Ctx) {
	ctx.JSON(fiber.Map{
		"message": "Tasks resources",
	})
}

// Store new resource
func (p TaskService) Store(ctx *fiber.Ctx) {
	var err error
	task := new(entities.Task)

	if err = ctx.BodyParser(task); err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"message": "Whoops! We could not process your request",
		})
		return
	}

	p.DB.Save(&task)
	ctx.JSON(fiber.Map{
		"message": "Congratulations you have created a new task",
	})
}
