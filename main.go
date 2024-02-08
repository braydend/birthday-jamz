package main

import (
	"github.com/braydend/birthday-jamz/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString(handlers.HelloWorldHandler())
    })

	app.Get("/:name", func(c *fiber.Ctx) error {
		return c.SendString(handlers.HelloToHandler(c.Params("name")))
	})

    app.Listen(":3000")
}