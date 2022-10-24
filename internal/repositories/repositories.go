package repositories

import (
	"context"

	"github.com/youtube-api-golang/internal/models"
	"github.com/youtube-api-golang/internal/repositories/query"
	"go.mongodb.org/mongo-driver/mongo"
)

type repositories struct {
	db *mongo.Database
}

type Repositories interface {
	GetHealthCheck(ctx context.Context) (*models.Health, error)

	// user
	GetUserByUsername(ctx context.Context, username string) *query.User
	GetUserByEmail(ctx context.Context, email string) *query.User
	GetUserByID(ctx context.Context, userID string) (user *models.UserResponse, err error)
	CreateUser(ctx context.Context, user query.User) (userID string, err error)
	UpdateUser(ctx context.Context, user query.User, userID string) (updatedUser *models.UserResponse, err error)
	DeleteUser(ctx context.Context, userID string) (deletedUserID string, err error)

	// subscribtion
	Subscribe(ctx context.Context, params query.UpdateSub) error
	Unsubscribe(ctx context.Context, params query.UpdateSub) error

	// video
	AddVideo(ctx context.Context, video query.Video) (videoID string, err error)
	GetVideoByID(ctx context.Context, videoID string) (video *models.Video, err error)
	UpdateVideo(ctx context.Context, video query.Video, videoID string) (updatedVideo *models.Video, err error)
	DeleteVideo(ctx context.Context, videoID string) (deletedVideoID string, err error)
	UpdateView(ctx context.Context, params query.UpdateView) (viewCount *int, err error)
	GetRandomVideos(ctx context.Context, limit int64) (videoSlice []models.Video, err error)
	GetTrendingVideos(ctx context.Context, limit int64) (videoSlice []models.Video, err error)
	GetVideosFromSubscribedChannels(ctx context.Context, subscribedChannels []string) (videos []models.Video, err error)
}

func NewRepositories(d *mongo.Database) Repositories {
	return &repositories{db: d}
}
