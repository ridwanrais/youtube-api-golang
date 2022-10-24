package shared

import (
	"encoding/json"
	"math/rand"

	configs "github.com/youtube-api-golang/configs/auth"
)

func GenerateRandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func InterfaceToTokenClaim(i interface{}) (*configs.AccessTokenClaim, error) {
	claim := configs.AccessTokenClaim{}

	byteData, _ := json.Marshal(i)
	err := json.Unmarshal(byteData, &claim)
	if err != nil {
		return nil, err
	}

	return &claim, nil
}

// func MapVideoResponse()
