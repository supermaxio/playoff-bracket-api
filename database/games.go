package database

import (
	"context"
	"fmt"
	"log"

	"github.com/supermaxio/nflplayoffbracket/config"
	"github.com/supermaxio/nflplayoffbracket/constants"
	"github.com/supermaxio/nflplayoffbracket/types"
	"go.mongodb.org/mongo-driver/bson"
)

func GetGames() ([]types.Game, error) {
	coll := mongoClient.Database(config.GetMongoDbName()).Collection(constants.GAMES_COLLECTION_NAME)

	cursor, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []types.Game{}, err
	}

	var results []types.Game
	if err = cursor.All(context.TODO(), &results); err != nil {
		return []types.Game{}, err
	}

	return results, nil
}

func CreateGame(game types.Game) (types.Game, error) {
	coll := mongoClient.Database(config.GetMongoDbName()).Collection(constants.GAMES_COLLECTION_NAME)

	// validation

	_, err := coll.InsertOne(context.TODO(), game)
	if err != nil {
		return types.Game{}, err
	}

	createdGame, _ := FindGame(game.ID)

	log.Printf(fmt.Sprintf("Successfully added game: %d", createdGame.ID))

	return createdGame, nil
}

func UpdateGame(gameID int, game types.Game) (types.Game, error) {
	collection := mongoClient.Database(config.GetMongoDbName()).Collection(constants.GAMES_COLLECTION_NAME)

	// validation

	updateByGamename := bson.D{{Key: "id", Value: gameID}}
	update := bson.D{{Key: "$set", Value: game}}
	_, err := collection.UpdateOne(context.TODO(), updateByGamename, update)
	if err != nil {
		return types.Game{}, err
	}

	updatedGame, err := FindGame(gameID)
	if err != nil {
		return types.Game{}, err
	}

	return updatedGame, nil
}

func FindGame(gameID int) (resultGame types.Game, err error) {
	coll := mongoClient.Database(config.GetMongoDbName()).Collection(constants.GAMES_COLLECTION_NAME)

	//validation

	err = coll.FindOne(context.TODO(), bson.D{{Key: "id", Value: gameID}}).Decode(&resultGame)
	if err != nil {
		return
	}

	return
}

func DeleteGame(gameID int) (int, error) {
	collection := mongoClient.Database(config.GetMongoDbName()).Collection(constants.GAMES_COLLECTION_NAME)

	// validation

	query := bson.D{{Key: "id", Value: gameID}}
	deletedResult, err := collection.DeleteOne(context.TODO(), query)
	if err != nil {
		return 0, err
	}

	return int(deletedResult.DeletedCount), nil
}
