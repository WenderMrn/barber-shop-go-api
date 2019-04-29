package service

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/wendermrn/barber-shop-go-api/pkg/controller"
	"github.com/wendermrn/barber-shop-go-api/pkg/model"
)

//CustomerServ service struct
type CustomerServ struct {
}

//Create a Customer
func (ctrl *CustomerServ) Create(c echo.Context) (err error) {
	data := getData(c)

	customerCtrl := controller.CustomerCtrl{}
	customer := model.Customer{}

	err = c.Bind(&customer)
	if err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	err = customerCtrl.Create(&customer)
	if err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusInternalServerError, data)
	}
	return c.JSON(http.StatusCreated, customer)
}

//Get a Customer
func (ctrl *CustomerServ) Get(c echo.Context) (err error) {
	data := getData(c)

	customerCtrl := controller.CustomerCtrl{}

	customerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	customer, err := customerCtrl.Get(customerID)
	if err == gorm.ErrRecordNotFound {
		data["message"] = "Customer not found"
		return c.JSON(http.StatusNotFound, data)
	} else if err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusInternalServerError, data)
	}
	return c.JSON(http.StatusOK, customer)
}

//Update a Customer
func (ctrl *CustomerServ) Update(c echo.Context) (err error) {
	data := getData(c)

	customerCtrl := controller.CustomerCtrl{}
	customer := model.Customer{}

	err = c.Bind(&customer)
	if err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	err = customerCtrl.Update(&customer)
	if err == gorm.ErrRecordNotFound {
		data["message"] = "Customer not found"
		return c.JSON(http.StatusNotFound, data)
	} else if err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusInternalServerError, data)
	}
	return c.JSON(http.StatusOK, customer)
}

//Delete a Customer
func (ctrl *CustomerServ) Delete(c echo.Context) (err error) {
	data := getData(c)

	customerCtrl := controller.CustomerCtrl{}

	customerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, data)
	}

	err = customerCtrl.Delete(customerID)
	if err == gorm.ErrRecordNotFound {
		data["message"] = "Customer not found"
		return c.JSON(http.StatusNotFound, data)
	} else if err != nil {
		data["message"] = err.Error()
		return c.JSON(http.StatusInternalServerError, data)
	}
	data["message"] = "Customer deleted successfully"
	return c.JSON(http.StatusOK, data)
}
