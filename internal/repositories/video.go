package repositories

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/youtube-api-golang/internal/models"
	"github.com/youtube-api-golang/internal/repositories/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repositories) AddVideo(ctx context.Context, video query.Video) (videoID string, err error) {
	coll := r.db.Collection("videos")

	result, err := coll.InsertOne(ctx, video)
	if err != nil {
		return "", err
	}

	videoID = result.InsertedID.(primitive.ObjectID).Hex()

	return videoID, nil
}

func (r *repositories) GetVideoByID(ctx context.Context, videoID string) (video *models.Video, err error) {
	coll := r.db.Collection("videos")
	id, _ := primitive.ObjectIDFromHex(videoID)

	var result *query.Video
	err = coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the videoID %s\n", videoID)
		return nil, errors.New("Video tidak ditemukan")
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var dislikes []string
	for _, v := range result.Dislikes {
		dislikes = append(dislikes, v.Hex())
	}

	var likes []string
	for _, v := range result.Likes {
		likes = append(likes, v.Hex())
	}

	userIDString := result.UserID.Hex()
	viewCount := int64(len(result.Views))

	videoResponse := &models.Video{
		ID:          result.ID.Hex(),
		Description: &result.Description,
		Dislikes:    dislikes,
		Likes:       likes,
		ImgURL:      &result.ImgURL,
		Tags:        result.Tags,
		Title:       &result.Title,
		UserID:      &userIDString,
		VideoURL:    &result.VideoURL,
		Views:       &viewCount,
	}

	return videoResponse, nil
}

func (r *repositories) UpdateVideo(ctx context.Context, video query.Video, videoID string) (updatedVideo *models.Video, err error) {
	coll := r.db.Collection("videos")

	video.ID, _ = primitive.ObjectIDFromHex(videoID)
	update := bson.M{
		"$set": video,
	}

	var updatedVideoMongo query.Video

	err = coll.FindOneAndUpdate(ctx, bson.D{{"_id", video.ID}}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&updatedVideoMongo)
	if err != nil {
		return nil, err
	}

	var dislikes []string
	for _, v := range updatedVideoMongo.Dislikes {
		dislikes = append(dislikes, v.Hex())
	}

	var likes []string
	for _, v := range updatedVideoMongo.Likes {
		likes = append(likes, v.Hex())
	}

	userIDString := updatedVideoMongo.UserID.Hex()
	viewCount := int64(len(updatedVideoMongo.Views))

	updatedVideo = &models.Video{
		ID:          videoID,
		Description: &updatedVideoMongo.Description,
		Dislikes:    dislikes,
		Likes:       likes,
		ImgURL:      &updatedVideoMongo.ImgURL,
		Tags:        updatedVideoMongo.Tags,
		Title:       &updatedVideoMongo.Title,
		UserID:      &userIDString,
		VideoURL:    &updatedVideoMongo.VideoURL,
		Views:       &viewCount,
	}

	return updatedVideo, nil
}

func (r *repositories) DeleteVideo(ctx context.Context, videoID string) (deletedVideoID string, err error) {
	coll := r.db.Collection("videos")
	id, _ := primitive.ObjectIDFromHex(videoID)

	var deletedVideo *query.Video
	err = coll.FindOneAndDelete(ctx, bson.D{{"_id", id}}).Decode(&deletedVideo)
	if err != nil {
		return "", err
	}

	return deletedVideo.ID.Hex(), nil
}

func (r *repositories) UpdateView(ctx context.Context, params query.UpdateView) (viewCount *int, err error) {
	coll := r.db.Collection("videos")

	videoID, _ := primitive.ObjectIDFromHex(params.VideoID)
	userID, _ := primitive.ObjectIDFromHex(params.UserID)
	var video *query.Video

	updateView := bson.M{
		"$addToSet": bson.M{
			"views": userID,
		},
	}

	err = coll.FindOneAndUpdate(ctx, bson.D{{"_id", videoID}}, updateView).Decode(&video)
	if err != nil {
		return nil, err
	}

	views := len(video.Views)
	return &views, nil
}

