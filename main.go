package main

import (
	"github.com/andrewesteves/taskee-api/middlewares"
	"github.com/andrewesteves/taskee-api/routes"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	db := Connect()
	defer db.Close()

	app.Use(cors.New())
	app.Use(middlewares.NewAuth(middlewares.ConfigAuth{
		DB: db,
	}))

	app.Get("/", func(c *fiber.Ctx) {
		c.JSON(map[string]string{
			"message": "Welcome to Taskee API",
		})
	})

	routes.Users(app, db)
	routes.Projects(app, db)
	routes.Tasks(app, db)

	app.Listen(3010)
}
