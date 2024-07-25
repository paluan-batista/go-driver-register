package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-driver-register/internal/container"
	"go-driver-register/internal/server/routes"
	"go-driver-register/internal/settings"
)

func StartServer() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	container.InitContainer()

	routes.SetUpHealthRoute(e)
	routes.StartDriverRoutes(e)
	routes.StartVehicleRoutes(e)

	e.Logger.Fatal(e.Start(settings.GetServerConfig().Host))
}
