package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"tkbai-fe/config"
	"tkbai-fe/models"
)

func GetCertificateByID(ctx echo.Context, certificateID, certificateHolder string) (result struct {
	ResponseCode    string
	AdditionalInfo  models.ToeflCertificate
	ResponseMessage string
}, err error) {
	url := config.APIHost + config.ApiPrefix + "/admin/data/toefl/id/" + certificateID + "/name/" + certificateHolder

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		loggers.Err(err).Msg("Error creating request")
		return result, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		loggers.Err(err).Msg("Error sending request")
		return result, err
	}

	loggers.Debug().Str("Response Status", resp.Status).Msg("")
	if resp.StatusCode == 401 || resp.StatusCode == 403 {
		DeleteCookie(ctx, "accessToken")
		DeleteCookie(ctx, "refreshToken")
		DeleteCookie(ctx, "idToken")
		DeleteCookie(ctx, "expiry")
		err := ctx.Redirect(http.StatusSeeOther, config.AdminLoginURL)
		return result, err
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}

	err = resp.Body.Close()
	if err != nil {
		loggers.Err(err).Msg("Error closing body")
		return result, err
	}

	return result, err
}
