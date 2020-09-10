package service

import (
	"github.com/labstack/echo/v4"
)

func getData(c echo.Context) (data map[string]interface{}) {
	aux := c.Get("data")

	if aux == nil {
		aux = map[string]interface{}{}
	}

	return aux.(map[string]interface{})
}
