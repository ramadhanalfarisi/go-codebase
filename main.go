package main

import (
	"log"

	"github.com/ramadhanalfarisi/go-codebase-kocak/app"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Your application running on http://localhost:8080")
	app := app.App{}
	app.ConnectDB()
	app.Routes()
	app.Run()
}