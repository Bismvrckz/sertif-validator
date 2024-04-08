package config

import (
	"github.com/labstack/echo/v4"
)

type Apps struct {
	Web *echo.Echo
}
