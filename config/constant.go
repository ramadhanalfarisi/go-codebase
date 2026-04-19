package config

import (
	"os"
)

var (
	DEBUG           string
	MIGRATIONS_PATH string
	PORT_APP        string
	ENVIRONMMENT    string
)

func InitializeConstants() {
	DEBUG = os.Getenv("DEBUG")                     // is debug ?
	MIGRATIONS_PATH = os.Getenv("MIGRATIONS_PATH") // migrations path
	PORT_APP = os.Getenv("PORT_APP")
	ENVIRONMMENT = os.Getenv("ENVIRONMENT") // your environment (testing, development, production)
}
