package migrations

import (
	"fmt"
	"log"

	"github.com/DragostinH/CallCentreOrderManagementTool/database"
	"github.com/DragostinH/CallCentreOrderManagementTool/models"
)

func Migrate() {
	fmt.Println("Starting migration...")

	err := database.DB.AutoMigrate(
		&models.Address{},
		&models.Category{},
		&models.Customer{},
		&models.Order{},
		&models.OrderItem{},
		&models.PriceInfo{},
		&models.Product{},
	)

	if err != nil {
		log.Fatal("Issue with migrations...", err)
	}

	fmt.Println("Completed migrations...")
}

func ClearTableData() {
	fmt.Println("Clearing tables...")

	database.DB.Exec("DELETE FROM order_items")
	database.DB.Exec("DELETE FROM orders")
	database.DB.Exec("DELETE FROM product_categories")
	database.DB.Exec("DELETE FROM products")
	database.DB.Exec("DELETE FROM customers")
	database.DB.Exec("DELETE FROM categories")

	fmt.Println("Cleared tables...")
}
