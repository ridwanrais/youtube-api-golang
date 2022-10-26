package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/youtube-api-golang/internal/apis/operations/comment"
	"github.com/youtube-api-golang/internal/models"
	"github.com/youtube-api-golang/internal/repositories/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *useCase) AddComment(ctx context.Context, params comment.PostCommentParams, tokenUserID string) (commentID string, err error) {
	_, err = u.repo.GetVideoByID(ctx, *params.Body.VideoID)
	if err != nil {
		return "", err
	}

	userID, _ := primitive.ObjectIDFromHex(tokenUserID)
	videoID, _ := primitive.ObjectIDFromHex(*params.Body.VideoID)
	now := time.Now()

	newComment := query.Comment{
		UserID:      userID,
		VideoID:     videoID,
		Description: *params.Body.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	commentID, err = u.repo.AddComment(ctx, newComment)
	if err != nil {
		return "", err
	}

	return commentID, nil
}

func (u *useCase) DeleteComment(ctx context.Context, params comment.DeleteCommentParams, tokenUserID string) (deletedCommentID string, err error) {
	_, err = u.repo.GetVideoByID(ctx, *params.Body.VideoID)
	if err != nil {
		return "", errors.New("Video tidak ditemukan")
	}

	commentRes, err := u.repo.GetCommentByID(ctx, *params.Body.CommentID)
	if err != nil {
		return "", errors.New("Komentar tidak ditemukan")
	}

	if *commentRes.UserID != tokenUserID {
		return "", errors.New("Anda tidak berhak melakukan perintah ini")
	}

	deletedCommentID, err = u.repo.DeleteComment(ctx, *params.Body.CommentID)
	if err != nil {
		return "", err
	}

	return deletedCommentID, nil
}

func (u *useCase) GetCommentsByVideoID(ctx context.Context, videoID string) (commentSlice []models.Comment, err error) {
	_, err = u.repo.GetVideoByID(ctx, videoID)
	if err != nil {
		return nil, err
	}

	comments, err := u.repo.GetCommentsByVideoID(ctx, videoID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
