package config

import (
	"os"
)

var (
	DEBUG           string
	MIGRATIONS_PATH string
	PORT_API        string
	ENVIRONMMENT    string
	PORT_GRAPHQL    string
	PORT_GRPC       string
)

func InitializeConstants() {
	DEBUG = os.Getenv("DEBUG")                     // is debug ?
	MIGRATIONS_PATH = os.Getenv("MIGRATIONS_PATH") // migrations path
	PORT_API = os.Getenv("PORT_API")
	ENVIRONMMENT = os.Getenv("ENVIRONMENT") // your environment (testing, development, production)
	PORT_GRAPHQL = os.Getenv("PORT_GRAPHQL")
	PORT_GRPC = os.Getenv("PORT_GRPC")
}
