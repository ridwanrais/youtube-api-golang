package usecase

import (
	"context"
	"errors"
	"net/mail"
	"time"

	"github.com/youtube-api-golang/internal/apis/operations/user"
	"github.com/youtube-api-golang/internal/shared"

	"github.com/youtube-api-golang/internal/repositories/query"
)

func (u *useCase) RegisterUser(ctx context.Context, params user.PostUserRegisterParams) (*user.PostUserRegisterOKBody, error) {
	_, err := mail.ParseAddress(*params.Body.Email)
	if err != nil {
		return nil, shared.CreateError(400, "email tidak valid")
	}

	// validate if username is already taken
	res := u.repo.GetUserByUsername(ctx, *params.Body.Username)
	if res != nil {
		return nil, shared.CreateError(400, "username tidak tersedia")
	}

	res = u.repo.GetUserByEmail(ctx, *params.Body.Email)
	if res != nil {
		return nil, shared.CreateError(400, "email sudah dipakai")
	}

	hashedPassword, _ := shared.HashPassword(*params.Body.Password)

	now := time.Now()

	newUser := query.User{
		Username:  *params.Body.Username,
		Email:     *params.Body.Email,
		Password:  hashedPassword,
		Img:       params.Body.Img,
		CreatedAt: now,
		UpdatedAt: now,
	}

	userID, err := u.repo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, shared.CreateError(500, "error saat create user")
	}

	return &user.PostUserRegisterOKBody{
		Message: "success",
		UserID:  userID,
	}, nil
}

func (u *useCase) UpdateUser(ctx context.Context, params user.PutUserIDParams, tokenUserID string) (*user.PutUserIDOKBody, error) {
	if params.ID != tokenUserID {
		return nil, errors.New("Anda tidak berhak melakukan perintah ini")
	}

	var newUser query.User

	if params.Body.Username != "" {
		res := u.repo.GetUserByUsername(ctx, params.Body.Username)
		if res != nil {
			return nil, shared.CreateError(400, "username tidak tersedia")
		}

		newUser.Username = params.Body.Username
	}

	if params.Body.Email != "" {
		_, err := mail.ParseAddress(params.Body.Email)
		if err != nil {
			return nil, shared.CreateError(400, "email tidak valid")
		}

		res := u.repo.GetUserByEmail(ctx, params.Body.Email)
		if res != nil {
			return nil, shared.CreateError(400, "email sudah dipakai")
		}

		newUser.Email = params.Body.Email
	}

	if params.Body.Password != "" {
		hashedPassword, _ := shared.HashPassword(params.Body.Password)

		newUser.Password = hashedPassword
	}

	if params.Body.Img != "" {
		newUser.Img = params.Body.Img
	}

	now := time.Now()
	newUser.UpdatedAt = now

	updatedUser, err := u.repo.UpdateUser(ctx, newUser, params.ID)
	if err != nil {
		return nil, err
	}

	return &user.PutUserIDOKBody{
		Message:     "success",
		UpdatedUser: updatedUser,
	}, nil
}

func (u *useCase) DeleteUser(ctx context.Context, params user.DeleteUserIDParams, tokenUserID string) (*user.DeleteUserIDOKBody, error) {
	if params.ID != tokenUserID {
		return nil, errors.New("Anda tidak berhak melakukan perintah ini")
	}

	deletedUserID, err := u.repo.DeleteUser(ctx, params.ID)
	if err != nil {
		return nil, shared.CreateError(500, "error saat delete user")
	}

	return &user.DeleteUserIDOKBody{
		Message: "Success",
		UserID:  deletedUserID,
	}, nil
}

func (u *useCase) GetUserByID(ctx context.Context, params user.GetUserIDParams) (*user.GetUserIDOKBody, error) {
	userResponse, err := u.repo.GetUserByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return &user.GetUserIDOKBody{
		Message: "Success",
		User:    userResponse,
	}, nil

}
