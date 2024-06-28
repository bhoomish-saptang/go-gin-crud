package database

import (
	"context"

	"github.com/go-gin/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllUserDetails() ([]constants.User, error) {

	AllUserDetails := []constants.User{}

	result, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return AllUserDetails, err
	}
	for result.Next(context.Background()) {
		var User constants.User
		result.Decode(&User)
		AllUserDetails = append(AllUserDetails, User)
	}
	return AllUserDetails, nil

}
