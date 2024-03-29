package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/ramadhanalfarisi/go-codebase-kocak/helpers"
	"github.com/ramadhanalfarisi/go-codebase-kocak/models"
)

func AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] != nil {
			authorization := r.Header.Get("Authorization")
			if !strings.Contains(authorization, "Bearer") {
				response := &models.Response{Code: 401, Status: "failed", Message: []string{"Token must be Bearer type"}}
				json, err := json.Marshal(response)
				if err != nil {
					log.Println(err)
					return
				}
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(json)
			} else {
				tokenString := strings.Replace(authorization, "Bearer ", "", -1)
				decodeToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
					if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Signing Method invalid")
					} else if method != helpers.JWT_SIGNING_METHOD {
						return nil, fmt.Errorf("Signing Method invalid")
					}

					err_claims := t.Claims.Valid()
					if err_claims != nil {
						return nil, err_claims
					}

					return helpers.JWT_SIGNATURE_KEY, nil
				})
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
				claims, ok := decodeToken.Claims.(jwt.MapClaims)

				if !ok || !decodeToken.Valid {
					http.Error(w, "Not Valid", http.StatusBadGateway)
				}

				ctx := context.WithValue(context.Background(), "userDetail", claims)
				r = r.WithContext(ctx)
				handler.ServeHTTP(w, r)
			}
		} else {
			response := &models.Response{Code: 401, Status: "failed", Message: []string{"Authorization header is required"}}
			json, err := json.Marshal(response)
			if err != nil {
				log.Println(err)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(json)
		}
	})
}
