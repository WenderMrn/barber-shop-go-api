package controller

import (
	"github.com/wendermrn/barber-shop-go-api/pkg/dao/gorm"
	"github.com/wendermrn/barber-shop-go-api/pkg/model"
)

//CustomerCtrl customer controller struct
type CustomerCtrl struct {
}

//Create a Customer
func (ctrl *CustomerCtrl) Create(u *model.Customer) error {
	customerDao := gorm.CustomerDAO{}
	return customerDao.Create(u)
}
