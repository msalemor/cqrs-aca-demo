package main

import (
	"goservices/configs"
	"goservices/routes"
	"log"
)

func main() {

	configs.Init()
	routes.MapVendorRoutes(configs.Config.App)

	log.Fatal(configs.Config.App.Listen(":3000"))
}
