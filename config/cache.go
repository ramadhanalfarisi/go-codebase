package config

import (
	"os"
)

var (
	REDIS_URL      string
	REDIS_PASSWORD string
)

func InitializeCacheConfig() {
	REDIS_URL = os.Getenv("REDIS_URL")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
}
