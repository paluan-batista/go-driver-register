package routes

import (
	"github.com/labstack/echo/v4"
	"go-driver-register/internal/container"
	"go-driver-register/internal/repository"
	"go-driver-register/internal/usecases"
	"go-driver-register/pkg/http/handler"
)

func StartVehicleRoutes(e *echo.Echo) {
	db := container.GetDB()
	vehicleUsecase := usecases.NewVehicleUsecase(repository.NewVehicleRepository(db), repository.NewDriverRepository(db))
	vehicleHandler := handler.NewVehicleHandler(vehicleUsecase)
	e.POST("/vehicles", vehicleHandler.CreateVehicle)
	e.GET("/vehicles", vehicleHandler.GetAllVehicles)
	e.GET("/vehicles/:id", vehicleHandler.GetVehicleByID)
	e.PUT("/vehicles/:id", vehicleHandler.UpdateVehicle)
	e.DELETE("/vehicles/:id", vehicleHandler.DeleteVehicle)
	e.POST("/vehicles/:vehicle_id/assign/:driver_id", vehicleHandler.AssignDriver)
}
