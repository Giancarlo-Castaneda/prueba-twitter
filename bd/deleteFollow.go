package bd

import (
	"context"
	"time"

	"github.com/GicGa-iOS/prueba-twitter/models"
)

/*DeleteFollow delete the relationship in database*/
func DeleteFollow(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
