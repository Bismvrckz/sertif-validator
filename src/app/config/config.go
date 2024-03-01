package config

import (
	"sertif_validator/app/utils"

	_ "github.com/joho/godotenv/autoload"
)

var (
	BaseURL              = utils.GetEnv("BASE_URL", "/sv")
	JwtKey               = utils.GetEnv("SV_JWT_KEY", "LmPZJbddZ9uXW4JE7g6N9R8ZdmDRv5vYihZJRBcOz7U=")
	IAMURL               = utils.GetEnv("KC_URL", "https://cicd.jatelindo.co.id/kc")
	IAMClientSecret      = utils.GetEnv("KC_SECRET", "hGt92tXINrDm8kGHZJkYMBy65a609NIx")
	IAMClientID          = utils.GetEnv("KC_ID", "sv")
	IAMRealm             = utils.GetEnv("KC_REALM", "dev_sv")
	IAMState             = utils.GetEnv("KC_STATE", "authExt")
	IAMRedirectURL       = utils.GetEnv("KC_REDIRECT", "")
	IAMConfigURL         = IAMURL + "/realms/" + IAMRealm
	IAMRedirect302Url    = utils.GetEnv("302REDIRECT", "")
	IAMLogoutRedirectUrl = utils.GetEnv("LOGOUT_REDIRECT", "")
	DbValidator          = utils.GetEnv("DB_SERTIF_VALIDATOR_URL", "")
)
