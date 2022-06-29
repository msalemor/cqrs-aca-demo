package main

import (
	"goservices/configs"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	configs.Init()

	app := configs.Config.App

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(configs.Config.SPort))
}
