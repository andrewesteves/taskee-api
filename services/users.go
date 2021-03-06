package services

import (
	"github.com/andrewesteves/taskee-api/entities"
	"github.com/andrewesteves/taskee-api/utils"
	"github.com/andrewesteves/taskee-api/validations"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// UserService type
type UserService struct {
	DB *gorm.DB
}

// Register new users
func (u UserService) Register(ctx *fiber.Ctx) {
	var err error
	user := new(entities.User)

	if err := ctx.BodyParser(user); err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"message": "Whoops! We could not process your request",
		})
		return
	}

	if len(validations.UserRegister(*user)) > 0 {
		ctx.Status(422).JSON(validations.UserRegister(*user))
		return
	}

	var userDB entities.User
	u.DB.Where("email = ?", user.Email).First(&userDB)

	if userDB.ID != 0 {
		ctx.Status(422).JSON(fiber.Map{
			"message": "We already have a registered user with this e-mail",
		})
		return
	}

	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"message": "Whoops! We could not process your credentials correctly",
		})
		return
	}

	user.Token, err = utils.GenerateToken()
	if err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"message": "Whoops! We could not generate your access key",
		})
		return
	}

	u.DB.Save(&user)
	ctx.JSON(user)
}

// Login users
func (u UserService) Login(ctx *fiber.Ctx) {
	user := new(entities.User)
	var userDB entities.User
	var err error

	if err = ctx.BodyParser(user); err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"message": "Whoops! We could not process your request",
		})
		return
	}

	if len(validations.UserLogin(*user)) > 0 {
		ctx.Status(422).JSON(validations.UserLogin(*user))
		return
	}

	u.DB.Where("email = ?", user.Email).First(&userDB)

	if !utils.CompareHash(userDB.Password, user.Password) {
		ctx.Status(401).JSON(fiber.Map{
			"message": "These credentials do not match with our records",
		})
		return
	}

	userDB.Token, err = utils.GenerateToken()
	if err != nil {
		ctx.Status(503).JSON(fiber.Map{
			"message": "Whoops! We could not generate your access key",
		})
		return
	}
	u.DB.Save(&userDB)

	ctx.JSON(fiber.Map{
		"name":  userDB.Name,
		"email": userDB.Email,
		"token": userDB.Token,
	})
}

// Logout users
func (u UserService) Logout(ctx *fiber.Ctx) {
	user := ctx.Locals("user").(entities.User)

	user.Token = ""
	u.DB.Save(&user)

	ctx.JSON(fiber.Map{
		"message": "You are logged out",
	})
}
