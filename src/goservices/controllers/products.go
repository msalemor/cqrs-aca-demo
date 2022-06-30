package controllers

import (
	"context"
	"goservices/configs"
	"goservices/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	var product models.Product
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	validate := configs.App.Validate
	if validationErr := validate.Struct(&product); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newProduct := models.Product{
		ID:          primitive.NewObjectID(),
		Name:        product.Name,
		Price:       product.Price,
		Weight:      product.Weight,
		Size:        product.Size,
		CreatedDate: time.Now().UTC(),
	}

	result, err := configs.App.Collection.InsertOne(ctx, newProduct)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}

	return c.Status(http.StatusCreated).JSON(models.Response{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func PutProduct(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
func DeleteProduct(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
