package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/susilolab/gotodo/models"
)

func TodoServices(app *fiber.App) {
	todo := app.Group("todo")

	// `/todo`
	todo.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Todo")
	})

	// `/todo/1`
	todo.Get("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Todo" + c.Params("id"))
	})

	// `/todo`
	todo.Post("/", func(c *fiber.Ctx) error {
		t := new(models.Todo)
		if err := c.BodyParser(t); err != nil {
			return err
		}

		return c.JSON(t)
	})
}
