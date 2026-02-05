package main

import (
	"log"

	"github.com/DragostinH/CallCentreOrderManagementTool/database"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database.Connect()
}
