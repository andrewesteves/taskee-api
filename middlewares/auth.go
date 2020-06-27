package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber"
)

// NewAuth middleware
func NewAuth() func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		token := ctx.Fasthttp.Request.Header.Peek("taskee-token")
		fmt.Println(string(token))
		ctx.Next()
	}
}
