package gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/wendermrn/barber-shop-go-api/pkg/dao"
	"github.com/wendermrn/barber-shop-go-api/pkg/database"
	"github.com/wendermrn/barber-shop-go-api/pkg/model"
)

var dbo *gorm.DB

//CustomerDAO struct
type CustomerDAO struct {
	dao.DAO
}

func init() {
	if dbo == nil {
		d, err := database.GetDBG()
		if err != nil {
			log.Error("Connect Error")
			return
		}
		dbo = d
	}
}

//Create a Customer
func (c *CustomerDAO) Create(u *model.Customer) error {
	return dbo.Save(&u).Error
}

//Update a Customer
func (c *CustomerDAO) Update(u *model.Customer) error {
	_, err := c.FindByID(u.ID)
	if err != nil {
		return err
	}
	return dbo.Save(&u).Error
}

//FindByID a Customer
func (c *CustomerDAO) FindByID(ID uint64) (*model.Customer, error) {
	customer := model.Customer{}
	return &customer, dbo.Find(&customer, ID).Error
}

//Delete a Customer
func (c *CustomerDAO) Delete(u *model.Customer) error {
	return dbo.Delete(&u).Error
}
