package view_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func LoginView(ctx echo.Context) error {
	htmlData := HtmlData{
		"prefix": base_url,
	}

	return ctx.Render(http.StatusOK, "admin.login", htmlData)
}
