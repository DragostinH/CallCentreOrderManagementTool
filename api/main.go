package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/DragostinH/CallCentreOrderManagementTool/database"
	"github.com/DragostinH/CallCentreOrderManagementTool/models"
	"github.com/DragostinH/CallCentreOrderManagementTool/seeders"
	"github.com/DragostinH/CallCentreOrderManagementTool/server"
	"github.com/joho/godotenv"
)

func main() {
	mode := flag.String("mode", "server", "What to run: server or setup?")

	flag.Parse()
	// 1. Load config
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using defaults")
	}

	if *mode == "seed" {
		fmt.Println("Starting seeding process...")
		SeedAndMigrate()

	} else {
		fmt.Println("Starting server...")
		Start()

	}

}

func Start() {

	server.ServerStart()
}

func SeedAndMigrate() {

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
