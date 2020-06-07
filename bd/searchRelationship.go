package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/GicGa-iOS/prueba-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*SearchRelationship search relationship between two users*/
func SearchRelationship(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("relation")

	condition := bson.M{
		"userid":         t.UserID,
		"userrelationid": t.UserRelationID,
	}

	var result models.Relation
	fmt.Println(result)
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
