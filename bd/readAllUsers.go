package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/GicGa-iOS/prueba-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ReadAllUsers read the registered users in the system, you receive "R" in those who bring only returns those related*/
func ReadAllUsers(ID string, page int64, search string, kind string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search}, //(?i) indica que no distinga mayusculas
	}
	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var found, includ bool

	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		includ = false

		found, err = SearchRelationship(r)
		if kind == "new" && !found {
			includ = true
		}
		if kind == "follow" && found {
			includ = true
		}

		if r.UserRelationID == ID {
			includ = false
		}

		if includ {
			s.Password = ""
			s.Biography = ""
			s.WebSite = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cur.Close(ctx)
	return results, true
}
