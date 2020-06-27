package services

import (
	"github.com/andrewesteves/taskee-api/entities"
	"github.com/andrewesteves/taskee-api/utils"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// User type
type User struct {
	DB *gorm.DB
}

// Register new users
func (u User) Register(ctx *fiber.Ctx) {
	var err error
	user := new(entities.User)

	if err := ctx.BodyParser(user); err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"message": "Whoops! We could not process your request",
		})
	}

	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"message": "Whoops! We could not process your credentials correctly",
		})
	}

	user.Token, err = utils.GenerateToken()
	if err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"message": "Whoops! We could not generate your access key",
		})
	}

	u.DB.Save(&user)
	ctx.JSON(user)
}
