package controllers

import (
	"github.com/aldoshb/fiber-gorm/config/database"
	"github.com/aldoshb/fiber-gorm/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllSaleItem(c *fiber.Ctx) error {

	var saleItems []models.SaleItem

	database.DB.Preload("Sale").Find(&saleItems)

	return c.JSON(fiber.Map{
		"Items": saleItems,
	})
}

func CreateSaleItem(c *fiber.Ctx) error {
	saleItems := new(models.SaleItem)

	// parse body request to object struct
	if err := c.BodyParser(saleItems); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	// if sale.SaleCode == "" {
	// 	return c.Status(404).JSON(fiber.Map{
	// 		"error": "Sale Code is required",
	// 	})
	// }
	if saleItems.SaleID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "Customer ID is required",
		})
	}
	database.DB.Create(&saleItems)

	return c.JSON(fiber.Map{
		"message": "Create data sale items successfully",
		"item":    saleItems,
	})
}
