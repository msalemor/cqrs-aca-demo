package routes

import (
	"goservices/controllers"

	"github.com/gofiber/fiber/v2"
)

func MapProductsRoutes(app *fiber.App) {
	app.Post("/product", controllers.CreateProduct)
	app.Put("/product/:id", controllers.PutProduct)
	app.Delete("/product/:id", controllers.DeleteProduct)
}
