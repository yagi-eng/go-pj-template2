package infrastructure

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/yagi-eng/go-pj-template2/apigen"
	"github.com/yagi-eng/go-pj-template2/domain/myerror"
	"go.uber.org/zap"
)

func newErrorResp(msg string) *apigen.Error {
	// don't return error detail for security in PRD
	if os.Getenv("IS_PRD") == "true" {
		return nil
	}
	return &apigen.Error{
		Message: msg,
	}
}

// CustomHTTPErrorHandler HTTP errorをハンドリングする
func CustomHTTPErrorHandler(err error, c echo.Context) {
	unknownErrResp := newErrorResp("Unknown error")

	me, ok := err.(*myerror.MyError)
	if ok {
		switch {
		case me.Code >= 200:
			zap.S().Errorf("Server error: %+v", me.Err)
			c.JSON(http.StatusInternalServerError, newErrorResp(me.Error()))
			return
		case me.Code == 101:
			zap.S().Infof("Client error: %v", me.Err)
			c.JSON(http.StatusNotFound, newErrorResp(me.Error()))
			return
		case me.Code >= 100:
			zap.S().Infof("Client error: %v", me.Err)
			c.JSON(http.StatusBadRequest, newErrorResp(me.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, unknownErrResp)
		return
	}

	he, ok := err.(*echo.HTTPError)
	if ok {
		if msg, ok := he.Message.(string); ok {
			httpCode := he.Code
			// echo error
			if httpCode == http.StatusNotFound {
				c.JSON(http.StatusNotFound, newErrorResp(msg))
				return
			}
			if httpCode == http.StatusMethodNotAllowed {
				c.JSON(http.StatusMethodNotAllowed, newErrorResp(msg))
				return
			}
			// oapi-codegen error
			zap.S().Infof("oapi-codegen error: %v", he)
			c.JSON(http.StatusBadRequest, newErrorResp(msg))
			return
		}
	}

	// don't expect reaching here
	zap.S().Errorf("Unknown HTTP error: %v", he)
	c.JSON(http.StatusInternalServerError, unknownErrResp)
}
