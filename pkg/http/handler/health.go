package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Health(c echo.Context) error {
	response := map[string]string{
		"status": "UP",
	}
	return c.JSON(http.StatusOK, response)
}
