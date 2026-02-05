package seeders

import (
	"fmt"
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

func SeedCustomersWithOrders() {}