func (r *repositories) GetRandomVideos(ctx context.Context, limit int64) (videoSlice []models.Video, err error) {
	coll := r.db.Collection("videos")

	pipeline := []bson.D{{{"$sample", bson.D{{"size", limit}}}}}
	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var videos []query.Video
	err = cursor.All(ctx, &videos)
	if err != nil {
		return nil, err
	}

	for _, v := range videos {

		var dislikes []string
		for _, v := range v.Dislikes {
			dislikes = append(dislikes, v.Hex())
		}

		var likes []string
		for _, v := range v.Likes {
			likes = append(likes, v.Hex())
		}

		userIDString := v.UserID.Hex()
		viewCount := int64(len(v.Views))

		video := models.Video{
			ID:          v.ID.Hex(),
			Description: &v.Description,
			Dislikes:    dislikes,
			Likes:       likes,
			ImgURL:      &v.ImgURL,
			Tags:        v.Tags,
			Title:       &v.Title,
			UserID:      &userIDString,
			VideoURL:    &v.VideoURL,
			Views:       &viewCount,
		}

		videoSlice = append(videoSlice, video)
	}

	return videoSlice, nil
}

func (r *repositories) GetTrendingVideos(ctx context.Context, limit int64) (videoSlice []models.Video, err error) {
	coll := r.db.Collection("videos")

	cursor, err := coll.Find(ctx, bson.D{{}}, options.Find().SetLimit(limit))
	if err != nil {
		return nil, err
	}

	var videos []query.Video
	err = cursor.All(ctx, &videos)
	if err != nil {
		return nil, err
	}

	for _, v := range videos {
		var dislikes []string
		for _, v := range v.Dislikes {
			dislikes = append(dislikes, v.Hex())
		}

		var likes []string
		for _, v := range v.Likes {
			likes = append(likes, v.Hex())
		}

		userIDString := v.UserID.Hex()
		viewCount := int64(len(v.Views))

		video := models.Video{
			ID:          v.ID.Hex(),
			Description: &v.Description,
			Dislikes:    dislikes,
			Likes:       likes,
			ImgURL:      &v.ImgURL,
			Tags:        v.Tags,
			Title:       &v.Title,
			UserID:      &userIDString,
			VideoURL:    &v.VideoURL,
			Views:       &viewCount,
		}

		videoSlice = append(videoSlice, video)
	}

	sort.Slice(videoSlice, func(i, j int) bool {
		return *videoSlice[i].Views > *videoSlice[j].Views
	})

	fmt.Println("TESTTTTTTTTTTTTTT", videoSlice[0])

	return videoSlice, nil
}

func (r *repositories) GetVideosFromSubscribedChannels(ctx context.Context, subscribedChannels []string) (videos []models.Video, err error) {
	coll := r.db.Collection("videos")

	var totalVideos []query.Video
	for _, ownerID := range subscribedChannels {
		ownerIDObject, _ := primitive.ObjectIDFromHex(ownerID)

		cursor, err := coll.Find(ctx, bson.D{{"userID", ownerIDObject}})
		if err != nil {
			return nil, err
		}

		var resVideos []query.Video
		err = cursor.Decode(&resVideos)
		if err != nil {
			return nil, err
		}

		totalVideos = append(totalVideos, resVideos...)
	}

	sort.Slice(totalVideos, func(i, j int) bool {
		return totalVideos[i].CreatedAt.Before(totalVideos[i].CreatedAt)
	})

	var videoSlice []models.Video
	for _, v := range totalVideos {

		var dislikes []string
		for _, v := range v.Dislikes {
			dislikes = append(dislikes, v.Hex())
		}

		var likes []string
		for _, v := range v.Likes {
			likes = append(likes, v.Hex())
		}

		userIDString := v.UserID.Hex()
		viewCount := int64(len(v.Views))

		createdAtStr := v.CreatedAt.Format("2022-10-23T17:20:36.524+00:00")
		updatedAtStr := v.UpdatedAt.Format("2022-10-23T17:20:36.524+00:00")

		video := models.Video{
			ID:          v.ID.Hex(),
			Description: &v.Description,
			Dislikes:    dislikes,
			Likes:       likes,
			ImgURL:      &v.ImgURL,
			Tags:        v.Tags,
			Title:       &v.Title,
			UserID:      &userIDString,
			VideoURL:    &v.VideoURL,
			Views:       &viewCount,
			CreatedAt:   createdAtStr,
			UpdatedAt:   updatedAtStr,
		}

		videoSlice = append(videoSlice, video)
	}

	return videoSlice, nil
}