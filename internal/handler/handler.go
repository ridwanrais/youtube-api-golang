package handler

import (
	"github.com/youtube-api-golang/configs"
	"github.com/youtube-api-golang/internal/apis/operations/auth"
	"github.com/youtube-api-golang/internal/apis/operations/comment"
	"github.com/youtube-api-golang/internal/apis/operations/health"
	"github.com/youtube-api-golang/internal/apis/operations/like"
	"github.com/youtube-api-golang/internal/apis/operations/subscription"
	"github.com/youtube-api-golang/internal/apis/operations/user"
	"github.com/youtube-api-golang/internal/apis/operations/video"
	"github.com/youtube-api-golang/internal/usecase"
)

type handler struct {
	useCase usecase.UseCase
}

type Handlers interface {
	// health check
	GetHealthHandler() health.GetHealthHandlerFunc

	// user
	RegisterUserHandler() user.PostUserRegisterHandlerFunc
	UpdateUserHandler() user.PutUserIDHandlerFunc
	DeleteUserHandler() user.DeleteUserIDHandlerFunc
	GetUserByIDHandler() user.GetUserIDHandlerFunc

	// auth
	LoginHandler() auth.PostAuthLoginHandlerFunc

	// subscription
	SubscribeHandler() subscription.PatchSubUserIDHandlerFunc
	UnsubscribeHandler() subscription.PatchUnsubUserIDHandlerFunc

	// video
	AddVideoHandler() video.PostVideoHandlerFunc
	UpdateVideoHandler() video.PutVideoIDHandlerFunc
	DeleteVideoHandler() video.DeleteVideoIDHandlerFunc
	GetVideoByIDHandler() video.GetVideoIDHandlerFunc
	UpdateViewHandler() video.PatchVideoViewIDHandlerFunc
	GetRandomVideosHandler() video.GetVideoRandomHandlerFunc
	GetTrendingVideosHandler() video.GetVideoTrendHandlerFunc
	GetVideosFromSubscribedChannelsHandler() video.GetVideoSubHandlerFunc
	GetVideosByTagsHandler() video.GetVideoTagsHandlerFunc
	SearchVideosHandler() video.GetVideoSearchHandlerFunc

	// comment
	AddCommentHandler() comment.PostCommentHandlerFunc
	DeleteCommentHandler() comment.DeleteCommentHandlerFunc
	GetCommentsByVideoIDHandler() comment.GetCommentsVideoIDHandlerFunc

	// like
	UpdateVideoLikeHandler() like.PatchLikeVideoIDHandlerFunc
	UpdateVideoDislikeHandler() like.PatchDislikeVideoIDHandlerFunc
}

func NewHandler() Handlers {
	return &handler{
		useCase: configs.GetUsecase(),
	}
}
