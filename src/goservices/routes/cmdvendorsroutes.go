package routes

import (
	"goservices/controllers"

	"github.com/gofiber/fiber/v2"
)

func MapCommandVendorsRoutes(app *fiber.App) {
	app.Post("/vendor", controllers.CreateVendor)
	app.Post("/vendors", controllers.CreateVendors)
	app.Put("/vendor/:id", controllers.EditVendor)
	app.Delete("/vendor/:id", controllers.DeleteVendor)
}
