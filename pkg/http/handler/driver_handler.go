package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-driver-register/internal/domain/entities"
	"go-driver-register/internal/usecases"
	"net/http"
)

type DriverHandler struct {
	usecase usecases.DriverUsecase
}

func NewDriverController(usecase usecases.DriverUsecase) *DriverHandler {
	return &DriverHandler{usecase}
}

func (c *DriverHandler) CreateDriver(ctx echo.Context) error {
	var driver entities.Driver
	if err := ctx.Bind(&driver); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := c.usecase.CreateDriver(&driver); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, driver)
}

func (c *DriverHandler) GetAllDrivers(ctx echo.Context) error {
	drivers, err := c.usecase.GetAllDrivers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, drivers)
}

func (c *DriverHandler) GetDriverByID(ctx echo.Context) error {
	id := uuid.MustParse(ctx.Param("id"))
	driver, err := c.usecase.GetDriverByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, driver)
}

func (c *DriverHandler) UpdateDriver(ctx echo.Context) error {
	var driver entities.Driver
	if err := ctx.Bind(&driver); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := c.usecase.UpdateDriver(&driver); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, driver)
}

func (c *DriverHandler) DeleteDriver(ctx echo.Context) error {
	id := uuid.MustParse(ctx.Param("id"))

	if err := c.usecase.DeleteDriver(id); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusNoContent)
}
