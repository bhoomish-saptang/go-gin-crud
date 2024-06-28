package database

import (
	"context"
	"log"

	"github.com/go-gin/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func FindUserDetailsByID(id string) (constants.User, error) {
	var findedUser constants.User
	filter := bson.M{"id": id}
	err := collection.FindOne(context.Background(), filter).Decode(&findedUser)
	if err != nil {
		log.Println("error occur in MongoDB User details document", err)
		return findedUser, err
	}
	return findedUser, nil
}
