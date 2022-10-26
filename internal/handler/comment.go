package handler

import (
	"context"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/youtube-api-golang/internal/apis/operations/comment"
	"github.com/youtube-api-golang/internal/models"
	"github.com/youtube-api-golang/internal/shared"
)

func (h *handler) AddCommentHandler() comment.PostCommentHandlerFunc {
	return func(params comment.PostCommentParams, tokenData interface{}) middleware.Responder {
		tokenClaim, err := shared.InterfaceToTokenClaim(tokenData)
		if err != nil {
			log.Println(err)
		}

		commentID, err := h.useCase.AddComment(context.Background(), params, tokenClaim.UserID)

		if err != nil {
			return comment.NewPostCommentBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return comment.NewPostCommentOK().WithPayload(
			&comment.PostCommentOKBody{
				Message:   "Success",
				CommentID: commentID,
			},
		)
	}
}

func (h *handler) DeleteCommentHandler() comment.DeleteCommentHandlerFunc {
	return func(params comment.DeleteCommentParams, tokenData interface{}) middleware.Responder {
		tokenClaim, err := shared.InterfaceToTokenClaim(tokenData)
		if err != nil {
			log.Println(err)
		}

		deletedCommentID, err := h.useCase.DeleteComment(context.Background(), params, tokenClaim.UserID)

		if err != nil {
			return comment.NewDeleteCommentBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		return comment.NewDeleteCommentOK().WithPayload(
			&comment.DeleteCommentOKBody{
				Message:   "Success",
				CommentID: deletedCommentID,
			},
		)
	}
}

func (h *handler) GetCommentsByVideoIDHandler() comment.GetCommentsVideoIDHandlerFunc {
	return func(params comment.GetCommentsVideoIDParams) middleware.Responder {
		comments, err := h.useCase.GetCommentsByVideoID(context.Background(), params.VideoID)

		if err != nil {
			return comment.NewGetCommentsVideoIDBadRequest().WithPayload(&models.Error{Code: "500", Message: err.Error()})
		}

		var arr []*models.Comment
		for i, _ := range comments {
			arr = append(arr, &comments[i])
		}
		return comment.NewGetCommentsVideoIDOK().WithPayload(
			&comment.GetCommentsVideoIDOKBody{
				Message:  "Success",
				Comments: arr,
			},
		)
	}
}
