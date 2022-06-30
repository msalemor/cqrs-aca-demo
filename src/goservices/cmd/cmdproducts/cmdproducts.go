package main

import (
	"goservices/configs"
	"goservices/routes"
	"log"
)

func main() {
	configs.Init(".env-products")
	config := configs.App
	app := config.FiberApp

	routes.MapProductsRoutes(app)
	log.Fatal(app.Listen(config.SPort))
}
