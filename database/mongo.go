package database

import (
	"context"
	"log"

	"github.com/supermaxio/nflplayoffbracket/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func MongoConnect() {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.GetMongoDbConnection()))
	if err != nil {
		panic(err)
	}

	mongoClient = client
	log.Println("Successfully connected to mongo db.")

	// Ping the primary
	MongoPing()
}

func MongoPing() {
	if err := mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		MongoDisconnect()
		panic(err)
	}
	log.Println("Successfully pinged mongo db.")
}

func MongoDisconnect() {
	if err := mongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	log.Println("Successfully Disconnected from mongo db.")
}
