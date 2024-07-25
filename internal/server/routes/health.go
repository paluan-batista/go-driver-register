package routes

import (
	"github.com/labstack/echo/v4"
	"go-driver-register/pkg/http/handler"
)

func SetUpHealthRoute(h *echo.Echo) {
	h.GET("/health", handler.Health)
}
