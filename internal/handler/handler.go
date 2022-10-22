package handler

import (
	"github.com/youtube-api-golang/configs"
	"github.com/youtube-api-golang/internal/apis/operations/auth"
	"github.com/youtube-api-golang/internal/apis/operations/health"
	"github.com/youtube-api-golang/internal/apis/operations/subscription"
	"github.com/youtube-api-golang/internal/apis/operations/user"
	"github.com/youtube-api-golang/internal/usecase"
)

type handler struct {
	useCase usecase.UseCase
}

type Handlers interface {
	// health check
	GetHealthHandler() health.GetHealthHandlerFunc

	// user
	RegisterUserHandler() user.PostUserRegisterHandlerFunc
	UpdateUserHandler() user.PutUserIDHandlerFunc
	DeleteUserHandler() user.DeleteUserIDHandlerFunc
	GetUserByIDHandler() user.GetUserIDHandlerFunc

	// auth
	LoginHandler() auth.PostUserLoginHandlerFunc

	// subscription
	SubscribeHandler() subscription.PatchSubUserIDHandlerFunc
	UnsubscribeHandler() subscription.PatchUnsubUserIDHandlerFunc
}

func NewHandler() Handlers {
	return &handler{
		useCase: configs.GetUsecase(),
	}
}
