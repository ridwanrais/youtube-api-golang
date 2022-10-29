package query

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Video struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Title       string               `bson:"title"`
	Description string               `bson:"description"`
	VideoURL    string               `bson:"videoUrl"`
	ImgURL      string               `bson:"imgUrl"`
	Views       []primitive.ObjectID `bson:"views"`
	Dislikes    []primitive.ObjectID `bson:"dislikes"`
	Likes       []primitive.ObjectID `bson:"likes"`
	Tags        []string             `bson:"tags"`
	UserID      primitive.ObjectID   `bson:"userID"`
	CreatedAt   time.Time            `bson:"createdAt"`
	UpdatedAt   time.Time            `bson:"updatedAt"`
}

func NewVideo(video Video) Video {
	if video.Tags == nil {
		video.Tags = []string{}
	}

	return Video{
		Title:       video.Title,
		Description: video.Description,
		VideoURL:    video.VideoURL,
		ImgURL:      video.ImgURL,
		Views:       []primitive.ObjectID{},
		Dislikes:    []primitive.ObjectID{},
		Likes:       []primitive.ObjectID{},
		Tags:        video.Tags,
		UserID:      video.UserID,
		CreatedAt:   video.CreatedAt,
		UpdatedAt:   video.UpdatedAt,
	}
}

type UpdateView struct {
	VideoID string
	UserID  string
}

type UpdateLike struct {
	VideoID string
	UserID  string
}
