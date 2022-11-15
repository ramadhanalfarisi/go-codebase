package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ramadhanalfarisi/go-codebase-kocak/helpers"
)

func main() {
	response := &helpers.Response{Code: 400, Status: "failed", Message: "parameter :id invalid"}
	json,err := json.Marshal(response)
	if err != nil {
		log.Fatal(json)
	}
	fmt.Println(string(json))
}