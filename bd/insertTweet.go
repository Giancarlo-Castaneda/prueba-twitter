package bd

import (
	"context"
	"time"

	"github.com/GicGa-iOS/prueba-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertTweet save the tweet in DB*/
func InsertTweet(t models.SaveTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	register := bson.M{
		"userid":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}

	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
