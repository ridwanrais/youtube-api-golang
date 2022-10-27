package handler

import (
	"context"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/youtube-api-golang/internal/apis/operations/like"
	"github.com/youtube-api-golang/internal/models"
	"github.com/youtube-api-golang/internal/shared"
)

func (h *handler) UpdateVideoLikeHandler() like.PatchLikeVideoIDHandlerFunc {
	return func(params like.PatchLikeVideoIDParams, tokenData interface{}) middleware.Responder {
		tokenClaim, err := shared.InterfaceToTokenClaim(tokenData)
		if err != nil {
			log.Println(err)
		}

		result, err := h.useCase.UpdateVideoLike(context.Background(), params, tokenClaim.UserID)

		if err != nil {
			return like.NewPatchLikeVideoIDBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return like.NewPatchLikeVideoIDOK().WithPayload(&like.PatchLikeVideoIDOKBody{
			Message:   "Success",
			LikeCount: int64(*result),
		})
	}
}

func (h *handler) UpdateVideoDislikeHandler() like.PatchDislikeVideoIDHandlerFunc {
	return func(params like.PatchDislikeVideoIDParams, tokenData interface{}) middleware.Responder {
		tokenClaim, err := shared.InterfaceToTokenClaim(tokenData)
		if err != nil {
			log.Println(err)
		}

		result, err := h.useCase.UpdateVideoDislike(context.Background(), params, tokenClaim.UserID)

		if err != nil {
			return like.NewPatchDislikeVideoIDBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return like.NewPatchDislikeVideoIDOK().WithPayload(&like.PatchDislikeVideoIDOKBody{
			Message:      "Success",
			DislikeCount: int64(*result),
		})
	}
}
