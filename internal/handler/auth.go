package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/youtube-api-golang/internal/apis/operations/auth"
	"github.com/youtube-api-golang/internal/models"
)

func (h *handler) LoginHandler() auth.PostAuthLoginHandlerFunc {
	return func(params auth.PostAuthLoginParams) middleware.Responder {
		res, err := h.useCase.Login(context.Background(), params)

		if err != nil {
			return auth.NewPostAuthLoginBadRequest().WithPayload(&models.Error{Code: "400", Message: err.Error()})
		}

		return auth.NewPostAuthLoginOK().WithPayload(res).WithSetCookie(fmt.Sprintf("access_token=%s; Path=/; Expires=%s; HttpOnly", res.Value, res.ExpiredAt)).WithAccessControlAllowOrigin(os.Getenv("CLIENT_URL"))
	}
}
