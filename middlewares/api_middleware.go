package middlewares

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/ramadhanalfarisi/go-codebase-kocak/helpers"
)

func ApiMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" || r.Method == "DELETE" {
			params := mux.Vars(r)
			id := params["id"]
			if id == "" {
				response := &helpers.Response{Code: 400, Status: "failed", Message: []string{"parameter :id have to entered"}}
				json, err := json.Marshal(response)
				if err != nil {
					log.Fatal(err)
				}
				w.WriteHeader(http.StatusBadRequest)
				w.Write(json)
				return
			} else {
				_, err := uuid.Parse(id)
				if err != nil {
					response := &helpers.Response{Code: 400, Status: "failed", Message: []string{"parameter :id invalid"}}
					json, err := json.Marshal(response)
					if err != nil {
						log.Fatal(err)
					}
					w.WriteHeader(http.StatusBadRequest)
					w.Write(json)
					return
				}
			}
		}

		w.Header().Add("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}
