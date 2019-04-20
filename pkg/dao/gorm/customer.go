package gorm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/wendermrn/barber-shop-go-api/pkg/dao"
	"github.com/wendermrn/barber-shop-go-api/pkg/database"
	"github.com/wendermrn/barber-shop-go-api/pkg/model"
)

var dbo *gorm.DB

//CustomerDAO struct
type CustomerDAO struct {
	dao.IGenericDAO
}

func init() {
	fmt.Println("sakjskajska")
	if dbo == nil {
		d, err := database.GetDBG()
		if err != nil {
			log.Error("Connect Error")
			return
		}
		dbo = d
	}
}

//Create create a Customer
func (c *CustomerDAO) Create(u *model.Customer) error {
	return dbo.Create(&u).Error
}

//FindByID create a Customer
func (c *CustomerDAO) FindByID(ID uint) (*model.Customer, error) {
	customer := model.Customer{}
	err := dbo.Find(&customer, ID).Error
	return &customer, err
}
