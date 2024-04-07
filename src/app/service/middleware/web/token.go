package webMiddleware

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"sertif_validator/app/config"
	"sertif_validator/app/service/handler"
)

func AdminValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		//handler.loggers.Debug().Msg("==========================> Start		| AdminValidateToken")

		htmlData := ctx.Get("htmlData").(map[string]interface{})

		url := config.AppHost + config.ApiPrefix + "/entry/validate"

		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			//handler.loggers.Err(err).Msg("Error creating request")
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", htmlData["accessToken"].(string))
		req.AddCookie(&http.Cookie{Name: "refreshToken", Value: htmlData["refreshToken"].(string)})
		req.AddCookie(&http.Cookie{Name: "expiry", Value: htmlData["expiry"].(string)})

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			//handler.loggers.Err(err).Msg("Error sending request")
			return err
		}

		//handler.loggers.Debug().Str("Response Status", resp.Status).Msg("")
		if resp.StatusCode == 401 || resp.StatusCode == 403 {
			handler.DeleteCookie(ctx, "accessToken")
			handler.DeleteCookie(ctx, "refreshToken")
			handler.DeleteCookie(ctx, "idToken")
			handler.DeleteCookie(ctx, "expiry")
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

		//handler.loggers.Debug().Any("AdditionalInfo.Message", result.AdditionalInfo.Message).Msg("")

		err = resp.Body.Close()
		if err != nil {
			//handler.loggers.Err(err).Msg("Error closing body")
			return err
		}

		handler.WriteCookie(ctx, "accessToken", result.AdditionalInfo.AccessToken, config.BaseURL, 24)
		handler.WriteCookie(ctx, "refreshToken", result.AdditionalInfo.RefreshToken, config.BaseURL, 24)
		handler.WriteCookie(ctx, "expiry", result.AdditionalInfo.Expiry, config.BaseURL, 24)

		htmlData["refreshToken"] = result.AdditionalInfo.RefreshToken
		htmlData["accessToken"] = result.AdditionalInfo.AccessToken
		htmlData["expiry"] = result.AdditionalInfo.Expiry

		ctx.Set("htmlData", htmlData)

		//handler.loggers.Debug().Msg("==========================> Done		| ValidateTokenMiddleware")
		return next(ctx)
	}
}
