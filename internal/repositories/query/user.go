package query

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty"`
	Username        string               `bson:"username,omitempty"`
	Email           string               `bson:"email,omitempty"`
	Password        string               `bson:"password,omitempty"`
	Img             string               `bson:"img,omitempty,omitempty"`
	Subscribers     int                  `bson:"subscribers,omitempty"`
	SubscribedUsers []primitive.ObjectID `bson:"subscribedUsers,omitempty"`
	CreatedAt       time.Time            `bson:"createdAt,omitempty"`
	UpdatedAt       time.Time            `bson:"updatedAt,omitempty"`
}

type UpdateSub struct {
	SubscribingID string
	SubscribedID  string
}
