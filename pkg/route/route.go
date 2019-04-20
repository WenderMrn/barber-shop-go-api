package route

import (
	"github.com/labstack/echo"
	"github.com/wendermrn/barber-shop-go-api/pkg/service"
)

// Setup carrega todas as rotas da aplicação
func Setup(e *echo.Echo) {

	// Customers
	customerServ := service.CustomerServ{}
	groupCustomers := e.Group("/customers")
	groupCustomers.POST("", customerServ.Create)
}
