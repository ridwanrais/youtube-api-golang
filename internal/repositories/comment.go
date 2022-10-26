package repositories

import (
	"context"
	"fmt"

	"github.com/youtube-api-golang/internal/models"
	"github.com/youtube-api-golang/internal/repositories/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repositories) AddComment(ctx context.Context, comment query.Comment) (commentID string, err error) {
	coll := r.db.Collection("comments")

	result, err := coll.InsertOne(ctx, comment)
	if err != nil {
		return "", err
	}

	commentID = result.InsertedID.(primitive.ObjectID).Hex()

	return commentID, nil
}

func (r *repositories) DeleteComment(ctx context.Context, commentID string) (deletedCommentID string, err error) {
	coll := r.db.Collection("comments")
	id, _ := primitive.ObjectIDFromHex(commentID)

	var deletedComment *query.Comment
	err = coll.FindOneAndDelete(ctx, bson.D{{"_id", id}}).Decode(&deletedComment)
	if err != nil {
		return "", err
	}

	return deletedComment.ID.Hex(), nil
}

func (r *repositories) GetCommentsByVideoID(ctx context.Context, videoID string) (commentSlice []models.Comment, err error) {
	coll := r.db.Collection("comments")
	id, _ := primitive.ObjectIDFromHex(videoID)

	cursor, err := coll.Find(ctx, bson.D{{"videoID", id}})
	if err != nil {
		return nil, err
	}

	var comments []query.Comment
	err = cursor.All(ctx, &comments)
	if err != nil {
		return nil, err
	}

	for _, v := range comments {
		id := v.ID.Hex()
		userID := v.UserID.Hex()
		videoID := v.VideoID.Hex()

		createdAtStr := v.CreatedAt.Format("2022-10-23T17:20:36.524+00:00")
		updatedAtStr := v.UpdatedAt.Format("2022-10-23T17:20:36.524+00:00")

		comment := models.Comment{
			ID:          id,
			UserID:      &userID,
			VideoID:     &videoID,
			Description: &v.Description,
			CreatedAt:   createdAtStr,
			UpdatedAt:   updatedAtStr,
		}

		commentSlice = append(commentSlice, comment)
	}

	return commentSlice, nil
}

func (r *repositories) GetCommentByID(ctx context.Context, commentID string) (*models.Comment, error) {
	coll := r.db.Collection("comments")
	id, _ := primitive.ObjectIDFromHex(commentID)

	var comment query.Comment
	err := coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&comment)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	userID := comment.UserID.Hex()
	videoID := comment.VideoID.Hex()

	createdAtStr := comment.CreatedAt.Format("2022-10-23T17:20:36.524+00:00")
	updatedAtStr := comment.UpdatedAt.Format("2022-10-23T17:20:36.524+00:00")

	result := models.Comment{
		ID:          commentID,
		UserID:      &userID,
		VideoID:     &videoID,
		Description: &comment.Description,
		CreatedAt:   createdAtStr,
		UpdatedAt:   updatedAtStr,
	}

	return &result, nil
}
