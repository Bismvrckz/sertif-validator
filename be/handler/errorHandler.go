package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"tkbai-be/config"
)

func InitErrHandler(ein *config.Apps) {
	loggers := config.Log

	// return route response
	//var code int
	//var response models.Status

	ein.Api.HTTPErrorHandler = func(err error, ctx echo.Context) {
		loggers.Debug().Msg(err.Error())

		var report *echo.HTTPError
		_ = errors.As(err, &report)

		err = ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": err.Error(),
		})
		if err != nil {
			loggers.Error().Err(err).Msg("")
		}
	}
}
