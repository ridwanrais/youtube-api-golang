package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/youtube-api-golang/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repositories) GetHealthCheck(ctx context.Context) (*models.Health, error) {

	coll := r.client.Database("sample_mflix").Collection("movies")
	title := "Back to the Future"

	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
	}

	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonData)

	return nil, nil
}
