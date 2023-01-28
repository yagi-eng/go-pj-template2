package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagi-eng/go-pj-template2/apigen"
)

func (c *controller) PostUsers(ctx echo.Context) error {
	var req apigen.PostUsersJSONRequestBody
	if err := bindAndValidate(ctx, &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return ctx.NoContent(http.StatusCreated)
}

func (c *controller) PutUsers(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (c *controller) GetUsersUserId(ctx echo.Context, userId apigen.UserIdPath) error {
	return ctx.JSON(http.StatusOK, userId)
}
