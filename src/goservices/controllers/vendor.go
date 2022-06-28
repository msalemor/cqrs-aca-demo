package controllers

import "github.com/gofiber/fiber/v2"

func GetVendor(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func PutVendor(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
func DeleteVendor(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
