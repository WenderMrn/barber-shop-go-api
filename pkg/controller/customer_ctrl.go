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
	customerDAO := gorm.CustomerDAO{}
	return customerDAO.Create(u)
}

//Update a Customer
func (ctrl *CustomerCtrl) Update(u *model.Customer) error {
	customerDAO := gorm.CustomerDAO{}
	return customerDAO.Update(u)
}

//Get a Customer
func (ctrl *CustomerCtrl) Get(ID uint64) (*model.Customer, error) {
	customerDAO := gorm.CustomerDAO{}
	customer, err := customerDAO.FindByID(ID)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

//Delete a Customer
func (ctrl *CustomerCtrl) Delete(ID uint64) error {
	customerDAO := gorm.CustomerDAO{}
	customer, err := customerDAO.FindByID(ID)
	if err != nil {
		return err
	}
	return customerDAO.Delete(customer)
}
