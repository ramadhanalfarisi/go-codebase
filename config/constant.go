package config

import (
	"os"
)

var (
	DEBUG              string
	MIGRATIONS_PATH    string
	PORT_API           string
	ENVIRONMMENT       string
	PORT_GRAPHQL       string
	PORT_GRPC          string
	GRPC_SERVER        string
	PPROF_API_PORT     string
	PPROF_GRAPHQL_PORT string
	PPROF_GRPC_PORT    string
)

func InitializeConstants() {
	DEBUG = os.Getenv("DEBUG")                     // is debug ?
	MIGRATIONS_PATH = os.Getenv("MIGRATIONS_PATH") // migrations path
	PORT_API = os.Getenv("PORT_API")
	ENVIRONMMENT = os.Getenv("ENVIRONMENT") // your environment (testing, development, production)
	PORT_GRAPHQL = os.Getenv("PORT_GRAPHQL")
	PORT_GRPC = os.Getenv("PORT_GRPC")
	GRPC_SERVER = os.Getenv("GRPC_SERVER")
	PPROF_API_PORT = os.Getenv("PPROF_API_PORT")
	PPROF_GRAPHQL_PORT = os.Getenv("PPROF_GRAPHQL_PORT")
	PPROF_GRPC_PORT = os.Getenv("PPROF_GRPC_PORT")
}
