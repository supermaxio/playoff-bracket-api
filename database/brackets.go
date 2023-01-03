package database

import (
	"context"
	"strings"

	"github.com/supermaxio/nflplayoffbracket/constants"
	"github.com/supermaxio/nflplayoffbracket/types"
	"go.mongodb.org/mongo-driver/bson"
)

func FindBracket(username string) (resultBracket types.Bracket) {
	collection := mongoClient.Database(constants.MONGO_DB_NAME).Collection(constants.BRACKETS_COLLECTION_NAME)

	//validation
	username = strings.ToLower(username)

	query := bson.D{{Key: "username", Value: username}}
	err := collection.FindOne(context.TODO(), query).Decode(&resultBracket)
	if err != nil {
		return
	}

	return
}

func CreateBracket(bracket types.Bracket) types.Bracket {
	collection := mongoClient.Database(constants.MONGO_DB_NAME).Collection(constants.BRACKETS_COLLECTION_NAME)

	bracket.Username = strings.ToLower(bracket.Username)

	_, err := collection.InsertOne(context.TODO(), bracket)
	if err != nil {
		panic(err)
	}

	createdBracket := FindBracket(bracket.Username)

	return createdBracket
}

func UpdateBracket(bracket types.Bracket) types.Bracket {
	collection := mongoClient.Database(constants.MONGO_DB_NAME).Collection(constants.BRACKETS_COLLECTION_NAME)

	// validation
	bracket.Username = strings.ToLower(bracket.Username)

	updateByUsername := bson.D{{Key: "username", Value: bracket.Username}}
	update := bson.D{{Key: "$set", Value: bracket}}
	_, err := collection.UpdateOne(context.TODO(), updateByUsername, update)
	if err != nil {
		panic(err)
	}

	updatedBracket := FindBracket(bracket.Username)

	return updatedBracket
}

func DeleteBracket(username string) int {
	collection := mongoClient.Database(constants.MONGO_DB_NAME).Collection(constants.BRACKETS_COLLECTION_NAME)

	// validation
	username = strings.ToLower(username)

	query := bson.D{{Key: "username", Value: username}}
	deletedResult, err := collection.DeleteOne(context.TODO(), query)
	if err != nil {
		panic(err)
	}

	return int(deletedResult.DeletedCount)
}
