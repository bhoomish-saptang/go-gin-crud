package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection
var CollectionAuth *mongo.Collection

func ConnectMongodb() {
	db_name := os.Getenv("DB_NAME")
	collection_name_auth := os.Getenv("COLLECTION_NAME_AUTH")
	collection_name := os.Getenv("COLLECTION_NAME")
	var err error

	clientOptions := options.Client().ApplyURI(os.Getenv("URI"))
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println("Unable to MongoDb client: ", err.Error())
	}

	collection = client.Database(db_name).Collection(collection_name)
	CollectionAuth = client.Database(db_name).Collection(collection_name_auth)

}
