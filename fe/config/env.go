package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var (
	SERVERPort              = GetEnv("FE_WEB_PORT", ":9071")
	ApiPrefix               = BaseURL + GetEnv("FE_API_PREFIX", "/api")
	APIHost                 = GetEnv("FE_API_HOST", "http://localhost:9070")
	WebHost                 = "http://localhost" + SERVERPort
	BaseURL                 = GetEnv("FE_BASE_URL", "/tkbai")
	WebTemplatesPath        = GetEnv("FE_WEB_TEMPLATES_PATH", "/root/tkbai-dashboard/fe/public/view/*.html")
	WebStaticUrl            = GetEnv("FE_WEB_STATIC_URL", "/static")
	WebStaticPath           = GetEnv("FE_WEB_STATIC_PATH", "/root/tkbai-dashboard/fe/public")
	AdminLoginURL           = BaseURL + "/login/admin"
	JwtKey                  = GetEnv("FE_SV_JWT_KEY", "LmPZJbddZ9uXW4JE7g6N9R8ZdmDRv5vYihZJRBcOz7U=")
	IAMURL                  = GetEnv("FE_KC_URL", "http://localhost:8080")
	IAMDockerURL            = GetEnv("FE_KC_DOCKER_URL", "http://keycloak:8080")
	IAMClientSecret         = GetEnv("FE_KC_SECRET", "mGDmTC5F0RrteSbmlpndSOv1JbKFKiFG")
	IAMClientID             = GetEnv("FE_KC_ID", "tkbai")
	IAMRealm                = GetEnv("FE_KC_REALM", "tkbai_dev")
	IAMState                = GetEnv("FE_KC_STATE", "authExt")
	IAMConfigURL            = IAMURL + "/realms/" + IAMRealm
	IAMDockerConfigURL      = IAMDockerURL + "/realms/" + IAMRealm
	IAMLoginRedirectPath    = GetEnv("FE_KC_LOGIN_REDIRECT_PATH", "")
	IAMLoginRedirect302Path = GetEnv("FE_KC_LOGIN_302REDIRECT_PATH", "")
	IAMLogoutRedirectPath   = GetEnv("FE_KC_LOGOUT_REDIRECT_PATH", "")
	IAMLoginRedirectURL     = "http://localhost:9070" + BaseURL + IAMLoginRedirectPath
	IAMLoginRedirect302URL  = "http://localhost:9070" + BaseURL + IAMLoginRedirect302Path
	IAMLogoutRedirectURL    = "http://localhost:9070" + BaseURL + IAMLogoutRedirectPath
	TkbaiDB                 = GetEnv("FE_TKBAI_DB_URL", "root:03IZmt7eRMukIHdoZahl@tcp(mysql:3306)/tkbai")
)

func GetEnv(key, fallback string) (value string) {
	logger := Log
	if value, ok := os.LookupEnv(key); ok {
		logger.Debug().Str(key, value).Msg("Env")
		return value
	}
	logger.Error().Str(key, fallback).Msg("Fallback")
	return fallback
}
