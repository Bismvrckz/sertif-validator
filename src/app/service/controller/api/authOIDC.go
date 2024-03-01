package api_controller

import (
	"net/http"
	"sertif_validator/app/config"
	"sertif_validator/app/logging"
	"sertif_validator/app/service/handler"
	"strings"
	"time"

	"encoding/json"

	oidc "github.com/coreos/go-oidc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	_ "github.com/rs/zerolog"
	"golang.org/x/oauth2"
)

type OidcData struct {
	AuthCode  string
	AuthToken *oauth2.Token
	UserData  jwt.Claims
}

var (
	oauth2Config = oauth2.Config{
		ClientID:     config.IAMClientID,
		ClientSecret: config.IAMClientSecret,
		RedirectURL:  config.IAMRedirectURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
)

// TODO refactor me

func LoginOIDC(ctx echo.Context) (err error) {
	loggers := logging.Log

	ctxCore := ctx.Request().Context()
	provider, err := oidc.NewProvider(ctxCore, config.IAMConfigURL)
	if err != nil {
		loggers.Error().Stack().Err(err).Msg(err.Error())
		return err
	}
	oauth2Config.Endpoint = provider.Endpoint()

	oidcConfig := &oidc.Config{
		ClientID: config.IAMClientID,
	}
	verifier := provider.Verifier(oidcConfig)

	rawAccessToken := ctx.Request().Header.Get("Authorization")

	if rawAccessToken == "" {
		return ctx.Redirect(http.StatusFound, oauth2Config.AuthCodeURL(config.IAMState))
	}
	parts := strings.Split(rawAccessToken, " ")

	if len(parts) != 2 {
		return ctx.NoContent(http.StatusBadRequest)
	}

	var _, errVerifier = verifier.Verify(ctxCore, parts[1])
	if errVerifier != nil {

		return ctx.Redirect(http.StatusFound, oauth2Config.AuthCodeURL(config.IAMState))
	}

	//TODO il8n
	var respn interface{} = &Response{
		ResponseCode:    "00",
		AdditionalInfo:  "",
		ResponseMessage: "sukses",
	}

	return ctx.JSON(http.StatusOK, respn)
}

func LoginCallbackOIDC(ctx echo.Context) (err error) {
	loggers := logging.Log

	if ctx.Request().URL.Query().Get("state") != config.IAMState {
		return ctx.HTML(http.StatusBadRequest, "state did not match")
	}

	ctxCore := ctx.Request().Context()
	provider, err := oidc.NewProvider(ctxCore, config.IAMConfigURL)
	if err != nil {
		loggers.Error().Stack().Err(err).Msg("")
	}

	oauth2Config.Endpoint = provider.Endpoint()

	oidcConfig := &oidc.Config{
		ClientID: config.IAMClientID,
	}
	verifier := provider.Verifier(oidcConfig)

	oauth2Token, err := oauth2Config.Exchange(ctxCore, ctx.Request().URL.Query().Get("code"))
	if err != nil {
		return ctx.HTML(http.StatusInternalServerError, "Failed to exchange token: "+err.Error())
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return ctx.HTML(http.StatusInternalServerError, "No id_token field in oauth2 token.")
	}

	idToken, err := verifier.Verify(ctxCore, rawIDToken)
	if err != nil {
		return ctx.HTML(http.StatusInternalServerError, "Failed to verify ID Token: "+err.Error())
	}

	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	}{
		oauth2Token,
		new(json.RawMessage),
	}
	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		return ctx.HTML(http.StatusInternalServerError, err.Error())
	}

	tokenExp := handler.GenerateJwtString(jwt.MapClaims{
		"tokenExp": oauth2Token.Expiry.Format("2006-01-02T15:04:05.999999-07:00"),
	})

	return ctx.HTML(http.StatusOK,
		`<!DOCTYPE html>
<html>
<head>
<meta http-equiv="refresh" content="0; url='`+config.IAMRedirect302Url+`'" />
</head>
<body>
<script type="text/javascript">
function setCookie(name, value, days, path) {
var expires = "";
if (days) {
var date = new Date();
date.setTime(date.getTime() + days * 24 * 60 * 60 * 1000);
expires = "; expires=" + date.toUTCString();
}

document.cookie = name + "=" + value + expires + "; path=" + path;
}

setCookie("accessToken", "`+oauth2Token.AccessToken+`", 1, "`+config.BaseURL+`");
setCookie("refreshToken", "`+oauth2Token.RefreshToken+`", 1, "`+config.BaseURL+`");
setCookie("expiry", "`+tokenExp+`", 1, "`+config.BaseURL+`");
setCookie("idToken", '`+rawIDToken+`', 1, "`+config.BaseURL+`");
</script>
</body>
</html>	
`)
}

