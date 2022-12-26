package config

import (
	"os"

	"github.com/supermaxio/nflplayoffbracket/constants"
)

var mongoDbConnection string

func Setup() {
	mongoDbConnection = setMongoDbConnection()
}

func setMongoDbConnection() string {
	return os.Getenv(constants.MONGO_DB_CONNECTION_ENV)
}

func GetMongoDbConnection() string {
	return mongoDbConnection
}
