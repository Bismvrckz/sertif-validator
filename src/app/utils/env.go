package utils

import (
	"excelvis/app/logging"
	"os"
)

func GetEnv(key, fallback string) string {
	logger := logging.Log
	if value, ok := os.LookupEnv(key); ok {
		logger.Debug().Str(key, value).Msg("Env")
		return value
	}
	logger.Debug().Str(key, fallback).Msg("fallback")
	return fallback
}
