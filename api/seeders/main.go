package seeders

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/DragostinH/CallCentreOrderManagementTool/database"
	"github.com/DragostinH/CallCentreOrderManagementTool/models"
	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	SeedAll()
}

func SeedAll() {
	SeedCategories()
	SeedProducts()
	SeedCustomersWithOrders()
}

func SeedCategories() {
	fmt.Println("Seeding categories...")
	catNames := []string{"Beverages", "Food", "Household", "Electronics", "Clothing"}
	var categories []models.Category
	for _, name := range catNames {
		cat := models.Category{Name: name}
		database.DB.FirstOrCreate(&cat, models.Category{Name: name})
		categories = append(categories, cat)
	}

	fmt.Printf("Seeded %d categories \n", len(catNames))
}

func SeedProducts() {
	fmt.Println("Seeding Products...")
	measures := []string{"kg", "mg", "l", "each", "pack"}
	var products []models.Product
	var allCategories []models.Category
	database.DB.Find(&allCategories)
	if len(allCategories) == 0 {
		log.Fatal("Found 0 categories. Not seeded?")
	}
	for i := 0; i < 100; i++ {
		unitPrice := gofakeit.Price(1, 100)
		measure := measures[rand.Intn(len(measures))]
		measureAmount := 1
		if measure != "each" {
			measureAmount = gofakeit.Number(1, 10)
		}

		product := models.Product{
			ProductUID:  gofakeit.UUID(),
			Name:        gofakeit.ProductName(),
			Image:       gofakeit.ImageURL(300, 300),
			ProductType: gofakeit.ProductMaterial(),
			Eans:        gofakeit.DigitN(13),
			IsAlcoholic: gofakeit.Bool(),
			UnitPrice: models.PriceInfo{
				Price:         unitPrice,
				Measure:       measure,
				MeasureAmount: measureAmount,
			},
			RetailPrice: models.PriceInfo{
				Price:         unitPrice * 1.2,
				Measure:       measure,
				MeasureAmount: measureAmount,
			},
			Categories: []models.Category{
				allCategories[rand.Intn(len(allCategories))],
			},
		}
		database.DB.Create(&product)
		products = append(products, product)
	}
	fmt.Printf("Seeded %d products", len(products))

}

// func SeedCustomersWithOrders() {
// 	status := []string{"pending", "shipped", "delivered"}
// 	var total float64
// 	var items []models.OrderItem
// 	val, ok := database.DB.Get("products")

// 	if !ok {
// 		log.Fatal("no products")
// 	}

// 	products, ok := val.([]models.Product)
// 	if !ok {
// 		log.Fatal("cant assign products from DB to product variable")
// 	}
// 	fmt.Println("Seeding customers with orders...")

// 	for i := 0; i < 50; i++ {
// 		cstmr := models.Customer{
// 			FirstName: gofakeit.FirstName(),
// 			LastName:  gofakeit.LastName(),
// 			Phone:     gofakeit.Phone(),
// 			Email:     gofakeit.Email(),
// 			Address: models.Address{
// 				PostCode: gofakeit.Zip(),
// 				City:     gofakeit.City(),
// 				Street:   gofakeit.Street(),
// 			},
// 			CustomerNumber: gofakeit.LetterN(8),
// 		}

// 		database.DB.Create(&cstmr)

// 		numItems := gofakeit.Number(1, 5)

// 		for j := 0; j < numItems; j++ {
// 			product := products[rand.Intn(len(products))]
// 			quantity := gofakeit.Number(1, 5)
// 			linePrice := product.RetailPrice.Price * float64(quantity)
// 			total += linePrice

// 			items = append(items, models.OrderItem{
// 				ProductID: product.ID,
// 				Quantity:  quantity,
// 				Price:     linePrice,
// 			})
// 		}
// 		order := models.Order{
// 			CustomerID: cstmr.ID,
// 			OrderDate:  gofakeit.Date(),
// 			Status:     status[rand.Intn(len(status))],
// 			Total:      total,
// 			Items:      items,
// 		}
// 		database.DB.Create(&order)

// 	}
// 	fmt.Println("Created 50 Customers and Orders")
// 	fmt.Println("Seeding complete!")

// }

func SeedCustomersWithOrders() {
	fmt.Println("Seeding customers with orders...")

	// FIX: Correct way to get products from GORM
	var allProducts []models.Product
	database.DB.Find(&allProducts)

	if len(allProducts) == 0 {
		fmt.Println("No products found to attach to orders. Run SeedProducts first.")
		return
	}

	statusOptions := []string{"pending", "shipped", "delivered"}

	for i := 0; i < 50; i++ {
		// 1. Create Customer
		cstmr := models.Customer{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Phone:     gofakeit.Phone(),
			Email:     gofakeit.Email(),
			Address: models.Address{
				PostCode: gofakeit.Zip(),
				City:     gofakeit.City(),
				Street:   gofakeit.Street(),
			},
			CustomerNumber: gofakeit.LetterN(8),
		}
		database.DB.Create(&cstmr)

		// 2. Generate Items for an Order
		numItems := gofakeit.Number(1, 5)
		var currentOrderItems []models.OrderItem
		var orderTotal float64

		for j := 0; j < numItems; j++ {
			product := allProducts[rand.Intn(len(allProducts))]
			quantity := gofakeit.Number(1, 5)
			linePrice := product.RetailPrice.Price * float64(quantity)
			orderTotal += linePrice

			currentOrderItems = append(currentOrderItems, models.OrderItem{
				ProductID: product.ID,
				Quantity:  quantity,
				Price:     linePrice,
			})
		}

		// 3. Create the Order
		order := models.Order{
			CustomerID: cstmr.ID,
			OrderDate:  gofakeit.Date(),
			Status:     statusOptions[rand.Intn(len(statusOptions))],
			Total:      orderTotal,
			Items:      currentOrderItems,
		}
		database.DB.Create(&order)
	}
	fmt.Println("Created 50 Customers and their Orders.")
}
