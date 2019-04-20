package model

//Customer struct
type Customer struct {
	User
	HairType      string         `json:"hair_type"`
	ServiceOrders []ServiceOrder `gorm:"foreignkey:CustomerID"`
}
