package middleware

import (
	"errors"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	configs "github.com/youtube-api-golang/configs/auth"
)

func ValidateHeader(bearerCookie string) (interface{}, error) {
	bearerToken := strings.Split(bearerCookie, " ")[1]
	claims := configs.AccessTokenClaim{}

	token, err := jwt.ParseWithClaims(bearerToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_TOKEN_KEY")), nil
	})

	if err != nil {
		return nil, err
	}
	if token.Valid {
		return &claims, nil
	}
	return nil, errors.New("invalid token")
}
