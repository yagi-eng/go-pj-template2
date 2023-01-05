package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Healthz health check
func (c *controller) GetHealthz(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}
