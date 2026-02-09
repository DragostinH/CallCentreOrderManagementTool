package customer

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DragostinH/CallCentreOrderManagementTool/database"
	"github.com/DragostinH/CallCentreOrderManagementTool/models"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	var customers []models.Customer

	database.DB.Find(&customers)

	if len(customers) == 0 {
		log.Fatal("No customers found... Are  they seeded?")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
