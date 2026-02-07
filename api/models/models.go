package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	PostCode string `json:"post_code"`
	City     string `json:"city"`
	Street   string `json:"street"`
}

type PriceInfo struct {
	Price         float64 `json:"price"`
	Measure       string  `json:"measure"`
	MeasureAmount int     `json:"measure_amount"`
}

type Customer struct {
	gorm.Model
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	Phone          string  `gorm:"unique" json:"phone"`
	Email          string  `gorm:"unique" json:"email"`
	Address        Address `gorm:"embedded;embeddedPrefix:address_"`
	CustomerNumber string  `gorm:"unique" json:"customer_number"`
	Orders         []Order `json:"orders,omitempty"`
}

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

type Product struct {
	gorm.Model
	ProductUID  string     `gorm:"unique" json:"product_uid"`
	Name        string     `gorm:"unique" json:"name"`
	Image       string     `json:"image"`
	ProductType string     `json:"product_type"`
	Eans        string     `json:"eans"`
	IsAlcoholic bool       `json:"is_alcoholic"`
	UnitPrice   PriceInfo  `gorm:"embedded;embeddedPrefix:unit_"`
	RetailPrice PriceInfo  `gorm:"embedded;embeddedPrefix:retail_"`
	Categories  []Category `gorm:"many2many:product_categories;" json:"categories"`
}

type Order struct {
	gorm.Model
	CustomerID uint        `json:"customer_id"`
	OrderID    uint        `json:"customer_id"`
	OrderDate  time.Time   `json:"order_date"`
	Status     string      `json:"status"`
	Total      float64     `json:"total"`
	Items      []OrderItem `json:"items"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"unique" json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
