package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagi-eng/go-pj-template2/apigen"
	"github.com/yagi-eng/go-pj-template2/util"
)

func (c *controller) GetTest(ctx echo.Context, params apigen.GetTestParams) error {
	o := apigen.TestObject{
		Id:   util.Pointer(10),
		Name: util.Pointer("Mike"),
		Q:    util.Pointer(params.Q),
	}
	return ctx.JSON(http.StatusOK, o)
}
