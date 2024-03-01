package routes

import (
	"net/http"
	"sertif_validator/app/config"
	view_controller "sertif_validator/app/service/controller/web"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func WebAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		sess, err := session.Get("session", ctx)
		if err != nil {
			return err
		}

		if sess.Values["UserName"] == nil && ctx.Path() != config.BaseURL+"/login" {
			return ctx.Redirect(http.StatusSeeOther, config.BaseURL+"/login")
		}

		if ctx.Path() == config.BaseURL+"/login" {
			return next(ctx)
		}

		ctx.Set("UserLevel", sess.Values["UserLevel"].(string))
		ctx.Set("UserName", sess.Values["UserName"].(string))
		ctx.Set("UserID", sess.Values["UserID"].(string))
		return next(ctx)
	}
}

func webRoutes(srv *echo.Echo) {

	/*------------------------------------------ VIEWS ------------------------------------------*/
	web := srv.Group(config.BaseURL)

	// web.GET("/login", view_controller.LoginView)
	web.GET("/", view_controller.DashboardView)

	web.GET("/admin", view_controller.LoginView)
	web.GET("/admin/dash", view_controller.AdminDashboardView)
}
