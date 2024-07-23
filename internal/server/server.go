package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-driver-register/internal/container"
	"go-driver-register/internal/domain/entities"
	"go-driver-register/internal/server/routes"
)

func StartServer() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	container.NewContainer()
	err := container.PostgresDB.AutoMigrate(&entities.Driver{}, entities.Vehicle{})
	if err != nil {
		panic("failed to migrate")
	}

	err = routes.Register(e)
	if err != nil {
		e.Logger.Fatal()
	}
	routes.StartDriverRoutes(e)
	routes.StartVehicleRoutes(e)
}
