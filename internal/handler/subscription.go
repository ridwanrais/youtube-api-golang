package handler

import (
	"context"
	"log"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/youtube-api-golang/internal/apis/operations/subscription"
	"github.com/youtube-api-golang/internal/models"
	"github.com/youtube-api-golang/internal/shared"
)

func (h *handler) SubscribeHandler() subscription.PatchSubUserIDHandlerFunc {
	return func(params subscription.PatchSubUserIDParams, tokenData interface{}) middleware.Responder {
		tokenClaim, err := shared.InterfaceToTokenClaim(tokenData)
		if err != nil {
			log.Println(err)
		}

		err = h.useCase.Subscribe(context.Background(), params, tokenClaim.UserID)

		if err != nil {
			return subscription.NewPatchSubUserIDBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return subscription.NewPatchSubUserIDOK().WithPayload(
			&subscription.PatchSubUserIDOKBody{Message: "Success"},
		).WithAccessControlAllowOrigin(os.Getenv("CLIENT_URL"))
	}
}

func (h *handler) UnsubscribeHandler() subscription.PatchUnsubUserIDHandlerFunc {
	return func(params subscription.PatchUnsubUserIDParams, tokenData interface{}) middleware.Responder {
		tokenClaim, err := shared.InterfaceToTokenClaim(tokenData)
		if err != nil {
			log.Println(err)
		}

		err = h.useCase.Unsubscribe(context.Background(), params, tokenClaim.UserID)

		if err != nil {
			return subscription.NewPatchUnsubUserIDBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return subscription.NewPatchUnsubUserIDOK().WithPayload(
			&subscription.PatchUnsubUserIDOKBody{Message: "Success"},
		).WithAccessControlAllowOrigin(os.Getenv("CLIENT_URL"))
	}
}
