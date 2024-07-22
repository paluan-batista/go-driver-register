package routes

import (
	"github.com/labstack/echo/v4"
	"go-driver-register/pkg/http"
)

func Register(e *echo.Echo) error {
	err := http.Health(e.AcquireContext())
	if err != nil {
		return err
	}
	return nil
}
