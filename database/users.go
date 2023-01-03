package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/supermaxio/nflplayoffbracket/constants"
	"github.com/supermaxio/nflplayoffbracket/types"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user types.User) types.User {
	coll := mongoClient.Database(constants.MONGO_DB_NAME).Collection(constants.USERS_COLLECTION_NAME)

	// validation
	user.Username = strings.ToLower(user.Username)

	_, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}

	createdUser := FindUser(user.Username)

	log.Printf(fmt.Sprintf("Successfully added user: %s", createdUser.Username))

	return createdUser
}

func FindUser(username string) (resultUser types.User) {
	coll := mongoClient.Database(constants.MONGO_DB_NAME).Collection(constants.USERS_COLLECTION_NAME)

	//validation
	username = strings.ToLower(username)

	err := coll.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}).Decode(&resultUser)
	if err != nil {
		return
	}

	return
}
