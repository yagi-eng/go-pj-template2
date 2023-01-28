package infrastructure

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagi-eng/go-pj-template2/apigen"
	"go.uber.org/zap"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	unknownErrResp := apigen.Error{
		Code:    "500-000",
		Message: "Unknown error",
	}

	he, ok := err.(*echo.HTTPError)
	if !ok {
		// e.g. panic
		zap.S().Errorf("Unknown error: %v", err)
		c.JSON(http.StatusInternalServerError, unknownErrResp)
		return
	}

	httpCode := he.Code
	switch err := he.Message.(type) {
	case error:
		var errResp apigen.Error
		switch {
		case httpCode >= 500:
			zap.S().Errorf("Server error: %+v", err)
			errResp = apigen.Error{
				Code:    "500-000", // FIXME
				Message: err.Error(),
			}
		case httpCode >= 400:
			zap.S().Infof("Client error: %v", err)
			errResp = apigen.Error{
				Code:    "400-000", // FIXME
				Message: err.Error(),
			}
		}
		c.JSON(http.StatusInternalServerError, errResp)
	case string:
		// echo error
		if httpCode == http.StatusNotFound {
			errResp := apigen.Error{
				Code:    "404-000",
				Message: he.Message.(string),
			}
			c.JSON(http.StatusNotFound, errResp)
			return
		}
		if httpCode == http.StatusMethodNotAllowed {
			errResp := apigen.Error{
				Code:    "405-000",
				Message: he.Message.(string),
			}
			c.JSON(http.StatusMethodNotAllowed, errResp)
			return
		}
		// oapi-codegen error
		zap.S().Infof("oapi-codegen error: %v", he)
		errResp := apigen.Error{
			Code:    "400-000",
			Message: he.Message.(string),
		}
		c.JSON(http.StatusBadRequest, errResp)
	default:
		// don't expect reaching here
		zap.S().Errorf("Unknown HTTP error: %v", he)
		c.JSON(http.StatusInternalServerError, unknownErrResp)
	}
}
