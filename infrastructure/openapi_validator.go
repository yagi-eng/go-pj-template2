package infrastructure

import (
	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/yagi-eng/go-pj-template2/apigen"
	"go.uber.org/zap"
)

func CreateOapiRequestValidator() echo.MiddlewareFunc {
	oapi, err := apigen.GetSwagger()
	if err != nil {
		zap.S().Fatalf("Cannot get openapi spec: %v", err)
	}
	oapi.Servers = nil
	opt := &oapimiddleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: openapi3filter.NoopAuthenticationFunc,
		},
	}
	return oapimiddleware.OapiRequestValidatorWithOptions(oapi, opt)
}
