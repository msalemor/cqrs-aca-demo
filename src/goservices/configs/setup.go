package configs

import (
	"github.com/gofiber/fiber/v2"
)

type AppConfig struct {
	App      *fiber.App
	MongoURI string
	SPort    string
	DB       string
}

var Config AppConfig

func Init() {
	app := fiber.New()
	Config = AppConfig{App: app, MongoURI: "", SPort: ":3000", DB: ""}
}
