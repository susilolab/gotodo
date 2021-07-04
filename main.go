package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/susilolab/gotodo/services"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})
	services.TodoServices(app)

	app.Listen(":3000")
}
