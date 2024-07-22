package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-driver-register/internal/domain/entities"
	"go-driver-register/internal/usecases"
	"net/http"
)

type VehicleHandler struct {
	usecase usecases.VehicleUsecase
}

func NewVehicleHandler(usecase usecases.VehicleUsecase) *VehicleHandler {
	return &VehicleHandler{usecase}
}

func (c *VehicleHandler) CreateVehicle(ctx echo.Context) error {
	var vehicle entities.Vehicle
	if err := ctx.Bind(&vehicle); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := c.usecase.CreateVehicle(&vehicle); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, vehicle)
}

func (c *VehicleHandler) GetAllVehicles(ctx echo.Context) error {
	vehicles, err := c.usecase.GetAllVehicles()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, vehicles)
}

func (c *VehicleHandler) GetVehicleByID(ctx echo.Context) error {
	id := uuid.MustParse(ctx.Param("id"))
	vehicle, err := c.usecase.GetVehicleByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, vehicle)
}

func (c *VehicleHandler) UpdateVehicle(ctx echo.Context) error {
	var vehicle entities.Vehicle
	if err := ctx.Bind(&vehicle); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := c.usecase.UpdateVehicle(&vehicle); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, vehicle)
}

func (c *VehicleHandler) DeleteVehicle(ctx echo.Context) error {
	id := uuid.MustParse(ctx.Param("id"))
	if err := c.usecase.DeleteVehicle(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (c *VehicleHandler) AssignDriver(ctx echo.Context) error {

	vehicleID := uuid.MustParse(ctx.Param("vehicle_id"))

	driverID := uuid.MustParse(ctx.Param("driver_id"))
	if err := c.usecase.AssignDriver(vehicleID, driverID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Driver assigned successfully"})
}
