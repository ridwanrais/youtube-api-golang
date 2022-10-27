package usecase

import (
	"context"

	"github.com/youtube-api-golang/internal/apis/operations/like"
	"github.com/youtube-api-golang/internal/repositories/query"
)

func (u *useCase) UpdateVideoLike(ctx context.Context, params like.PatchLikeVideoIDParams, tokenUserID string) (likeCount *int, err error) {
	res, err := u.repo.UpdateVideoLike(ctx, query.UpdateLike{
		VideoID: params.VideoID,
		UserID:  tokenUserID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *useCase) UpdateVideoDislike(ctx context.Context, params like.PatchDislikeVideoIDParams, tokenUserID string) (likeCount *int, err error) {
	res, err := u.repo.UpdateVideoDislike(ctx, query.UpdateLike{
		VideoID: params.VideoID,
		UserID:  tokenUserID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
