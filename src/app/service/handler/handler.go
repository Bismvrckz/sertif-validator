package handler

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"sertif_validator/app/config"
	"sertif_validator/app/logging"
)

var (
	loggers = logging.Log
)

func GenerateJwtString(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	secret := []byte(config.JwtKey)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		loggers.Error().Stack().Err(err).Msg("Error signing token")
		return ""
	}

	return signedToken
}

func ParseJwtString(tokenString, claim string) (interface{}, error) {

	// Define the secret key used for signing the token
	secret := []byte(config.JwtKey)

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		loggers.Error().Stack().Err(err).Msg("Error parsing token")
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		loggers.Error().Stack().Err(err).Msg("Token is not valid")
		return nil, err
	}

	// Access claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Access the claims
		tokenExp := claims[claim].(string)

		return tokenExp, nil
	} else {
		loggers.Error().Stack().Err(err).Msg("Error reading claims")
		return nil, nil
	}
}
