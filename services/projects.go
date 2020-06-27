package services

import (
	"github.com/andrewesteves/taskee-api/entities"
	"github.com/andrewesteves/taskee-api/validations"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// ProjectService type
type ProjectService struct {
	DB *gorm.DB
}

// Index list resources
func (p ProjectService) Index(ctx *fiber.Ctx) {
	var projects []entities.Project
	p.DB.Preload("Tasks").Where("user_id = ?", ctx.Locals("user").(entities.User).ID).Find(&projects)
	ctx.JSON(projects)
}

// Store new resource
func (p ProjectService) Store(ctx *fiber.Ctx) {
	var err error
	project := new(entities.Project)

	if err = ctx.BodyParser(project); err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"message": "Whoops! We could not process your request",
		})
		return
	}

	if len(validations.ProjectStore(*project)) > 0 {
		ctx.Status(422).JSON(validations.ProjectStore(*project))
		return
	}

	project.UserID = ctx.Locals("user").(entities.User).ID
	p.DB.Save(&project)
	ctx.JSON(fiber.Map{
		"message": "Congratulations you have created a new project",
	})
}
