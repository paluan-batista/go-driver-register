package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "App is healthy")
}
