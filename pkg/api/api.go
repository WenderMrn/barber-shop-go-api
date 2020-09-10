package api

import (
	"context"
	"net"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wendermrn/barber-shop-go-api/pkg/route"
)

// Run start api server
func Run(ctx context.Context) error {
	e := echo.New()

	// Setup middreware
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	// Configura as rotas da api
	route.Setup(e)

	return e.Start(net.JoinHostPort("localhost", "3003"))
}
