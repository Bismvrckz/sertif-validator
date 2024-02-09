package routes

import (
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
	VALIDATOR struct {
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

func ServiceVALIDATOR() *VALIDATOR {
	return &VALIDATOR{
		newEcho: echo.New(),
	}
}

func (srv *VALIDATOR) Start(port string) {
	/**=======================================================================================================================
	*?                                                   CONFIG
	*=======================================================================================================================**/

	/*------------------------------------------ middleware ------------------------------------------*/

	// logging route
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

	srv.newEcho.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			var respn interface{} = &api_controller.Response{
				Rc:   "02",
				Val:  report,
				Desc: "gagal",
			}

			c.JSON(report.Code, respn)

		} else {
			c.JSON(report.Code, report)
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

	api := srv.newEcho.Group(config.BaseURL+"/api", ApiAuthMiddleware)

	// api.POST("/login", api_controller.LoginUser)
	// api.POST("/logout", api_controller.LogoutUser)
	// api.GET("/session/check", api_controller.CheckSession)

	// Certificate
	api.GET("/certificate/validate/id/:certificate_id", api_controller.GetCertificateByID)

	/*------------------------------------------ server start ------------------------------------------*/
	srv.newEcho.Logger.Fatal(srv.newEcho.Start(port))
}
