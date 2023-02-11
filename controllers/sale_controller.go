package controllers

import (
	"github.com/aldoshb/fiber-gorm/config/database"
	"github.com/aldoshb/fiber-gorm/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllSale(c *fiber.Ctx) error {

	var sales []models.Sale

	database.DB.Debug().Preload("Customer").Preload("SaleItems").Find(&sales)

	return c.JSON(fiber.Map{
		"sales": sales,
	})
}
func GetSaleById(c *fiber.Ctx) error {

	var sales []models.Sale
	id := c.Params("id")
	database.DB.Debug().Preload("Customer").Preload("SaleItems").Where("id = ?", id).First(&sales)

	var totalPrice int
	row := database.DB.Debug().Table("sale_items").Where("sale_id = ?", id).Select("sum(price)").Row()
	row.Scan(&totalPrice)
	return c.JSON(fiber.Map{
		"sales":      sales,
		"totalPrice": totalPrice,
	})

}

func CreateSale(c *fiber.Ctx) error {
	sale := new(models.Sale)

	// parse body request to object struct
	if err := c.BodyParser(sale); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	if sale.SaleCode == "" {
		return c.Status(404).JSON(fiber.Map{
			"error": "Sale Code is required",
		})
	}
	if sale.CustomerID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "Customer ID is required",
		})
	}
	database.DB.Create(&sale)

	return c.JSON(fiber.Map{
		"message": "Create data sale successfully",
		"sale":    sale,
	})
}
