package usecase

import (
	"context"

	"github.com/youtube-api-golang/internal/models"
)

func (u *useCase) GetHealthcheck(ctx context.Context) (*models.Health, error) {
	return nil, nil
}
