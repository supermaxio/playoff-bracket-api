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
)

func GetUsers() ([]types.User, error) {
	coll := mongoClient.Database(config.GetMongoDbName()).Collection(constants.USERS_COLLECTION_NAME)

	cursor, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Println(err.Error())
		return []types.User{}, err
	}

	var results []types.User
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Println(err.Error())
		return []types.User{}, err
	}

	return results, nil
}

func CreateUser(user types.User) (types.User, error) {
	coll := mongoClient.Database(config.GetMongoDbName()).Collection(constants.USERS_COLLECTION_NAME)

	// validation
	user.Username = strings.ToLower(user.Username)

	_, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(err.Error())
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
		log.Println(err.Error())
		return types.User{}, err
	}

	updatedUser, err := FindUser(strings.ToLower(username))
	if err != nil {
		return types.User{}, err
	}

	return updatedUser, nil
}

func FindUser(username string) (resultUser types.User, err error) {
	coll := mongoClient.Database(config.GetMongoDbName()).Collection(constants.USERS_COLLECTION_NAME)

	//validation
	username = strings.ToLower(username)

	err = coll.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}).Decode(&resultUser)
	if err != nil {
		log.Println(err.Error())
		return
	}

	return
}

func DeleteUser(username string)  (int, error) {
	collection := mongoClient.Database(config.GetMongoDbName()).Collection(constants.USERS_COLLECTION_NAME)

	// validation
	username = strings.ToLower(username)

	query := bson.D{{Key: "username", Value: username}}
	deletedResult, err := collection.DeleteOne(context.TODO(), query)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	return int(deletedResult.DeletedCount), nil
}
