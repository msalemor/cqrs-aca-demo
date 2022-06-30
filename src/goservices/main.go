package main

import (
	"goservices/configs"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config := configs.New()
	app := config.FiberApp

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(config.SPort))
}
