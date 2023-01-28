package infrastructure

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// CustomHTTPErrorHandler HTTP errorをハンドリングする
func CustomHTTPErrorHandler(err error, c echo.Context) {
	unknownErrResp := errorResp{
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
		switch {
		case httpCode >= 500:
			zap.S().Errorf("Server error: %+v", err)
		case httpCode >= 400:
			zap.S().Infof("Client error: %v", err)
		}
		errResp := errorResp{
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResp)
	case string:
		// echo error
		if httpCode == http.StatusNotFound {
			c.String(http.StatusNotFound, "Not Found")
			return
		}
		if httpCode == http.StatusMethodNotAllowed {
			c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
			return
		}
		// don't expect reaching here
		zap.S().Errorf("Echo HTTP error: %v", he)
		c.JSON(http.StatusInternalServerError, unknownErrResp)
	default:
		// don't expect reaching here
		zap.S().Errorf("Unknown HTTP error: %v", he)
		c.JSON(http.StatusInternalServerError, unknownErrResp)
	}
}
