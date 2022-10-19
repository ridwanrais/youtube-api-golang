package handler

import (
	"context"

	"github.com/youtube-api-golang/internal/models"
)

func (h *handler) GetHealtcheck(ctx context.Context) (*models.Health, error) {
	return &models.Health{Status: "great"}, nil
}
