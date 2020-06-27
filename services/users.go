package services

import (
	"github.com/andrewesteves/taskee-api/entities"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// User type
type User struct {
	DB *gorm.DB
}

// Register new users
func (u User) Register(ctx *fiber.Ctx) {
	user := new(entities.User)
	if err := ctx.BodyParser(user); err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"message": "Whoops! We could not process your request",
		})
	}
	u.DB.Save(&user)
	ctx.JSON(user)
}
