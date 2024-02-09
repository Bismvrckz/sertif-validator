package routes

import (
	"cms-fello/app/config"
	view_controller "cms-fello/service/controller/web"
	"net/http"

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
	web := srv.Group(config.BaseURL, WebAuthMiddleware)

	web.GET("/login", view_controller.LoginView)
	web.GET("/markom", view_controller.DashboardView)

	// Artikel
	web.GET("/markom/artikel", view_controller.ListArtikelView)
	web.GET("/markom/artikel/tambah", view_controller.TambahArtikelView)
	web.GET("/markom/artikel/detail/id/:id_artikel", view_controller.DetailArtikelView)

	// Banner
	web.GET("/mobile/banner", view_controller.ListBannerView)
	web.GET("/mobile/banner/tambah", view_controller.TambahBannerView)
	web.GET("/mobile/banner/ubah/id/:id_banner", view_controller.UbahBannerView)
}
