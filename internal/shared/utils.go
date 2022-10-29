package shared

import (
	"encoding/json"
	"math/rand"
	"time"

	configs "github.com/youtube-api-golang/configs/auth"
	"github.com/youtube-api-golang/internal/models"
	"github.com/youtube-api-golang/internal/repositories/query"
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

func MapMongoVideoSliceResponse(videos []query.Video) (videoSlice []models.Video) {
	for _, v := range videos {

		var dislikes []string
		for _, v := range v.Dislikes {
			dislikes = append(dislikes, v.Hex())
		}

		var likes []string
		for _, v := range v.Likes {
			likes = append(likes, v.Hex())
		}

		userIDString := v.UserID.Hex()
		viewCount := int64(len(v.Views))

		createdAtStr := v.CreatedAt.Format(time.UnixDate)
		updatedAtStr := v.UpdatedAt.Format(time.UnixDate)

		video := models.Video{
			ID:          v.ID.Hex(),
			Description: v.Description,
			Dislikes:    dislikes,
			Likes:       likes,
			ImgURL:      v.ImgURL,
			Tags:        v.Tags,
			Title:       v.Title,
			UserID:      userIDString,
			VideoURL:    v.VideoURL,
			Views:       viewCount,
			CreatedAt:   createdAtStr,
			UpdatedAt:   updatedAtStr,
		}

		videoSlice = append(videoSlice, video)
	}

	return videoSlice
}
