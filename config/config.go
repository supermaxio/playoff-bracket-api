package config

import (
	"os"

	"github.com/supermaxio/nflplayoffbracket/constants"
)

var mongoDbConnection string
var jwtKey string

func Setup() {
	mongoDbConnection = setMongoDbConnection()
	jwtKey = setJwtKey()
}

func setMongoDbConnection() string {
	return os.Getenv(constants.MONGO_DB_CONNECTION_ENV)
}

func GetMongoDbConnection() string {
	return mongoDbConnection
}

func setJwtKey() string {
	return os.Getenv(constants.JWT_KEY_ENV)
}

func GetJwtKey() string {
	return jwtKey
}
