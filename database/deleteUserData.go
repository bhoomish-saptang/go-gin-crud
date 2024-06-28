package database

import (
	"context"
	"log"

	"github.com/go-gin/constants"
)

func DeleteUserData(data constants.User) {
	_, err := collection.DeleteOne(context.Background(), data)
	if err != nil {
		log.Println("error occur in MongoDB Insertion document", err.Error())
	}

}
