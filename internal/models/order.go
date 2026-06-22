package models

import (
    "gorm.io/gorm"
    "github.com/teris-io/shortid"
)


var (
	OrderStatuses = []string{"Order placed", "Preparing", "Baking", "Quality Check", "Ready"}
	PizzaTypes    = []string{
		"Margherita",
		"Pepperoni",
		"BBQ Chicken",
		"Veggie",
		"Hawaiian",
		"Meat Lovers",
		"Supreme",
		"Buffalo Chicken",
		"Four Cheese",
		"Spinach and Feta",
		"Mediterranean",
		"White Pizza",
		"Philly Cheesesteak",
		"Truffle Mushroom",
		"Caprese",
		"Breakfast Pizza",
	}
	PizzaSizes = []string{
		"Small",
		"Medium",
		"Large",
		"Extra Large",
	}
)

type OrderModel struct {
	DB *gorm.DB
}
type Order struct {
	ID           string      `json:"id" gorm:"primaryKey;size:14"`
	Status       string      `json:"status" gorm:"not null"`
	CustomerName string      `json:"customer_name" gorm:"not null"`
	Phone        string      `json:"phone" gorm:"not null"`
	Address      string      `json:"address" gorm:"not null"`
	Items        []OrderItem `json:"pizzas" gorm:"foreignKey:OrderID;"`
	CreatedAt    int64       `json:"created_at" gorm:"autoCreateTime"`
}
type OrderItem struct {
	ID           string `json:"id" gorm:"primaryKey;size:14"`
	OrderID      string `json:"order_id" gorm:"index"`
	Size         string `json:"size" gorm:"not null"`
	Pizza        string `json:"pizza" gorm:"not null"`
	Instructions string `json:"instructions"`
}
func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID == "" {
		o.ID = shortid.MustGenerate()
	}
	return nil
}
func (oi *OrderItem) BeforeCreate(tx *gorm.DB) error {
	if oi.ID == "" {
		oi.ID = shortid.MustGenerate()
	}
	return nil
}
func (o * OrderModel) CreateOrder (order *Order)error {
	return o.DB.Create(order).Error
}
func (o *OrderModel) GetOrder(id string) (*Order, error) {
	var order Order
	err := o.DB.Preload("Items").First(&order, "id = ?", id).Error
	return &order, err
}
