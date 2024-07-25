package routes

import (
	"github.com/labstack/echo/v4"
	"go-driver-register/internal/container"
	"go-driver-register/internal/repository"
	"go-driver-register/internal/usecases"
	"go-driver-register/pkg/http/handler"
)

func StartDriverRoutes(e *echo.Echo) {
	db := container.GetDB()
	driverRepository := repository.NewDriverRepository(db)
	driverUsecase := usecases.NewDriverUsecase(driverRepository)
	driverHandler := handler.NewDriverController(driverUsecase)
	e.POST("/drivers", driverHandler.CreateDriver)
	e.GET("/drivers", driverHandler.GetAllDrivers)
	e.GET("/drivers/:id", driverHandler.GetDriverByID)
	e.PUT("/drivers/:id", driverHandler.UpdateDriver)
	e.DELETE("/drivers/:id", driverHandler.DeleteDriver)
}
