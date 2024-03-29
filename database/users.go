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

func GetUsers() ([]types.User, error) {
	collection := mongoClient.Database(config.GetMongoDbName()).Collection(constants.USERS_COLLECTION_NAME)
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

func CreateUser(user types.User) (types.User, error) {
	collection := mongoClient.Database(config.GetMongoDbName()).Collection(constants.USERS_COLLECTION_NAME)

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

func UpdateUser(username string, user types.UserUpdate) (types.User, error) {
	collection := mongoClient.Database(config.GetMongoDbName()).Collection(constants.USERS_COLLECTION_NAME)

	// validation

	updateByUsername := bson.D{{Key: "username", Value: strings.ToLower(username)}}
	update := bson.D{{Key: "$set", Value: user}}
	_, err := collection.UpdateOne(context.TODO(), updateByUsername, update)
	if err != nil {
		return types.User{}, err
	}

	updatedUser, err := FindUser(strings.ToLower(username))
	if err != nil {
		return types.User{}, err
	}

	return updatedUser, nil
}

func FindUser(username string) (resultUser types.User, err error) {
	collection := mongoClient.Database(config.GetMongoDbName()).Collection(constants.USERS_COLLECTION_NAME)

	//validation
	username = strings.ToLower(username)

	err = collection.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}).Decode(&resultUser)
	if err != nil {
		return
	}

	return
}

func DeleteUser(username string) (int, error) {
	collection := mongoClient.Database(config.GetMongoDbName()).Collection(constants.USERS_COLLECTION_NAME)

	// validation
	username = strings.ToLower(username)

	query := bson.D{{Key: "username", Value: username}}
	deletedResult, err := collection.DeleteOne(context.TODO(), query)
	if err != nil {
		return 0, err
	}

	return int(deletedResult.DeletedCount), nil
}
