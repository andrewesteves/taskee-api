package services

import (
	"strconv"

	"github.com/andrewesteves/taskee-api/entities"
	"github.com/andrewesteves/taskee-api/validations"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// TaskService type
type TaskService struct {
	DB *gorm.DB
}

// Store new resource
func (t TaskService) Store(ctx *fiber.Ctx) {
	var err error
	task := new(entities.Task)

	if err = ctx.BodyParser(task); err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"message": "Whoops! We could not process your request",
		})
		return
	}

	if len(validations.TaskStore(*task)) > 0 {
		ctx.Status(422).JSON(validations.TaskStore(*task))
		return
	}

	t.DB.Save(&task)
	ctx.JSON(fiber.Map{
		"task":    task,
		"message": "Congratulations you have created a new task",
	})
}

// Update a resource
func (t TaskService) Update(ctx *fiber.Ctx) {
	var err error
	task := new(entities.Task)

	if err = ctx.BodyParser(task); err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"message": "Whoops! We could not process your request",
		})
		return
	}

	if len(validations.TaskStore(*task)) > 0 {
		ctx.Status(422).JSON(validations.TaskStore(*task))
		return
	}

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"message": "Whoops! We could not process your request",
		})
		return
	}

	var taskDB entities.Task
	t.DB.First(&taskDB, id)
	t.DB.Model(&taskDB).Updates(map[string]interface{}{
		"description":  task.Description,
		"completed_at": task.CompletedAt,
	})

	ctx.JSON(fiber.Map{
		"message": "You have successfully updated the task",
	})
}

// Destroy a resource
func (t TaskService) Destroy(ctx *fiber.Ctx) {
	var err error

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"message": "Whoops! We could not process your request",
		})
		return
	}

	var task entities.Task
	task.ID = uint(id)
	t.DB.Unscoped().Delete(&task)

	ctx.JSON(fiber.Map{
		"message": "You have successfully deleted the task",
	})
}
