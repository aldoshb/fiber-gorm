package main

import (
	"github.com/aldoshb/fiber-gorm/config/database"
	"github.com/aldoshb/fiber-gorm/config/database/migrations"
	"github.com/aldoshb/fiber-gorm/config/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.InitDB()

	migrations.Migration()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Okay",
			"name":    "Iron Man",
		})
	})

	routes.RouteInit(app)

	app.Listen(":8888")
}
