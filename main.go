package main

import (
	"log"

	"github.com/ramadhanalfarisi/go-codebase-kocak/app"
)

func main() {
	// Set log settings
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Your application running on http://localhost:8080")

	// Identify App struct
	app := app.App{}
	// Create database connection
	app.ConnectDB()
	// Create all API routers
	app.Routes()
	// Run the service
	app.Run()
}