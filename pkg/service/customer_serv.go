package service

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/wendermrn/barber-shop-go-api/pkg/controller"
	"github.com/wendermrn/barber-shop-go-api/pkg/model"
)

//CustomerServ service struct
type CustomerServ struct {
}

//Create create Customer
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
	return c.JSON(http.StatusOK, customer)
}
