package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/youtube-api-golang/internal/repositories/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repositories) GetUserByUsername(ctx context.Context, username string) *query.User {
	coll := r.db.Collection("users")

	var result *query.User
	err := coll.FindOne(ctx, bson.D{{"username", username}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the username %s\n", username)
		return nil
	}
	if err != nil {
		fmt.Println(err)
	}

	return result
}

func (r *repositories) GetUserByEmail(ctx context.Context, email string) *query.User {
	coll := r.db.Collection("users")

	var result *query.User
	err := coll.FindOne(ctx, bson.D{{"email", email}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the email %s\n", email)
		return nil
	}
	if err != nil {
		fmt.Println(err)
	}

	return result
}

func (r *repositories) GetUserByID(ctx context.Context, userID string) (user *query.User, err error) {
	coll := r.db.Collection("users")
	id, _ := primitive.ObjectIDFromHex(userID)

	var result *query.User
	err = coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the userID %s\n", userID)
		return nil, errors.New("User tidak ditemukan")
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, nil
}

func (r *repositories) CreateUser(ctx context.Context, user query.User) (userID string, err error) {
	coll := r.db.Collection("users")

	result, err := coll.InsertOne(ctx, user)
	if err != nil {
		return "", errors.New("error when creating user")
	}

	userID = result.InsertedID.(primitive.ObjectID).Hex()

	return userID, nil
}

func (r *repositories) UpdateUser(ctx context.Context, user query.User) (updatedUser *query.User, err error) {
	coll := r.db.Collection("users")

	update := bson.M{
		"$set": user,
	}

	err = coll.FindOneAndUpdate(ctx, bson.D{{"_id", user.ID}}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&updatedUser)
	if err != nil {
		return nil, errors.New("error when updating user")
	}

	return updatedUser, nil
}

func (r *repositories) DeleteUser(ctx context.Context, userID string) (deletedUserID string, err error) {
	coll := r.db.Collection("users")
	id, _ := primitive.ObjectIDFromHex(userID)

	var deletedUser *query.User
	err = coll.FindOneAndDelete(ctx, bson.D{{"_id", id}}).Decode(&deletedUser)
	if err != nil {
		return "", errors.New("error when deleting user")
	}

	return deletedUser.ID.Hex(), nil
}

// TODO: make validation to prevent duplicate subscribers
func (r *repositories) Subscribe(ctx context.Context, params query.UpdateSub) error {
	coll := r.db.Collection("users")

	subscribingID, _ := primitive.ObjectIDFromHex(params.SubscribingID)
	subscribedID, _ := primitive.ObjectIDFromHex(params.SubscribedID)

	session, err := r.db.Client().StartSession()
	if err != nil {
		panic(err)
	}
	if err = session.StartTransaction(); err != nil {
		panic(err)
	}

	if err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		updateSubscribedUsers := bson.M{
			"$push": bson.M{
				"subscribedUsers": subscribedID,
			},
		}

		_, err := coll.UpdateOne(ctx, bson.D{{"_id", subscribingID}}, updateSubscribedUsers)
		if err != nil {
			fmt.Println("ERROR: ", err)
			session.AbortTransaction(sc)
			return err
		}

		updateSubscribers := bson.M{
			"$inc": bson.M{
				"subscribers": 1,
			},
		}

		_, err = coll.UpdateOne(ctx, bson.D{{"_id", subscribedID}}, updateSubscribers)
		if err != nil {
			fmt.Println("ERROR: ", err)
			session.AbortTransaction(sc)
			return err
		}

		if err = session.CommitTransaction(sc); err != nil {
			return err
		}

		return nil
	}); err != nil {
		fmt.Println(err)
	}
	session.EndSession(ctx)

	return nil
}

// TODO: subscriber count must not be negative
func (r *repositories) Unsubscribe(ctx context.Context, params query.UpdateSub) error {
	coll := r.db.Collection("users")

	subscribingID, _ := primitive.ObjectIDFromHex(params.SubscribingID)
	subscribedID, _ := primitive.ObjectIDFromHex(params.SubscribedID)

	session, err := r.db.Client().StartSession()
	if err != nil {
		panic(err)
	}
	if err = session.StartTransaction(); err != nil {
		panic(err)
	}

	if err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		updateSubscribedUsers := bson.M{
			"$pull": bson.M{
				"subscribedUsers": subscribedID,
			},
		}

		_, err := coll.UpdateOne(ctx, bson.D{{"_id", subscribingID}}, updateSubscribedUsers)
		if err != nil {
			fmt.Println("ERROR: ", err)
			session.AbortTransaction(sc)
			return err
		}

		updateSubscribers := bson.M{
			"$inc": bson.M{
				"subscribers": -1,
			},
		}

		_, err = coll.UpdateOne(ctx, bson.D{{"_id", subscribedID}}, updateSubscribers)
		if err != nil {
			fmt.Println("ERROR: ", err)
			session.AbortTransaction(sc)
			return err
		}

		if err = session.CommitTransaction(sc); err != nil {
			return err
		}

		return nil
	}); err != nil {
		fmt.Println(err)
	}
	session.EndSession(ctx)

	return nil
}
