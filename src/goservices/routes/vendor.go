package routes

import (
	"goservices/controllers"

	"github.com/gofiber/fiber/v2"
)

func MapVendorRoutes(app *fiber.App) {
	app.Post("/vendor", controllers.GetVendor)
	app.Put("/vendor/:id", controllers.GetVendor)
	app.Delete("/vendor/:id", controllers.GetVendor)
}
