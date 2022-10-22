package configs

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenConfig struct {
	Name                string
	Type                string
	AccessTokenLifeTime time.Duration
}

type AccessTokenClaim struct {
	jwt.StandardClaims
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	UserRole     string `json:"user_role"`
	TokenID      string `json:"token_id"`
	QrStringCode string `json:"qr_string_code"`
}
