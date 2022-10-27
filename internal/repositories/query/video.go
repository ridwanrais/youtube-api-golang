package query

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Video struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Title       string               `bson:"title,omitempty"`
	Description string               `bson:"description,omitempty"`
	VideoURL    string               `bson:"videoUrl,omitempty"`
	ImgURL      string               `bson:"imgUrl,omitempty"`
	Views       []primitive.ObjectID `bson:"views,omitempty"`
	Dislikes    []primitive.ObjectID `bson:"dislikes,omitempty"`
	Likes       []primitive.ObjectID `bson:"likes,omitempty"`
	Tags        []string             `bson:"tags,omitempty"`
	UserID      primitive.ObjectID   `bson:"userID,omitempty"`
	CreatedAt   time.Time            `bson:"createdAt,omitempty"`
	UpdatedAt   time.Time            `bson:"updatedAt,omitempty"`
}

type UpdateView struct {
	VideoID string
	UserID  string
}

type UpdateLike struct {
	VideoID string
	UserID  string
}
