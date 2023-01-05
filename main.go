package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yagi-eng/go-pj-template2/apigen"
	"github.com/yagi-eng/go-pj-template2/controller"
	"github.com/yagi-eng/go-pj-template2/infrastructure"
	"github.com/yagi-eng/go-pj-template2/repository"
	"github.com/yagi-eng/go-pj-template2/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// okay to use 8080 in prd, because Render detects which port is open
const port = 8080

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("cannot read .env")
	}

	e := echo.New()
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = util.CustomHTTPErrorHandler
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("FRONT_DOMAIN")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			m, p := c.Request().Method, c.Request().URL.Path
			return (m == "GET" && p == "/healthz")
		},
	}))

	var logger *zap.Logger
	if os.Getenv("IS_PRD") == "true" {
		zapConfig := zap.NewProductionConfig()
		var level zapcore.Level
		level.Set(os.Getenv("LOGGER_LEVEL"))
		zapConfig.Level = zap.NewAtomicLevelAt(level)
		zapConfig.DisableStacktrace = true
		logger, _ = zapConfig.Build()
	} else {
		zapConfig := zap.NewDevelopmentConfig()
		zapConfig.DisableStacktrace = true
		logger, _ = zapConfig.Build()
	}

	defer logger.Sync()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	// DB
	db := infrastructure.Connect()
	repository.Migrate(db)

	handler := controller.NewController(db)
	apigen.RegisterHandlers(e, handler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
