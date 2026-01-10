package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/supermaxio/nflplayoffbracket/config"
	"github.com/supermaxio/nflplayoffbracket/constants"
	"github.com/supermaxio/nflplayoffbracket/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMessages() ([]types.User, error) {
	collection := mongoClient.Database(config.GetMongoDbName()).Collection(constants.MESSAGES_COLLECTION_NAME)
	// create a bson.D to specify the sort
	sort := bson.D{{Key: "rank", Value: 1}} // 1 for ascending, -1 for descending

	// create the find options
	findOptions := options.Find().SetSort(sort)

	// perform the find
	cursor, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return []types.User{}, err
	}

	var results []types.User
	if err = cursor.All(context.TODO(), &results); err != nil {
		return []types.User{}, err
	}

	return results, nil
}

func CreateMessage(user types.User) (types.User, error) {
	collection := mongoClient.Database(config.GetMongoDbName()).Collection(constants.MESSAGES_COLLECTION_NAME)

	// validation
	user.Username = strings.ToLower(user.Username)

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return types.User{}, err
	}

	createdUser, _ := FindUser(user.Username)

	log.Printf(fmt.Sprintf("Successfully added user: %s", createdUser.Username))

	return createdUser, nil
}

func FindMessageByUser(username string) (resultUser types.User, err error) {
	collection := mongoClient.Database(config.GetMongoDbName()).Collection(constants.MESSAGES_COLLECTION_NAME)

	//validation
	username = strings.ToLower(username)

	err = collection.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}).Decode(&resultUser)
	if err != nil {
		return
	}

	return
}

func DeleteMessage(username string) (int, error) {
	collection := mongoClient.Database(config.GetMongoDbName()).Collection(constants.MESSAGES_COLLECTION_NAME)

	// validation
	username = strings.ToLower(username)

	query := bson.D{{Key: "username", Value: username}}
	deletedResult, err := collection.DeleteOne(context.TODO(), query)
	if err != nil {
		return 0, err
	}

	return int(deletedResult.DeletedCount), nil
}
