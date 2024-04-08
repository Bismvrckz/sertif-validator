package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"tkbai-fe/config"
)

type TokenStruct struct {
	AccessToken  string
	Expiry       string
	Message      string
	RefreshToken string
	IdToken      string
}

func AdminGetCookieMid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		loggers.Debug().Msg("==========================> Start		| GetCookieMiddleware")
		htmlData := make(map[string]interface{})

		loggers.Debug().Any("config.WebPrefix", config.BaseURL).Msg("PATH")
		loggers.Debug().Any("ctx.Path()", ctx.Path()).Msg("PATH")
		if !strings.Contains(ctx.Path(), config.BaseURL) {
			return ctx.Redirect(http.StatusSeeOther, config.BaseURL+"/")
		}

		accessToken, err := ReadCookie(ctx, "accessToken")
		if err != nil {
			config.LogErr(err, "")
			return ctx.Redirect(http.StatusSeeOther, config.AdminLoginURL)
		}

		refreshToken, err := ReadCookie(ctx, "refreshToken")
		if err != nil {
			config.LogErr(err, "")
			return ctx.Redirect(http.StatusSeeOther, config.AdminLoginURL)
		}

		idToken, err := ReadCookie(ctx, "idToken")
		if err != nil {
			config.LogErr(err, "")
			return ctx.Redirect(http.StatusSeeOther, config.AdminLoginURL)
		}

		expiry, err := ReadCookie(ctx, "expiry")
		if err != nil {
			config.LogErr(err, "")
			return ctx.Redirect(http.StatusSeeOther, config.AdminLoginURL)
		}

		loggers.Debug().Any("accessToken", accessToken != nil).Msg("COOKIE")
		loggers.Debug().Any("refreshToken", refreshToken != nil).Msg("COOKIE")
		loggers.Debug().Any("idToken", idToken != nil).Msg("COOKIE")
		loggers.Debug().Any("expiry", expiry != nil).Msg("COOKIE")

		htmlData["refreshToken"] = refreshToken.Value
		htmlData["accessToken"] = accessToken.Value
		htmlData["idToken"] = idToken.Value
		htmlData["expiry"] = expiry.Value

		htmlData["baseURL"] = config.BaseURL
		htmlData["apiPrefix"] = config.ApiPrefix
		htmlData["apiHost"] = config.APIHost
		htmlData["prefix"] = config.BaseURL

		ctx.Set("htmlData", htmlData)

		loggers.Debug().Msg("==========================> Done		| GetCookieMiddleware")
		return next(ctx)
	}
}

func AdminValidateTokenMid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		loggers.Debug().Msg("==========================> Start		| ValidateTokenMiddleware")

		htmlData := ctx.Get("htmlData").(map[string]interface{})

		url := config.APIHost + config.ApiPrefix + "/entry/validate"

		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			loggers.Err(err).Msg("Error creating request")
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", htmlData["accessToken"].(string))
		req.AddCookie(&http.Cookie{Name: "refreshToken", Value: htmlData["refreshToken"].(string)})
		req.AddCookie(&http.Cookie{Name: "expiry", Value: htmlData["expiry"].(string)})

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			loggers.Err(err).Msg("Error sending request")
			return err
		}

		loggers.Debug().Str("Response Status", resp.Status).Msg("")
		if resp.StatusCode == 401 || resp.StatusCode == 403 {
			DeleteCookie(ctx, "accessToken")
			DeleteCookie(ctx, "refreshToken")
			DeleteCookie(ctx, "idToken")
			DeleteCookie(ctx, "expiry")
			return ctx.Redirect(http.StatusSeeOther, config.AdminLoginURL)
		}

		var result struct {
			ResponseCode    string
			ResponseMessage string
			AdditionalInfo  TokenStruct
		}
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return err
		}

		loggers.Debug().Any("AdditionalInfo.Message", result.AdditionalInfo.Message).Msg("")

		err = resp.Body.Close()
		if err != nil {
			loggers.Err(err).Msg("Error closing body")
			return err
		}

		WriteCookie(ctx, "accessToken", result.AdditionalInfo.AccessToken, config.BaseURL, 24)
		WriteCookie(ctx, "refreshToken", result.AdditionalInfo.RefreshToken, config.BaseURL, 24)
		WriteCookie(ctx, "expiry", result.AdditionalInfo.Expiry, config.BaseURL, 24)

		htmlData["refreshToken"] = result.AdditionalInfo.RefreshToken
		htmlData["accessToken"] = result.AdditionalInfo.AccessToken
		htmlData["expiry"] = result.AdditionalInfo.Expiry

		ctx.Set("htmlData", htmlData)

		loggers.Debug().Msg("==========================> Done		| ValidateTokenMiddleware")
		return next(ctx)
	}
}
