package usecase

import (
	"context"
	"errors"

	"github.com/youtube-api-golang/internal/apis/operations/subscription"
	"github.com/youtube-api-golang/internal/repositories/query"
)

func (u *useCase) Subscribe(ctx context.Context, params subscription.PatchSubUserIDParams, tokenUserID string) error {

	if tokenUserID == params.UserID {
		return errors.New("You can't subscribe to your own channel")
	}

	err := u.repo.Subscribe(ctx, query.UpdateSub{
		SubscribingID: tokenUserID,
		SubscribedID:  params.UserID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *useCase) Unsubscribe(ctx context.Context, params subscription.PatchUnsubUserIDParams, tokenUserID string) error {

	err := u.repo.Unsubscribe(ctx, query.UpdateSub{
		SubscribingID: tokenUserID,
		SubscribedID:  params.UserID,
	})
	if err != nil {
		return err
	}

	return nil
}
