package main

import (
	"log"

	"github.com/DragostinH/CallCentreOrderManagementTool/database"
	"github.com/DragostinH/CallCentreOrderManagementTool/models"
	"github.com/DragostinH/CallCentreOrderManagementTool/seeders"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load config
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using defaults")
	}

	// 2. Connect
	database.Connect()

	// 3. Migrate (Ensures tables exist before seeding)
	log.Println("Running migrations...")
	database.DB.AutoMigrate(
		&models.Category{},
		&models.Product{},
		&models.Customer{},
		&models.Order{},
		&models.OrderItem{},
	)

	// 4. Run the seeders
	log.Println("Starting seeding process...")
	seeders.SeedAll()
}
