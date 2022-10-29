package handler

import (
	"context"
	"log"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/youtube-api-golang/internal/apis/operations/user"
	"github.com/youtube-api-golang/internal/models"
	"github.com/youtube-api-golang/internal/shared"
)

func (h *handler) RegisterUserHandler() user.PostUserRegisterHandlerFunc {
	return func(params user.PostUserRegisterParams) middleware.Responder {
		result, err := h.useCase.RegisterUser(context.Background(), params)

		if err != nil {
			return user.NewPostUserRegisterBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return user.NewPostUserRegisterOK().WithPayload(&user.PostUserRegisterOKBody{
			Message: result.Message,
			UserID:  result.UserID,
		}).WithAccessControlAllowOrigin(os.Getenv("CLIENT_URL"))
	}
}

func (h *handler) UpdateUserHandler() user.PutUserIDHandlerFunc {
	return func(params user.PutUserIDParams, tokenData interface{}) middleware.Responder {
		tokenClaim, err := shared.InterfaceToTokenClaim(tokenData)
		if err != nil {
			log.Println(err)
		}

		res, err := h.useCase.UpdateUser(context.Background(), params, tokenClaim.UserID)

		if err != nil {
			return user.NewPutUserIDBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return user.NewPutUserIDOK().WithPayload(res).WithAccessControlAllowOrigin(os.Getenv("CLIENT_URL"))
	}
}

func (h *handler) DeleteUserHandler() user.DeleteUserIDHandlerFunc {
	return func(params user.DeleteUserIDParams, tokenData interface{}) middleware.Responder {
		tokenClaim, err := shared.InterfaceToTokenClaim(tokenData)
		if err != nil {
			log.Println(err)
		}

		res, err := h.useCase.DeleteUser(context.Background(), params, tokenClaim.UserID)

		if err != nil {
			return user.NewPutUserIDBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return user.NewDeleteUserIDOK().WithPayload(res).WithAccessControlAllowOrigin(os.Getenv("CLIENT_URL"))
	}
}

func (h *handler) GetUserByIDHandler() user.GetUserIDHandlerFunc {
	return func(params user.GetUserIDParams) middleware.Responder {
		res, err := h.useCase.GetUserByID(context.Background(), params)
		if err != nil {
			return user.NewGetUserIDBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return user.NewGetUserIDOK().WithPayload(res).WithAccessControlAllowOrigin(os.Getenv("CLIENT_URL"))
	}
}
