package main

import (
	"log"

	"github.com/braydend/birthday-jamz/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString(handlers.HelloWorldHandler())
    })

	app.Get("/catfact", func(c *fiber.Ctx) error {
		fact, err := handlers.GetCatFact()

		if (err != nil) {
			log.Println("error: " + err.Error())
			return c.SendStatus(500)
		}

		return c.SendString(fact.FactString)
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		return c.SendString(handlers.HelloToHandler(c.Params("name")))
	})

    app.Listen(":3000")
}