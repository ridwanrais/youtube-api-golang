package handler

import (
	"context"

	"github.com/youtube-api-golang/configs"
	"github.com/youtube-api-golang/internal/models"
	"github.com/youtube-api-golang/internal/usecase"
)

type handler struct {
	useCase usecase.UseCase
}

type Handlers interface {
	GetHealtcheck(ctx context.Context) (*models.Health, error)
}

func NewHandler() Handlers {
	return &handler{
		useCase: configs.GetUsecase(),
	}
}
