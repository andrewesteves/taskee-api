package main

import (
	"github.com/andrewesteves/taskee-api/routes"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	db := Connect()
	defer db.Close()

	app.Get("/", func(c *fiber.Ctx) {
		c.JSON(map[string]string{
			"message": "Welcome to Taskee API",
		})
	})

	routes.Users(app, db)

	app.Listen(3010)
}
