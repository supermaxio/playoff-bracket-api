package database

import (
	"context"
	"log"

	"github.com/supermaxio/nflplayoffbracket/constants"
	"github.com/supermaxio/nflplayoffbracket/types"
	"go.mongodb.org/mongo-driver/bson"
)

func RefreshStandings(standings []types.TeamStandingDB) {
	coll := mongoClient.Database(constants.MONGO_DB_NAME).Collection(constants.STANDINGS_COLLECTION_NAME)

	cursor, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		panic(err)
	}

	var results []types.TeamStandingDB
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	if len(results) == 0 {
		for _, standing := range standings {
			_, err := coll.InsertOne(context.TODO(), standing)
			if err != nil {
				panic(err)
			}
		}
	} else {
		for _, standing := range standings {
			update := bson.D{{Key: "$set", Value: bson.D{{Key: "record", Value: standing.Record}}}}
			_, err := coll.UpdateByID(context.TODO(), standing.ID, update)
			if err != nil {
				panic(err)
			}
		}
	}

	log.Println("Successfully updated standings")
}
