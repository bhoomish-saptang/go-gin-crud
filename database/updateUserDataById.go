package database

import (
	"context"
	"fmt"

	"github.com/go-gin/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUserDetailsByID(id string, userDetails constants.User) (constants.User, error) {

	filter := bson.M{"id": id}
	//var data constants.User
	result, err := collection.ReplaceOne(context.Background(), filter, userDetails)
	if err != nil {
		return constants.User{}, err
	}

	fmt.Println(".Printing a upserted ID", result.UpsertedID)
	return userDetails, nil
}
