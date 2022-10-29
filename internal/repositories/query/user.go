package query

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty"`
	Username        string               `bson:"username"`
	Email           string               `bson:"email"`
	Password        string               `bson:"password"`
	Img             string               `bson:"img"`
	Subscribers     int                  `bson:"subscribers"`
	SubscribedUsers []primitive.ObjectID `bson:"subscribedUsers"`
	CreatedAt       time.Time            `bson:"createdAt"`
	UpdatedAt       time.Time            `bson:"updatedAt"`
}

func NewUser(user User) User {
	user.SubscribedUsers = []primitive.ObjectID{}

	return user
}

type UpdateSub struct {
	SubscribingID string
	SubscribedID  string
}
