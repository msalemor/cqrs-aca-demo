package main

import (
	"goservices/configs"
	"goservices/routes"
	"log"
)

func main() {
	config := configs.New()
	app := config.FiberApp

	routes.MapVendorRoutes(app)
	log.Fatal(app.Listen(config.SPort))
}
