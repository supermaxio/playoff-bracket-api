package database

import (
	"context"
	"log"
	"strings"

	"github.com/supermaxio/nflplayoffbracket/config"
	"github.com/supermaxio/nflplayoffbracket/constants"
	"github.com/supermaxio/nflplayoffbracket/types"
	"go.mongodb.org/mongo-driver/bson"
)

func FindBracket(username string) (resultBracket types.Bracket, err error) {
	collection := mongoClient.Database(config.GetMongoDbName()).Collection(constants.BRACKETS_COLLECTION_NAME)

	//validation
	username = strings.ToLower(username)

	query := bson.D{{Key: "username", Value: username}}
	err = collection.FindOne(context.TODO(), query).Decode(&resultBracket)
	if err != nil {
		log.Println(err.Error())
		return
	}

	return
}

func CreateBracket(bracket types.Bracket) (types.Bracket, error) {
	collection := mongoClient.Database(config.GetMongoDbName()).Collection(constants.BRACKETS_COLLECTION_NAME)

	bracket.Username = strings.ToLower(bracket.Username)

	_, err := collection.InsertOne(context.TODO(), bracket)
	if err != nil {
		log.Println(err.Error())
		return types.Bracket{}, err
	}

	createdBracket, err := FindBracket(bracket.Username)
	if err != nil {
		return types.Bracket{}, err
	}

	return createdBracket, nil
}

func UpdateBracket(bracket types.Bracket) (types.Bracket, error) {
	collection := mongoClient.Database(config.GetMongoDbName()).Collection(constants.BRACKETS_COLLECTION_NAME)

	// validation
	bracket.Username = strings.ToLower(bracket.Username)

	updateByUsername := bson.D{{Key: "username", Value: bracket.Username}}
	update := bson.D{{Key: "$set", Value: bracket}}
	_, err := collection.UpdateOne(context.TODO(), updateByUsername, update)
	if err != nil {
		log.Println(err.Error())
		return types.Bracket{}, err
	}

	updatedBracket, err := FindBracket(bracket.Username)
	if err != nil {
		return types.Bracket{}, err
	}

	return updatedBracket, nil
}

func DeleteBracket(username string) (int, error) {
	collection := mongoClient.Database(config.GetMongoDbName()).Collection(constants.BRACKETS_COLLECTION_NAME)

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
