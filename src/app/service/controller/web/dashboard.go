package view_controller

import (
	"fmt"
	"net/http"
	"sertif_validator/app/config"
	"sertif_validator/app/service/handler"

	"github.com/labstack/echo/v4"
)

func DashboardView(ctx echo.Context) (err error) {
	htmlData := HtmlData{
		"prefix": base_url,
	}
	return ctx.Render(http.StatusOK, "public.dashboard", htmlData)
}

func AdminDashboardView(ctx echo.Context) (err error) {
	idToken, err := handler.ReadCookie(ctx, "idToken")
	if err != nil {
		fmt.Println(err.Error())
		return ctx.Redirect(http.StatusSeeOther, config.BaseURL+"/login")
	}

	htmlData := HtmlData{
		"prefix":    base_url,
		"apiPrefix": config.BaseURL + "/api",
		"idToken":   idToken.Value,
	}
	return ctx.Render(http.StatusOK, "admin.dashboard", htmlData)
}
