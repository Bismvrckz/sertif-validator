package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tkbai-fe/config"
)

func AdminLoginView(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "admin.login", map[string]interface{}{
		"apiHost":   config.APIHost,
		"apiPrefix": config.ApiPrefix,
		"prefix":    config.BaseURL,
	})
}

func AdminDashboardView(ctx echo.Context) (err error) {
	htmlData := ctx.Get("htmlData").(map[string]interface{})

	return ctx.Render(http.StatusOK, "admin.dashboard", htmlData)
}
