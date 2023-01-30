package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/yagi-eng/go-pj-template2/domain/myerror"
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

func bindAndValidate(c echo.Context, req interface{}) error {
	if err := c.Bind(req); err != nil {
		return myerror.New(myerror.BadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return myerror.New(myerror.BadRequest, err)
	}
	return nil
}
