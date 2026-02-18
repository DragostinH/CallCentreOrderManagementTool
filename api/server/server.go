package server

import (
	"log"
	"net/http"

	"github.com/DragostinH/CallCentreOrderManagementTool/database"
	"github.com/DragostinH/CallCentreOrderManagementTool/handlers/customer"
	"github.com/DragostinH/CallCentreOrderManagementTool/handlers/order"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func ServerStart() {
	database.Connect()
	corsSet := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})

	routes := chi.NewRouter()
	handler := corsSet.Handler(routes)

	// customer routes
	routes.Get("/customers", customer.GetCustomers)
	routes.Get("/customer/{customer_number}", customer.GetCustomer)
	routes.Get("/customers/search", customer.SearchCustomers)

	// order routes
	routes.Get("/orders", order.GetOrders)
	routes.Get("/order/{order_number}", order.GetOrder)

	log.Println("Server starting on port 3000...")
	http.ListenAndServe(":3000", handler)
}
