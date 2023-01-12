package database

import (
	"context"
	"log"

	"github.com/supermaxio/nflplayoffbracket/config"
	"github.com/supermaxio/nflplayoffbracket/constants"
	"github.com/supermaxio/nflplayoffbracket/types"
	"go.mongodb.org/mongo-driver/bson"
)

func GetStandings() ([]types.TeamStanding, error) {
	coll := mongoClient.Database(config.GetMongoDbName()).Collection(constants.STANDINGS_COLLECTION_NAME)
	cursor, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []types.TeamStanding{}, err
	}

	var results []types.TeamStanding
	if err = cursor.All(context.TODO(), &results); err != nil {
		return []types.TeamStanding{}, err
	}

	return results, nil
}

func RefreshStandings(standings []types.TeamStanding) error {
	coll := mongoClient.Database(config.GetMongoDbName()).Collection(constants.STANDINGS_COLLECTION_NAME)
	results, err := GetStandings()
	if err != nil {
		return err
	}

	if len(results) == 0 {
		for _, standing := range standings {
			_, err := coll.InsertOne(context.TODO(), standing)
			if err != nil {
				return err
			}
		}
	} else {
		for _, standing := range standings {
			updateModel := types.TeamStandingUpdateDB{
				Rank:         standing.Rank,
				Record:       standing.Record,
				Location:     standing.Location,
				Name:         standing.Name,
				Abbreviation: standing.Abbreviation,
				DisplayName:  standing.DisplayName,
				Conference:   standing.Conference,
			}
			updateByID := bson.D{{Key: "id", Value: standing.ID}}
			update := bson.D{{Key: "$set", Value: updateModel}}

			_, err := coll.UpdateOne(context.TODO(), updateByID, update)
			if err != nil {
				return err
			}
		}
	}

	log.Println("Successfully updated standings")

	return nil
}
