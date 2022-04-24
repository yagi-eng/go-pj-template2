package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Healthz health check
func Healthz(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
