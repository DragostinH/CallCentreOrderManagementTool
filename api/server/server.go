package server

import (
	"log"
	"net/http"

	"github.com/DragostinH/CallCentreOrderManagementTool/database"
	"github.com/DragostinH/CallCentreOrderManagementTool/handlers/customer"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func ServerStart() {
	corsSet := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:5173", "http://localhost:3000"}, // Your Vue URLs
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	database.Connect()

	routes := chi.NewRouter()
	handler := corsSet.Handler(routes)

	routes.Get("/customers", customer.GetCustomers)
	log.Println("Server starting on port 3000...")
	http.ListenAndServe(":3000", handler)
}
