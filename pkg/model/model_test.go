package model

import (
	"testing"

	"github.com/wendermrn/barber-shop-go-api/pkg/database"
)

func TestCreateService(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Create a service", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbo, err := database.GetDBG()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDBG() error = %v", err.Error())
			}

			dbo.AutoMigrate(&User{}, &ServiceOrder{}, &Service{})

			user := User{Name: "Wender", Email: "wender.mrn@gmail.com"}

			dbo.Create(&user)

			service := Service{Title: "Pintura", Description: "Pinturas em geral", UnitPrice: 100.00, Owner: user}

			dbo.Create(&service)

			services := []Service{}
			dbo.Preload("Owner").Preload("ServiceOrder").First(&services)
			t.Logf("Services: %+v", services)

		})
	}
}
