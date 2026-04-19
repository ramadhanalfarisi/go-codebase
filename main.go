package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ramadhanalfarisi/go-codebase/cmd"
	"github.com/ramadhanalfarisi/go-codebase/config"
)

func main() {
	// Set log settings
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	config.InitializeDatabaseConfig()
	config.InitializeCacheConfig()
	config.InitializeConstants()

	// Run the root command
	cmd.RootCmd.Execute()
}
