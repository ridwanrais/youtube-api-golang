package usecase

import (
	"context"

	"github.com/youtube-api-golang/internal/models"
	"github.com/youtube-api-golang/internal/repositories"
)

type useCase struct {
	repo repositories.Repositories
}

type UseCase interface {
	GetHealthcheck(ctx context.Context) (*models.Health, error)
}

// func NewUseCase() UseCase {
// 	return &useCase{}
// }

func NewUseCase(r repositories.Repositories) UseCase {
	return &useCase{repo: r}
}
