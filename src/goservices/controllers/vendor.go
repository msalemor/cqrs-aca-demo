package controllers

import (
	"context"
	"goservices/configs"
	"goservices/models"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetVendor(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	entityID := c.Params("id")
	var vendor models.Vendor
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(entityID)

	err := configs.App.Collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&vendor)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(vendor)

}
func GetVendors(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var vendors []models.Vendor
	defer cancel()

	results, err := configs.App.Collection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var vendor models.Vendor
		if err = results.Decode(&vendor); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(models.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		vendors = append(vendors, vendor)
	}

	return c.Status(http.StatusOK).JSON(vendors)
}

func CreateVendor(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	var vendor models.Vendor
	defer cancel()

	//validate the request body
	json := string(c.Request().Body())
	log.Println(json)

	if err := c.BodyParser(&vendor); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	validate := configs.App.Validate
	if validationErr := validate.Struct(&vendor); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newVendor := models.Vendor{
		ID:          primitive.NewObjectID(),
		VendorCode:  vendor.VendorCode,
		Name:        vendor.Name,
		Contact:     vendor.Contact,
		Phone:       vendor.Phone,
		Email:       vendor.Email,
		CreatedDate: time.Now().UTC(),
	}

	result, err := configs.App.Collection.InsertOne(ctx, newVendor)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}

	return c.Status(http.StatusCreated).JSON(models.Response{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func CreateVendors(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	var vendors []models.Vendor
	defer cancel()

	//validate the request body
	json := string(c.Request().Body())
	log.Println(json)

	if err := c.BodyParser(&vendors); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	// validate := configs.App.Validate
	// if validationErr := validate.Struct(&vendors); validationErr != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(models.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	// }

	newVendors := []interface{}{}

	for _, vendor := range vendors {
		v := models.Vendor{
			ID:          primitive.NewObjectID(),
			VendorCode:  vendor.VendorCode,
			Name:        vendor.Name,
			Contact:     vendor.Contact,
			Phone:       vendor.Phone,
			Email:       vendor.Email,
			CreatedDate: time.Now().UTC(),
		}
		newVendors = append(newVendors, v)
	}

	result, err := configs.App.Collection.InsertMany(ctx, newVendors)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}

	return c.Status(http.StatusCreated).JSON(models.Response{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func EditVendor(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ID := c.Params("userId")
	app := configs.App
	var vendor models.Vendor
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(ID)

	//validate the request body
	if err := c.BodyParser(&vendor); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := configs.App.Validate.Struct(&vendor); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"name": vendor.Name} //, "location": vendor.Location, "title": vendor.Title}

	result, err := app.Collection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated user details
	var updatedVendor models.Vendor
	if result.MatchedCount == 1 {
		err := app.Collection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedVendor)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(models.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(models.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedVendor}})
}
func DeleteVendor(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	vendorId := c.Params("id")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(vendorId)

	result, err := configs.App.Collection.DeleteOne(ctx, bson.M{"id": objId})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			models.Response{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "User with specified ID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		models.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "User successfully deleted!"}},
	)
}
