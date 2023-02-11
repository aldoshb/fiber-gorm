package controllers

import (
	"github.com/aldoshb/fiber-gorm/config/database"
	"github.com/aldoshb/fiber-gorm/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllCustomer(c *fiber.Ctx) error {

	var customers []models.Customer

	database.DB.Preload("Sale").Find(&customers)

	return c.JSON(fiber.Map{
		"customers": customers,
	})
}

func CreateCustomer(c *fiber.Ctx) error {
	customer := new(models.Customer)

	// parse body request to object struct
	if err := c.BodyParser(customer); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	if customer.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"err": "name is required",
		})
	}

	database.DB.Create(&customer)

	return c.Status(200).JSON(fiber.Map{
		"message":  "Create data successfully",
		"customer": customer,
	})
}
