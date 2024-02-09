package config

import (
	"sertif_validator/app/utils"

	_ "github.com/joho/godotenv/autoload"
)

var (
	BaseURL    string = utils.GetEnv("BASE_URL", "")
	DbValidator string = utils.GetEnv("DB_SERTIF_VALIDATOR_URL", "")
)
