package order

import (
	"encoding/json"
	"net/http"

	"github.com/DragostinH/CallCentreOrderManagementTool/database"
	"github.com/DragostinH/CallCentreOrderManagementTool/models"
	"github.com/go-chi/chi/v5"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order
	result := database.DB.Preload("Items").Find(&orders)

	if len(orders) == 0 {
		http.Error(w, "No orders found.", http.StatusBadRequest)
		return
	}

	if result.Error != nil {
		http.Error(w, "Couldn't find orders. Are they seeded?", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	orderNumber := chi.URLParam(r, "order_number")

	result := database.DB.
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Categories").
		Where("order_number = ?", orderNumber).
		First(&order)

	if result.Error != nil {
		http.Error(w, "Couldn't find anything with this order number when executing the query", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {

}
