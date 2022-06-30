package controllers

import (
	"context"
	"goservices/models"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func CreateVendor(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	var vendor models.Vendor
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&vendor); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&vendor); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newVendor := models.Vendor{
		Id:       primitive.NewObjectID(),
		Name:     vendor.Name,
		Location: vendor.Location,
		Title:    vendor.Title,
	}

	result, err := userCollection.InsertOne(ctx, newVendor)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}

	return c.Status(http.StatusCreated).JSON(models.Response{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func PutVendor(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
func DeleteVendor(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
