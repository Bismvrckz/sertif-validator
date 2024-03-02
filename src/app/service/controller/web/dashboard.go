package view_controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func DashboardView(ctx echo.Context) (err error) {
	htmlData := HtmlData{
		"prefix": base_url,
	}
	return ctx.Render(http.StatusOK, "public.dashboard", htmlData)
}

func AdminDashboardView(ctx echo.Context) (err error) {
	htmlData := HtmlData{
		"prefix":    base_url,
		"apiPrefix": base_url + "/api",
	}
	return ctx.Render(http.StatusOK, "admin.dashboard", htmlData)
}
