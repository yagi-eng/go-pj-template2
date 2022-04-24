package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// CustomHTTPErrorHandler HTTP errorをハンドリングする
func CustomHTTPErrorHandler(err error, c echo.Context) {
	unknownErrRsp := errorResp{
		Message: "unknown error",
	}

	he, ok := err.(*echo.HTTPError)
	if !ok {
		// panicなど予期せぬエラーが発生するとここに到達する
		zap.S().Errorf("Unknown error: %v", err)
		c.JSON(http.StatusInternalServerError, unknownErrRsp)
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

		errRsp := errorResp{
			Message: "something wrong",
		}
		c.JSON(httpCode, errRsp)
	case string:
		// echo側でエラーが発生するとこのcase文に入る

		if httpCode == http.StatusNotFound {
			errRsp := errorResp{
				Message: "url not found",
			}
			c.JSON(http.StatusNotFound, errRsp)
			return
		}
		// 現状ここに到達する想定はない
		zap.S().Errorf("Echo HTTP error: %v", he)
		c.JSON(http.StatusInternalServerError, unknownErrRsp)
	default:
		// 通常到達しない

		zap.S().Errorf("Unknown HTTP error: %v", he)
		c.JSON(http.StatusInternalServerError, unknownErrRsp)
	}
}
