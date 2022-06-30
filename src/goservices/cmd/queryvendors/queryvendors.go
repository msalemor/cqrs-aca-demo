package main

import (
	"goservices/configs"
	"goservices/routes"
	"log"
)

func main() {
	env := ".env-query-vendors"
	configs.Init(env)
	config := configs.App
	app := config.FiberApp

	routes.MapQueryVendorRoutes(app)
	log.Printf("Starting server at: 0.0.0.0:%s\n", config.SPort)
	log.Fatal(app.Listen(config.SPort))
}
