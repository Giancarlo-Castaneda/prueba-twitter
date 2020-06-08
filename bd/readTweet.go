package bd

import (
	"context"
	"log"
	"time"

	"github.com/GicGa-iOS/prueba-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ReadTweet reads tweets from a profile*/
func ReadTweet(ID string, page int64) ([]*models.ReturnTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	var results []*models.ReturnTweet

	condition := bson.M{
		"userid": ID,
	}

	option := options.Find()
	option.SetLimit(20)
	option.SetSort(bson.D{{Key: "date", Value: -1}})
	option.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, option)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var register models.ReturnTweet
		err := cursor.Decode(&register)
		if err != nil {
			return results, false
		}
		results = append(results, &register)
	}
	return results, true
}
