package main

import (
	"goservices/configs"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	configs.Init()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(configs.Config.App.Listen(":3000"))
}
