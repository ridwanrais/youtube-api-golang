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
	GetUserByID(ctx context.Context, userID string) (user *query.User, err error)
	CreateUser(ctx context.Context, user query.User) (userID string, err error)
	UpdateUser(ctx context.Context, user query.User) (updatedUser *query.User, err error)
	DeleteUser(ctx context.Context, userID string) (deletedUserID string, err error)

	// subscribtion
	Subscribe(ctx context.Context, params query.UpdateSub) error
	Unsubscribe(ctx context.Context, params query.UpdateSub) error
}

func NewRepositories(d *mongo.Database) Repositories {
	return &repositories{db: d}
}
