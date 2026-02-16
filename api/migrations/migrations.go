package migrations

import (
	"fmt"
	"log"

	"github.com/DragostinH/CallCentreOrderManagementTool/database"
	"github.com/DragostinH/CallCentreOrderManagementTool/models"
)

func CreateDatabase() {
	database.DB.Exec("CREATE DATABSE IF NOT EXISTS call_centre_db")

	database.DB.Exec("USE call_centre_db")
}

func ClearTable() {
	fmt.Println("Truncating data")

	database.DB.Exec("SET FOREIGN_KEY_CHECKS = 0")

	tables := []string{
		"order_items",
		"orders",
		"product_categories",
		"products",
		"customers",
		"categories",
	}

	for _, table := range tables {
		if err := database.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table)).Error; err != nil {
			log.Printf("Warning: Could not truncate %s: %v", table, err)
		}
	}

	database.DB.Exec("SET FOREIGN_KEY_CHECKS = 1")

	fmt.Println("Tables truncated!")
}

func Migrate() {
	fmt.Println("Starting migration...")

	err := database.DB.AutoMigrate(
		&models.Category{},
		&models.Customer{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	fmt.Println("Migration completed successfully.")
}
