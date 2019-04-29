package model

import (
	"time"
)

//Service model
type Service struct {
	ID          uint64  `json:"id" gorm:"primary_key"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UnitPrice   float64 `json:"unit_price"`
	Active      bool    `json:"active" gorm:"default:true"`

	OwnerID uint
	Owner   User

	ServiceOrders []ServiceOrder `json:"service_orders" gorm:"many2many:service_order_service"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//ServiceOrder model
type ServiceOrder struct {
	ID          uint64    `json:"id" gorm:"primary_key"`
	Status      string    `json:"status"`
	Quantity    int       `json:"quantity"`
	Observation string    `json:"observation"`
	PaymentDate time.Time `json:"payment_date"`
	TotalPrice  float64   `json:"total_price"`

	BarberID uint64
	Barber   User `json:"barber"`

	CustomerID uint64
	Customer   User `json:"customer"`

	Services []Service `json:"services" gorm:"many2many:service_order_service"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
