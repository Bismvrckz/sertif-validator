package config

import (
	"sertif_validator/app/utils"

	_ "github.com/joho/godotenv/autoload"
)

var (
	AppHost                 = utils.GetEnv("APP_HOST", "http://localhost:9070")
	ApiPrefix               = utils.GetEnv("API_PREFIX", "/api")
	BaseURL                 = utils.GetEnv("BASE_URL", "/tkbai")
	AdminLoginURL           = BaseURL + "/admin"
	JwtKey                  = utils.GetEnv("SV_JWT_KEY", "LmPZJbddZ9uXW4JE7g6N9R8ZdmDRv5vYihZJRBcOz7U=")
	IAMURL                  = utils.GetEnv("KC_URL", "http://localhost:8080")
	IAMDockerURL            = utils.GetEnv("KC_DOCKER_URL", "http://keycloak:8080")
	IAMClientSecret         = utils.GetEnv("KC_SECRET", "mGDmTC5F0RrteSbmlpndSOv1JbKFKiFG")
	IAMClientID             = utils.GetEnv("KC_ID", "tkbai")
	IAMRealm                = utils.GetEnv("KC_REALM", "tkbai_dev")
	IAMState                = utils.GetEnv("KC_STATE", "authExt")
	IAMConfigURL            = IAMURL + "/realms/" + IAMRealm
	IAMDockerConfigURL      = IAMDockerURL + "/realms/" + IAMRealm
	IAMLoginRedirectPath    = utils.GetEnv("KC_LOGIN_REDIRECT_PATH", "")
	IAMLoginRedirect302Path = utils.GetEnv("KC_LOGIN_302REDIRECT_PATH", "")
	IAMLogoutRedirectPath   = utils.GetEnv("KC_LOGOUT_REDIRECT_PATH", "")
	IAMLoginRedirectURL     = "http://localhost:9070" + BaseURL + IAMLoginRedirectPath
	IAMLoginRedirect302URL  = "http://localhost:9070" + BaseURL + IAMLoginRedirect302Path
	IAMLogoutRedirectURL    = "http://localhost:9070" + BaseURL + IAMLogoutRedirectPath
	DbValidator             = utils.GetEnv("DB_SERTIF_VALIDATOR_URL", "")
)
