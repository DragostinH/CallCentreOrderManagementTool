package customer

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DragostinH/CallCentreOrderManagementTool/database"
	"github.com/DragostinH/CallCentreOrderManagementTool/models"
	"github.com/go-chi/chi/v5"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	var customers []models.Customer

	database.DB.Preload("Orders.Items").Find(&customers)

	if len(customers) == 0 {
		http.Error(w, "No customers found... Are  they seeded?", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	customerNumber := chi.URLParam(r, "customer_number")

	if len(customerNumber) == 0 {
		http.Error(w, "No customer number provided", http.StatusBadRequest)
		return
	}

	result := database.DB.Preload("Orders").Where("customer_number = ?", customerNumber).First(&customer)

	if result.Error != nil {
		http.Error(w, "Something happened with the request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Printf("result %v", customer)

	json.NewEncoder(w).Encode(customer)
}

func SearchCustomers(w http.ResponseWriter, r *http.Request) {
	var customers []models.Customer
	queryParams := r.URL.Query()

	// form search params
	firstName := queryParams.Get("first_name")
	lastName := queryParams.Get("last_name")
	customerNumber := queryParams.Get("customer_number")
	email := queryParams.Get("email")
	phone := queryParams.Get("phone")
	orderNumber := queryParams.Get("order_number")
	postCode := queryParams.Get("post_code")

	query := database.DB.Preload("Orders").Model(&models.Customer{})

	if firstName != "" {
		query = query.Where("first_name LIKE ?", "%"+firstName+"%")
	}

	if lastName != "" {
		query = query.Where("last_name LIKE ?", "%"+lastName+"%")
	}

	if customerNumber != "" {
		query = query.Where("customer_number LIKE ?", "%"+customerNumber+"%")
	}

	if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}

	if phone != "" {
		query = query.Where("phone LIKE ?", "%"+phone+"%")
	}

	if orderNumber != "" {
		query = query.Joins("JOIN orders ON orders.customer_id = customers.id").Where("orders.id = ?", orderNumber)
	}

	if postCode != "" {
		query = query.Where("address_post_code LIKE ?", postCode)
	}

	result := query.Find(&customers)

	if result.Error != nil {
		http.Error(w, "Coouldn't find anything with given search params or something else happened...", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
