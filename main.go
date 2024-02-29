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

	app.Get("/top-songs", func(c *fiber.Ctx) error {
		date := c.Query("date")

		if (date == "") {
			return fiber.NewError(fiber.ErrBadRequest.Code, "Missing date parameter")
		}
		playlist, err := handlers.BuildPlaylistHandler(date);

		if err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
		}

		return c.SendString(playlist)
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		return c.SendString(handlers.HelloToHandler(c.Params("name")))
	})

    app.Listen(":3000")
}