package route

import (
	echo "github.com/labstack/echo/v4"
	"github.com/wendermrn/barber-shop-go-api/pkg/service"
)

// Setup carrega todas as rotas da aplicação
func Setup(e *echo.Echo) {

	// Customers
	customerServ := service.CustomerServ{}
	groupCustomers := e.Group("/customers")
	groupCustomers.POST("", customerServ.Create)
	groupCustomers.PUT("", customerServ.Update)
	groupCustomers.GET("/:id", customerServ.Get)
	groupCustomers.DELETE("/:id", customerServ.Delete)
}
