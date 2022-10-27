package usecase

import (
	"context"

	"github.com/youtube-api-golang/internal/apis/operations/auth"
	"github.com/youtube-api-golang/internal/apis/operations/comment"
	"github.com/youtube-api-golang/internal/apis/operations/like"
	"github.com/youtube-api-golang/internal/apis/operations/subscription"
	"github.com/youtube-api-golang/internal/apis/operations/user"
	"github.com/youtube-api-golang/internal/apis/operations/video"
	"github.com/youtube-api-golang/internal/models"
	"github.com/youtube-api-golang/internal/repositories"
)

type useCase struct {
	repo repositories.Repositories
}

type UseCase interface {
	GetHealthcheck(ctx context.Context) (*models.Health, error)

	// auth
	Login(ctx context.Context, params auth.PostUserLoginParams) (*models.LoginResponse, error)

	// user
	RegisterUser(ctx context.Context, params user.PostUserRegisterParams) (*user.PostUserRegisterOKBody, error)
	UpdateUser(ctx context.Context, params user.PutUserIDParams, tokenUserID string) (*user.PutUserIDOKBody, error)
	DeleteUser(ctx context.Context, params user.DeleteUserIDParams, tokenUserID string) (*user.DeleteUserIDOKBody, error)
	GetUserByID(ctx context.Context, params user.GetUserIDParams) (*user.GetUserIDOKBody, error)

	// subcription
	Subscribe(ctx context.Context, params subscription.PatchSubUserIDParams, tokenUserID string) error
	Unsubscribe(ctx context.Context, params subscription.PatchUnsubUserIDParams, tokenUserID string) error

	// video
	AddVideo(ctx context.Context, params video.PostVideoParams, tokenUserID string) (*video.PostVideoOKBody, error)
	UpdateVideo(ctx context.Context, params video.PutVideoIDParams, tokenUserID string) (*video.PutVideoIDOKBody, error)
	DeleteVideo(ctx context.Context, params video.DeleteVideoIDParams, tokenUserID string) (*video.DeleteVideoIDOKBody, error)
	GetVideoByID(ctx context.Context, params video.GetVideoIDParams) (*video.GetVideoIDOKBody, error)
	UpdateView(ctx context.Context, params video.PatchVideoViewIDParams, tokenUserID string) (viewCount *int, err error)
	GetRandomVideos(ctx context.Context) (videoSlice []models.Video, err error)
	GetTrendingVideos(ctx context.Context) (videoSlice []models.Video, err error)
	GetVideosFromSubscribedChannels(ctx context.Context, userID string) (videos []models.Video, err error)
	GetVideosByTags(ctx context.Context, params video.GetVideoTagsParams) (videos []models.Video, err error)
	SearchVideos(ctx context.Context, params video.GetVideoSearchParams) (videos []models.Video, err error)

	// comment
	AddComment(ctx context.Context, params comment.PostCommentParams, tokenUserID string) (commentID string, err error)
	DeleteComment(ctx context.Context, params comment.DeleteCommentParams, tokenUserID string) (deletedCommentID string, err error)
	GetCommentsByVideoID(ctx context.Context, videoID string) (commentSlice []models.Comment, err error)

	// like
	UpdateVideoLike(ctx context.Context, params like.PatchLikeVideoIDParams, tokenUserID string) (likeCount *int, err error)
	UpdateVideoDislike(ctx context.Context, params like.PatchDislikeVideoIDParams, tokenUserID string) (likeCount *int, err error)
}

func NewUseCase(r repositories.Repositories) UseCase {
	return &useCase{repo: r}
}
