package repositories

import (
	"context"

	"github.com/youtube-api-golang/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type repositories struct {
	client *mongo.Client
}

type Repositories interface {
	GetHealthCheck(ctx context.Context) (*models.Health, error)
}

func NewRepositories(c *mongo.Client) Repositories {
	return &repositories{client: c}
}
