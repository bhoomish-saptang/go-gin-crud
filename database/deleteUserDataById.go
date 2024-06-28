package database

import (
	"context"
	"log"

	"github.com/go-gin/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteUserDetailsByID(id string) (constants.User, error) {
	var findedUser constants.User
	filter := bson.M{"id": id}
	err := collection.FindOneAndDelete(context.Background(), filter).Decode(&findedUser)
	if err != nil {
		log.Println("error occur in MongoDB Insertion document", err)
		return findedUser, err
	}
	return findedUser, nil

}
