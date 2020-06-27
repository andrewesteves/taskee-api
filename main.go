package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.JSON(map[string]string{
			"message": "Welcome to Taskee API",
		})
	})

	app.Listen(3010)
}
