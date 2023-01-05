package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type controller struct {
	// TODO have usecase struct
	db *gorm.DB
}

func NewController(
	db *gorm.DB,
) *controller {
	return &controller{
		db: db,
	}
}

func errWithStack(code int, err error) error {
	return echo.NewHTTPError(code, errors.WithStack(err))
}

func bindAndValidate(c echo.Context, req interface{}) error {
	if err := c.Bind(req); err != nil {
		return errWithStack(http.StatusBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return errWithStack(http.StatusBadRequest, err)
	}
	return nil
}
