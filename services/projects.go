package services

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// ProjectService type
type ProjectService struct {
	DB *gorm.DB
}

// Index new users
func (p ProjectService) Index(ctx *fiber.Ctx) {
	ctx.JSON(fiber.Map{
		"message": "Projects resources",
	})
}
