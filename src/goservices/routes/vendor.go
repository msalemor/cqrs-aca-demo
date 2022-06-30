package routes

import (
	"goservices/controllers"

	"github.com/gofiber/fiber/v2"
)

func MapVendorRoutes(app *fiber.App) {
	app.Post("/vendor", controllers.CreateVendor)
	app.Put("/vendor/:id", controllers.PutVendor)
	app.Delete("/vendor/:id", controllers.DeleteVendor)
}
