package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	configs "github.com/youtube-api-golang/configs/auth"
	"github.com/youtube-api-golang/internal/apis/operations/auth"
	"github.com/youtube-api-golang/internal/models"
	"github.com/youtube-api-golang/internal/repositories/query"
	"github.com/youtube-api-golang/internal/shared"
)

func (u *useCase) Login(ctx context.Context, params auth.PostAuthLoginParams) (*models.LoginResponse, error) {
	fmt.Println(params)

	user := u.repo.GetUserByUsername(ctx, *params.Body.Username)
	if user == nil {
		return nil, errors.New("akun tidak terdaftar")
	}

	isPasswordValid := shared.CheckPasswordHash(*params.Body.Password, user.Password)
	if !isPasswordValid {
		return nil, errors.New("password yang anda masukkan salah")
	}

	return u.createAccessToken(ctx, user)
}

func (u *useCase) createAccessToken(ctx context.Context, params *query.User) (*models.LoginResponse, error) {
	accessTokenLifetimeStr := os.Getenv("ACCESS_TOKEN_LIFETIME")
	accessTokenLifetime, err := strconv.Atoi(accessTokenLifetimeStr)
	if err != nil {
		log.Panic(`ENV ACCESS_TOKEN_LIFETIME ERROR `, err)
	}

	tokenConfig := configs.TokenConfig{
		Name:                os.Getenv("TOKEN_NAME"),
		Type:                os.Getenv("TOKEN_TYPE"),
		AccessTokenLifeTime: time.Duration(accessTokenLifetime) * time.Hour,
	}

	tokenIDBytes := make([]byte, 12)
	_, err = rand.Read(tokenIDBytes)
	if err != nil {
		panic(err)
	}

	accessTokenClaim := configs.AccessTokenClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(tokenConfig.AccessTokenLifeTime)).Unix(),
			Issuer:    tokenConfig.Name,
		},
		UserID:   params.ID.Hex(),
		Username: params.Username,
		TokenID:  base64.StdEncoding.EncodeToString(tokenIDBytes),
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaim)
	tokenString, err := newToken.SignedString([]byte(os.Getenv("ACCESS_TOKEN_KEY")))
	if err != nil {
		return nil, err
	}

	now := time.Now()
	end := now.Add(tokenConfig.AccessTokenLifeTime)

	res := &models.LoginResponse{
		Name:      tokenConfig.Name,
		Type:      tokenConfig.Type,
		Value:     tokenString,
		ExpiredAt: end.Format(time.RFC3339),
	}

	return res, nil
}
