package usecase

import (
	"context"

	"github.com/youtube-api-golang/internal/models"
)

func (u *useCase) GetHealthcheck(ctx context.Context) (*models.Health, error) {
	_, err := u.repo.GetHealthCheck(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
