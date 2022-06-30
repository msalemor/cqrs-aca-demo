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
	var vendor models.Vendor
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&vendor); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	validate := configs.App.Validate
	if validationErr := validate.Struct(&vendor); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newVendor := models.Vendor{
		Id:       primitive.NewObjectID(),
		Name:     vendor.Name,
		Location: vendor.Location,
		Title:    vendor.Title,
	}

	result, err := configs.App.Collection.InsertOne(ctx, newVendor)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.ErrorResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}

	return c.Status(http.StatusCreated).JSON(models.ErrorResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func PutProduct(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
func PostProduct(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
