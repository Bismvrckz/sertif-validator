package webMiddleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sertif_validator/app/config"
	"sertif_validator/app/service/handler"
	"strings"
)

func AdminGetCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		//loggers.Debug().Msg("==========================> Start		| GetCookieMiddleware")
		htmlData := make(map[string]interface{})

		//loggers.Debug().Any("config.WebPrefix", config.WebPrefix).Msg("PATH")
		//loggers.Debug().Any("ctx.Path()", ctx.Path()).Msg("PATH")
		if !strings.Contains(ctx.Path(), config.BaseURL) {
			return ctx.Redirect(http.StatusSeeOther, config.BaseURL+"/")
		}

		accessToken, err := handler.ReadCookie(ctx, "accessToken")
		if err != nil {
			//loggers.Err(err).Msg("")
			return ctx.Redirect(http.StatusSeeOther, config.AdminLoginURL)
		}

		refreshToken, err := handler.ReadCookie(ctx, "refreshToken")
		if err != nil {
			//loggers.Err(err).Msg("")
			return ctx.Redirect(http.StatusSeeOther, config.AdminLoginURL)
		}

		idToken, err := handler.ReadCookie(ctx, "idToken")
		if err != nil {
			//loggers.Err(err).Msg("")
			return ctx.Redirect(http.StatusSeeOther, config.AdminLoginURL)
		}

		expiry, err := handler.ReadCookie(ctx, "expiry")
		if err != nil {
			//loggers.Err(err).Msg("")
			return ctx.Redirect(http.StatusSeeOther, config.AdminLoginURL)
		}

		//loggers.Debug().Any("accessToken", accessToken != nil).Msg("COOKIE")
		//loggers.Debug().Any("refreshToken", refreshToken != nil).Msg("COOKIE")
		//loggers.Debug().Any("idToken", idToken != nil).Msg("COOKIE")
		//loggers.Debug().Any("expiry", expiry != nil).Msg("COOKIE")

		htmlData["refreshToken"] = refreshToken.Value
		htmlData["accessToken"] = accessToken.Value
		htmlData["idToken"] = idToken.Value
		htmlData["expiry"] = expiry.Value

		htmlData["BaseURL"] = config.BaseURL
		//htmlData["apiPrefix"] = config.ApiPrefix
		//htmlData["apiHost"] = config.ApiHost

		ctx.Set("htmlData", htmlData)

		//loggers.Debug().Msg("==========================> Done		| GetCookieMiddleware")
		return next(ctx)
	}
}
