package routes

import (
	"errors"
	"net/http"
	"sertif_validator/app/config"
	"sertif_validator/app/logging"
	api_controller "sertif_validator/app/service/controller/api"
	view_controller "sertif_validator/app/service/controller/web"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
)

type (
	TkbaiApp struct {
		newEcho *echo.Echo
	}
)

func ApiAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		sess, err := session.Get("session", ctx)
		if err != nil {
			return err
		}

		if sess.Values["UserName"] == nil && ctx.Path() != config.BaseURL+"/api/login" {
			return err
		}

		if ctx.Path() == config.BaseURL+"/api/login" {
			return next(ctx)
		}

		ctx.Set("UserLevel", sess.Values["UserLevel"].(string))
		ctx.Set("UserName", sess.Values["UserName"].(string))
		ctx.Set("UserID", sess.Values["UserID"].(string))
		return next(ctx)
	}
}

func ServiceTKBAI() *TkbaiApp {
	return &TkbaiApp{
		newEcho: echo.New(),
	}
}

func (srv *TkbaiApp) Start(port string) {
	logger := logging.Log
	srv.newEcho.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
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
					Str("Host", values.Host).
					Str("RequestID", values.RequestID).
					Stack().Err(values.Error).Msg("")
			} else {
				logger.Info().
					Str("URI", values.URI).
					Str("METHOD", values.Method).
					Int("STATUS", values.Status).
					Str("IP", values.RemoteIP).
					Str("Host", values.Host).
					Str("RequestID", values.RequestID).
					Msg("Request")
			}
			return nil
		},
	}))

	srv.newEcho.HTTPErrorHandler = func(err error, ctx echo.Context) {
		var report *echo.HTTPError
		_ = errors.As(err, &report)

		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		err = ctx.JSON(report.Code, api_controller.Response{
			ResponseCode:    "02",
			AdditionalInfo:  report,
			ResponseMessage: "failed",
		})
		if err != nil {
			logger.Error().Err(err).Msg("")
		}
	}

	// recover
	srv.newEcho.Use(middleware.Recover())

	// set cors
	srv.newEcho.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PATCH, echo.PUT, echo.POST, echo.DELETE},
	}))

	srv.newEcho.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	// Web Routes Init
	view_controller.InitWeb(srv.newEcho)
	webRoutes(srv.newEcho)

	api := srv.newEcho.Group(config.BaseURL + "/api")

	api.GET("/auth/login", api_controller.LoginOIDC)
	api.GET("/auth/loginCallback", api_controller.LoginCallbackOIDC)
	api.GET("/auth/logout", api_controller.LogoutOIDC)
	api.GET("/auth/logoutCallback", api_controller.LogoutCallbackOIDC)
	api.GET("/auth/validate", api_controller.ValidateOIDC)

	// Admin
	api.GET("/admin/data/toefl/id/:test_id", api_controller.GetCertificateByID)
	api.GET("/admin/data/toefl/all", api_controller.GetCertificateAll)
	api.POST("/admin/data/toefl/csv", api_controller.PostCertificateCSV)

	// Certificate
	api.GET("/certificate/validate/id/:id", api_controller.ValidateCertificateByID)

	/*------------------------------------------ server start ------------------------------------------*/
	srv.newEcho.Logger.Fatal(srv.newEcho.Start(port))
}
