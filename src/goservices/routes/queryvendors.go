package routes

import (
	"goservices/controllers"

	"github.com/gofiber/fiber/v2"
)

func MapQueryVendorRoutes(app *fiber.App) {
	app.Get("/vendor/:id", controllers.GetVendor)
	app.Get("/vendors", controllers.GetVendors)
}
