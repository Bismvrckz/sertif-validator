package main

import (
	"log"
	"tkbai-be/config"
	"tkbai-be/databases"
	"tkbai-be/handler"
	"tkbai-be/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	a := new(config.Apps)

	a.Api = echo.New()

	// recover
	a.Api.Use(middleware.Recover())

	// set cors
	a.Api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{config.WebHost, config.APIHost},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PATCH, echo.PUT, echo.POST, echo.DELETE},
	}))

	//logging
	initLoggingMiddleware(a)

	//init handler
	handler.InitErrHandler(a)

	//add routes
	routes.BuildRoutes(a)

	err := databases.ConnectTkbaiDatabase()
	if err != nil {
		log.Fatal(err)
	}

	a.Api.Logger.Fatal(a.Api.Start(config.SERVERPort))
}

func initLoggingMiddleware(ein *config.Apps) {
	logger := config.Log
	// logging route

	ein.Api.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogURIPath:   true,
		LogStatus:    true,
		LogRemoteIP:  true,
		LogHost:      true,
		LogRequestID: true,
		LogError:     true,
		LogMethod:    true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			if values.Error != nil {
				logger.Error().
					Str("URI", values.URI).
					Str("METHOD", values.Method).
					Int("STATUS", values.Status).
					Str("IP", values.RemoteIP).
					Str("HOST", values.Host).
					Str("RequestID", values.RequestID).
					Stack().Err(values.Error).Msg("")
			} else {
				logger.Info().
					Str("URI", values.URI).
					Str("METHOD", values.Method).
					Int("STATUS", values.Status).
					Str("IP", values.RemoteIP).
					Str("HOST", values.Host).
					Str("RequestID", values.RequestID).
					Msg("Request")
			}
			return nil
		},
	}))
}
