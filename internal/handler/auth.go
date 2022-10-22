package handler

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/youtube-api-golang/internal/apis/operations/auth"
	"github.com/youtube-api-golang/internal/models"
)

func (h *handler) LoginHandler() auth.PostUserLoginHandlerFunc {
	return func(params auth.PostUserLoginParams) middleware.Responder {
		res, err := h.useCase.Login(context.Background(), params)

		if err != nil {
			return auth.NewPostUserLoginBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return auth.NewPostUserLoginOK().WithPayload(res).WithSetCookie(fmt.Sprintf("access_token=%s; Path=/; Expires=%s; HttpOnly", res.Value, res.ExpiredAt))
	}
}
