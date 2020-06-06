package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN is the connection object to DB*/
var MongoCN = ConnectBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://gianadmin:admin@cluster0-epfo5.mongodb.net/<dbname>?retryWrites=true&w=majority")

/*ConnectBD allows connection to the DB*/
func ConnectBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Println("Success connection to DB")
	return client
}

/*CheckConnection ping to DB*/
func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
