package middlewares

import (
	"fmt"

	"github.com/andrewesteves/taskee-api/entities"
	"github.com/andrewesteves/taskee-api/utils"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// ConfigAuth config for Auth
type ConfigAuth struct {
	DB *gorm.DB
}

// NewAuth middleware
func NewAuth(config ConfigAuth) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		if guard(string(ctx.Fasthttp.Request.URI().PathOriginal()), ctx.Method()) {
			token := ctx.Fasthttp.Request.Header.Peek("taskee-token")
			var user entities.User
			config.DB.Where("token = ?", string(token)).First(&user)

			if user.ID > 0 {
				ctx.Locals("user", user)
				ctx.Next()
				return
			}

			ctx.Status(401).JSON(fiber.Map{
				"message": "You are not allowed to access this path",
			})
			return
		}

		ctx.Next()
	}
}

// It will act as a guardian
func guard(path string, method string) bool {
	accessable := []string{
		"/-GET",
		"/users/register-POST",
		"/users/login-POST",
	}
	return !utils.Contains(accessable, fmt.Sprintf("%s-%s", path, method))
}
