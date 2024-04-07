package config

import "os"

var (
	SERVERPort              = GetEnv("SERVER_PORT", ":9070")
	AppHost                 = GetEnv("APP_HOST", "http://localhost:9070")
	ApiPrefix               = GetEnv("API_PREFIX", "/api")
	BaseURL                 = GetEnv("BASE_URL", "/tkbai")
	AdminLoginURL           = BaseURL + "/admin"
	JwtKey                  = GetEnv("SV_JWT_KEY", "LmPZJbddZ9uXW4JE7g6N9R8ZdmDRv5vYihZJRBcOz7U=")
	IAMURL                  = GetEnv("KC_URL", "http://localhost:8080")
	IAMDockerURL            = GetEnv("KC_DOCKER_URL", "http://keycloak:8080")
	IAMClientSecret         = GetEnv("KC_SECRET", "mGDmTC5F0RrteSbmlpndSOv1JbKFKiFG")
	IAMClientID             = GetEnv("KC_ID", "tkbai")
	IAMRealm                = GetEnv("KC_REALM", "tkbai_dev")
	IAMState                = GetEnv("KC_STATE", "authExt")
	IAMConfigURL            = IAMURL + "/realms/" + IAMRealm
	IAMDockerConfigURL      = IAMDockerURL + "/realms/" + IAMRealm
	IAMLoginRedirectPath    = GetEnv("KC_LOGIN_REDIRECT_PATH", "")
	IAMLoginRedirect302Path = GetEnv("KC_LOGIN_302REDIRECT_PATH", "")
	IAMLogoutRedirectPath   = GetEnv("KC_LOGOUT_REDIRECT_PATH", "")
	IAMLoginRedirectURL     = "http://localhost:9070" + BaseURL + IAMLoginRedirectPath
	IAMLoginRedirect302URL  = "http://localhost:9070" + BaseURL + IAMLoginRedirect302Path
	IAMLogoutRedirectURL    = "http://localhost:9070" + BaseURL + IAMLogoutRedirectPath
	TkbaiDB                 = GetEnv("DB_SERTIF_VALIDATOR_URL", "")
)

func GetEnv(key, fallback string) string {
	logger := Log
	if value, ok := os.LookupEnv(key); ok {
		logger.Debug().Str(key, value).Msg("Env")
		return value
	}
	logger.Debug().Str(key, fallback).Msg("Fallback")
	return fallback
}
