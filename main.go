package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yagi-eng/line-pay-nft/handler"
	"github.com/yagi-eng/line-pay-nft/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/time/rate"
)

// okay to use 8080 in prd, because Render detects which port is open
const port = 8080

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("cannot read .env")
	}

	e := echo.New()
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Logger())
	e.HTTPErrorHandler = util.CustomHTTPErrorHandler
	router(e)

	var logger *zap.Logger
	if os.Getenv("IS_LOCAL") == "true" {
		e.Use(middleware.CORS())
		zapConfig := zap.NewDevelopmentConfig()
		zapConfig.DisableStacktrace = true
		logger, _ = zapConfig.Build()
	} else {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{os.Getenv("FRONT_URL")},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))
		zapConfig := zap.NewProductionConfig()
		var level zapcore.Level
		level.Set(os.Getenv("LOGGER_LEVEL"))
		zapConfig.Level = zap.NewAtomicLevelAt(level)
		zapConfig.DisableStacktrace = true
		logger, _ = zapConfig.Build()
		limit, _ := strconv.Atoi(os.Getenv("RATE_LIMIT"))
		e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(limit))))
	}

	defer logger.Sync()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func router(e *echo.Echo) {
	e.GET("/healthz", handler.Healthz)
}
