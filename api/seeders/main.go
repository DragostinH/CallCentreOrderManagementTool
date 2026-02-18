package seeders

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/DragostinH/CallCentreOrderManagementTool/database"
	"github.com/DragostinH/CallCentreOrderManagementTool/models"
	"github.com/DragostinH/CallCentreOrderManagementTool/utils"
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

func SeedCustomersWithOrders() {
	fmt.Println("Seeding customers with orders...")

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

		// 2. Create the Order
		order := models.Order{
			CustomerID:  cstmr.ID,
			OrderNumber: uint(gofakeit.IntRange(000000000, 999999999)),
			OrderDate:   gofakeit.Date(),
			Status:      statusOptions[rand.Intn(len(statusOptions))],
		}
		database.DB.Create(&order)

		// 3. Generate Items for an Order
		numItems := gofakeit.Number(1, 50)
		var currentOrderItems []models.OrderItem
		var orderTotal float64

		for j := 0; j < numItems; j++ {
			product := allProducts[rand.Intn(len(allProducts))]
			quantity := gofakeit.Number(1, 5)
			linePrice := utils.RoundFloat(product.RetailPrice.Price*float64(quantity), 2)
			orderTotal += linePrice

			currentOrderItems = append(currentOrderItems, models.OrderItem{
				OrderID:   order.ID,
				ProductID: product.ID,
				Quantity:  quantity,
				Price:     linePrice,
			})
		}

		database.DB.Create(&currentOrderItems)
		order.Total = utils.RoundFloat(orderTotal, 2)
		database.DB.Save(&order)

	}
	fmt.Println("Created 50 Customers and their Orders.")
}
