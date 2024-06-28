package database

import (
	"context"
	"fmt"
	"log"

	"github.com/go-gin/constants"
)

func InsertUserData(data constants.User) error {
	insertResult, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Println("error occur in MongoDB Insertion document", err.Error())
		return err

	}
	fmt.Println("Inserted document ID:", insertResult.InsertedID)
	return nil
}
