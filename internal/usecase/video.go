package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/youtube-api-golang/internal/apis/operations/video"
	"github.com/youtube-api-golang/internal/models"
	"github.com/youtube-api-golang/internal/repositories/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *useCase) AddVideo(ctx context.Context, params video.PostVideoParams, tokenUserID string) (*video.PostVideoOKBody, error) {
	userID, _ := primitive.ObjectIDFromHex(tokenUserID)
	now := time.Now()
	newVideo := query.Video{
		UserID:      userID,
		Title:       *params.Body.Title,
		Description: *params.Body.Description,
		VideoURL:    *params.Body.VideoURL,
		ImgURL:      *params.Body.ImgURL,
		Tags:        params.Body.Tags,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	videoID, err := u.repo.AddVideo(ctx, newVideo)
	if err != nil {
		return nil, err
	}

	return &video.PostVideoOKBody{
		Message: "Success",
		VideoID: videoID,
	}, nil
}

func (u *useCase) UpdateVideo(ctx context.Context, params video.PutVideoIDParams, tokenUserID string) (*video.PutVideoIDOKBody, error) {
	res, err := u.repo.GetVideoByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	if res.UserID != tokenUserID {
		return nil, errors.New("Anda tidak berhak melakukan perintah ini")
	}

	var newVideo query.Video

	newVideo.Title = *params.Body.Title
	newVideo.Description = *params.Body.Description
	newVideo.ImgURL = *params.Body.ImgURL
	newVideo.Tags = params.Body.Tags
	newVideo.VideoURL = *params.Body.VideoURL

	now := time.Now()
	newVideo.UpdatedAt = now

	updatedVideo, err := u.repo.UpdateVideo(ctx, newVideo, params.ID)
	if err != nil {
		return nil, err
	}

	return &video.PutVideoIDOKBody{
		Message: "Success",
		Video:   updatedVideo,
	}, nil
}

func (u *useCase) DeleteVideo(ctx context.Context, params video.DeleteVideoIDParams, tokenUserID string) (*video.DeleteVideoIDOKBody, error) {
	res, err := u.repo.GetVideoByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	if res.UserID != tokenUserID {
		return nil, errors.New("Anda tidak berhak melakukan perintah ini")
	}

	deletedVideoID, err := u.repo.DeleteVideo(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return &video.DeleteVideoIDOKBody{
		Message: "Success",
		VideoID: deletedVideoID,
	}, nil
}

func (u *useCase) GetVideoByID(ctx context.Context, params video.GetVideoIDParams) (*video.GetVideoIDOKBody, error) {
	videoResponse, err := u.repo.GetVideoByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return &video.GetVideoIDOKBody{
		Message: "Success",
		Video:   videoResponse,
	}, nil
}

func (u *useCase) UpdateView(ctx context.Context, params video.PatchVideoViewIDParams, tokenUserID string) (viewCount *int, err error) {
	res, err := u.repo.UpdateView(ctx, query.UpdateView{
		VideoID: params.ID,
		UserID:  tokenUserID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *useCase) GetRandomVideos(ctx context.Context) (videoSlice []models.Video, err error) {
	res, err := u.repo.GetRandomVideos(ctx, 40)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *useCase) GetTrendingVideos(ctx context.Context) (videoSlice []models.Video, err error) {
	res, err := u.repo.GetTrendingVideos(ctx, 40)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *useCase) GetVideosFromSubscribedChannels(ctx context.Context, userID string) (videos []models.Video, err error) {
	user, err := u.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	subscribedChannels := user.SubscribedUsers

	res, err := u.repo.GetVideosFromSubscribedChannels(ctx, subscribedChannels)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *useCase) GetVideosByTags(ctx context.Context, params video.GetVideoTagsParams) (videos []models.Video, err error) {
	tagsArr := strings.Split(params.Tags, ",")

	videosArr, err := u.repo.GetVideosByTags(ctx, tagsArr, 40)
	if err != nil {
		return nil, err
	}

	return videosArr, nil
}

func (u *useCase) SearchVideos(ctx context.Context, params video.GetVideoSearchParams) (videos []models.Video, err error) {
	videosArr, err := u.repo.SearchVideos(ctx, params.Keyword, 40)
	if err != nil {
		return nil, err
	}

	return videosArr, nil
}