func LogoutOIDC(ctx echo.Context) (err error) {
	idToken := ctx.QueryParam("idToken")

	logoutUrl := config.IAMConfigURL + "/protocol/openid-connect/logout"
	logoutRedirectUrl := config.IAMLogoutRedirectUrl

	return ctx.Redirect(http.StatusSeeOther, logoutUrl+"?post_logout_redirect_uri="+logoutRedirectUrl+"&client_id=mav2&id_token_hint="+idToken)
}

func LogoutCallbackOIDC(ctx echo.Context) (err error) {
	// middlewares.DeleteCookie(ctx, "accessToken")
	// middlewares.DeleteCookie(ctx, "refreshToken")
	// middlewares.DeleteCookie(ctx, "idToken")
	// middlewares.DeleteCookie(ctx, "expiry")

	return ctx.HTML(http.StatusOK,
		`<!DOCTYPE html>
<html>
<head>
<meta http-equiv="refresh" content="0; url='`+config.IAMRedirect302Url+`'" />
</head>
<body></body>
</html>	`)
}

func ValidateOIDC(ctx echo.Context) (err error) {
	loggers := logging.Log

	var authBody AuthBody
	if err := ctx.Bind(&authBody); err != nil {
		return err
	}

	if authBody.AccessToken == "" || authBody.RefreshToken == "" || authBody.Expiry == "" {
		return ctx.JSON(http.StatusBadRequest, &Response{
			ResponseCode: "02",
			AdditionalInfo: map[string]interface{}{
				"authorizationBearer": authBody.AccessToken,
				"refreshToken":        authBody.RefreshToken,
				"expiry":              authBody.Expiry,
			},
			ResponseMessage: "gagal",
		})
	}

	tokenExp, err := handler.ParseJwtString(authBody.Expiry, "tokenExp")
	if err != nil {
		if strings.Contains(err.Error(), "token signature is invalid: signature is invalid") {
			return ctx.JSON(http.StatusUnauthorized, "Invalid Token")
		}
		return err
	}

	parsedTime, err := time.Parse("2006-01-02T15:04:05.999999-07:00", tokenExp.(string))
	if err != nil {
		return err
	}

	token := &oauth2.Token{
		AccessToken:  authBody.AccessToken,
		RefreshToken: authBody.RefreshToken,
		Expiry:       parsedTime,
		TokenType:    "Bearer",
	}

	if !token.Valid() {
		ctxCore := ctx.Request().Context()

		loggers.Info().Stack().Msg("accessToken not valid, refreshing")

		provider, err := oidc.NewProvider(ctxCore, config.IAMConfigURL)
		if err != nil {
			loggers.Error().Stack().Err(err).Msg(err.Error())
		}

		oauth2Config.Endpoint = provider.Endpoint()
		tokenSource := oauth2Config.TokenSource(ctxCore, token)
		newToken, err := tokenSource.Token()

		if err != nil {
			loggers.Error().Stack().Err(err).Msg("")

			if strings.Contains(err.Error(), "Token is not active") {
				return ctx.NoContent(http.StatusForbidden)
			} else if strings.Contains(err.Error(), "Session not active") {
				return ctx.NoContent(http.StatusUnauthorized)
			}

			return ctx.JSON(http.StatusBadRequest, "Failed to refresh token | "+err.Error())
		}

		expiry := handler.GenerateJwtString(jwt.MapClaims{
			"tokenExp": newToken.Expiry.Format("2006-01-02T15:04:05.999999-07:00"),
		})

		var respon interface{} = &Response{
			ResponseCode: "00",
			AdditionalInfo: map[string]interface{}{
				"accessToken":  newToken.AccessToken,
				"refreshToken": newToken.RefreshToken,
				"message":      "Token refreshed using refresh token.",
				"expiry":       expiry,
			},
			ResponseMessage: "sukses",
		}
		return ctx.JSON(http.StatusOK, respon)
	} else {
		loggers.Info().Stack().Msg("accessToken still valid")

		var respon interface{} = &Response{
			ResponseCode: "00",
			AdditionalInfo: map[string]interface{}{
				"accessToken":  authBody.AccessToken,
				"refreshToken": authBody.RefreshToken,
				"message":      "Token is still valid.",
				"expiry":       authBody.Expiry,
			},
			ResponseMessage: "sukses",
		}
		return ctx.JSON(http.StatusOK, respon)
	}
}
