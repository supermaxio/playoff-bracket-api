package constants

const (
	MONGO_DB_NAME             = "nflPlayoffBrackets"
	STANDINGS_COLLECTION_NAME = "standings"
	USERS_COLLECTION_NAME     = "users"

	ENV                     = "ENV"
	MONGO_DB_CONNECTION_ENV = "NFLPLAYOFFDBSECRET"
	JWT_KEY_ENV             = "NFLPLAYOFFJWTKEY"
	PORT_ENV                = "PORT"

	ENV_DEV   = "dev"
	ENV_TEST  = "test"
	ENV_STAGE = "stage"
	ENV_PROD  = "prod"

	DOMAIN_LOCAL  = "localhost"
	DOMAIN_DOCKER = "0.0.0.0"

	COOKIE_TOKEN = "playoff_bracket_token"
)
