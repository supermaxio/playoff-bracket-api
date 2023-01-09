package config

import (
	"os"

	"github.com/supermaxio/nflplayoffbracket/constants"
)

var mongoDbConnection string
var mongoDbName string
var jwtKey string
var port string
var env string
var domain string

func Setup() {
	env = setEnv()
	mongoDbConnection = setMongoDbConnection()
	mongoDbName = setMongoDbName()
	jwtKey = setJwtKey()
	port = setPort()
	domain = setDomain()
}

func setEnv() string {
	envString := os.Getenv(constants.ENV)
	if envString == "" {
		return constants.ENV_DEV
	} else {
		return envString
	}
}

func GetEnv() string {
	return env
}

func setMongoDbConnection() string {
	return os.Getenv(constants.MONGO_DB_CONNECTION_ENV)
}

func setMongoDbName() string {
	switch env {
	case constants.ENV_STAGE,
		constants.ENV_PROD:
		return constants.MONGO_DB_NAME_PROD
	default:
		return constants.MONGO_DB_NAME_DEV
	}
}

func setDomain() string {
	switch env {
	case constants.ENV_TEST,
		constants.ENV_STAGE,
		constants.ENV_PROD:
		return constants.DOMAIN_DOCKER
	default:
		return constants.DOMAIN_LOCAL
	}
}

func GetDomain() string {
	return domain
}

func GetMongoDbConnection() string {
	return mongoDbConnection
}

func GetMongoDbName() string {
	return mongoDbName
}

func setJwtKey() string {
	return os.Getenv(constants.JWT_KEY_ENV)
}

func GetJwtKey() string {
	return jwtKey
}

func setPort() string {
	envPort := os.Getenv(constants.PORT_ENV)
	if envPort == "" {
		return "8080"
	} else {
		return envPort
	}
}

func GetPort() string {
	return port
}
