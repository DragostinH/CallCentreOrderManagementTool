package seeders

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/DragostinH/CallCentreOrderManagementTool/database"
	"github.com/DragostinH/CallCentreOrderManagementTool/models"
	"github.com/brianvoe/gofakeit/v6"
)

func SeedAll() {
	// Seed Categories
	SeedCategories()
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
	var categories []models.Category
	if len(categories) == 0 {
		categories = []models.Category{{Name: "Beverages"}}
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
				categories[rand.Intn(len(categories))],
			},
		}
		database.DB.Create(&product)
		products = append(products, product)
	}
	fmt.Printf("Seeded %d products", len(products))

}

func SeedCustomersWithOrders() {
	status := []string{"pending", "shipped", "delivered"}
	var total float64
	var items []models.OrderItem
	val, ok := database.DB.Get("products")

	if !ok {
		log.Fatal("no products")
	}

	products, ok := val.([]models.Product)
	if !ok {
		log.Fatal("cant assign products from DB to product variable")
	}
	fmt.Println("Seeding customers with orders...")
	// what do you need first, orders or customers?
	// first u need to create orders and then customers
	// customer can have many orders but orders can have 1 customer
	// var orders []models.Order
	// var customers []models.Customer

	// for i := 0; i < 100; i++ {
	// 	order := models.Order{
	// 		OrderID:   gofakeit.UintRange(1, 400),
	// 		OrderDate: gofakeit.Date(),
	// 		Status:    gofakeit.State(),
	// 	}
	// }

	for i := 0; i < 50; i++ {
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

		numItems := gofakeit.Number(1, 5)

		for j := 0; j < numItems; j++ {
			product := products[rand.Intn(len(products))]
			quantity := gofakeit.Number(1, 5)
			linePrice := product.RetailPrice.Price * float64(quantity)
			total += linePrice

			items = append(items, models.OrderItem{
				ProductID: product.ID,
				Quantity:  quantity,
				Price:     linePrice,
			})
		}
		order := models.Order{
			CustomerID: cstmr.ID,
			OrderDate:  gofakeit.Date(),
			Status:     status[rand.Intn(len(status))],
			Total:      total,
			Items:      items,
		}
		database.DB.Create(&order)

	}
	fmt.Println("Created 50 Customers and Orders")
	fmt.Println("Seeding complete!")

}
