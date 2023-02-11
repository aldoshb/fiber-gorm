package routes

import (
	"github.com/aldoshb/fiber-gorm/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Get("/customers", controllers.GetAllCustomer)
	app.Post("/customers", controllers.CreateCustomer)

	app.Get("/sale", controllers.GetAllSale)
	app.Get("/sale/:id", controllers.GetSaleById)
	app.Post("/sale", controllers.CreateSale)

	app.Get("/saleItems", controllers.GetAllSaleItem)
	app.Post("/saleItems", controllers.CreateSaleItem)
}
