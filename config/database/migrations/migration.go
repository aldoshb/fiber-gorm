package migrations

import (
	"fmt"

	"github.com/aldoshb/fiber-gorm/config/database"
	"github.com/aldoshb/fiber-gorm/models"
)

func Migration() {
	err := database.DB.AutoMigrate(
		&models.Customer{},
		&models.Sale{},
		&models.SaleItem{},
	)

	if err != nil {
		fmt.Println("can't run migration!")
	}
	fmt.Println("Migration Complete!")
}
