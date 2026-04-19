package config

import (
	"os"
)

var (
	DB_URL  string
	DB_NAME string
	DB_SCHEMA string
)

func InitializeDatabaseConfig() {
	DB_URL = os.Getenv("DB_URL")
	DB_NAME = os.Getenv("DB_NAME")
	DB_SCHEMA = os.Getenv("DB_SCHEMA")
}
