package middlewares

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/ramadhanalfarisi/go-codebase/constants"
	"github.com/ramadhanalfarisi/go-codebase/drivers"
	"github.com/ramadhanalfarisi/go-codebase/helpers"
	"github.com/ramadhanalfarisi/go-codebase/services/user/grpc"
)

// Logger logs every request with method, path, status, and duration.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Read body to check operation name
		body, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(body)) // restore

		// Skip logging introspection queries
		if strings.Contains(string(body), "IntrospectionQuery") ||
			strings.Contains(string(body), "__schema") {
			next.ServeHTTP(w, r)
			return
		}

		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rw, r)

		log.Printf("%s %s | %d | %v",
			r.Method, r.URL.Path,
			rw.status,
			time.Since(start),
		)
	})
}

// Recovery catches panics and returns a 500 instead of crashing the server.
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("[Recovery] panic: %v\n%s", rec, debug.Stack())
				http.Error(w,
					`{"errors":[{"message":"internal server error"}]}`,
					http.StatusInternalServerError,
				)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// CORS adds Cross-Origin headers — adjust origins for production.
func CORS(allowedOrigins ...string) func(http.Handler) http.Handler {
	origins := "*"
	if len(allowedOrigins) > 0 {
		origins = allowedOrigins[0]
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", origins)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Request-ID")

			// Preflight
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// Chain composes multiple middleware into one, executed left → right.
//
//	chain := Chain(Recovery, RequestID, Logger, Auth, CORS())
//	http.Handle("/query", chain(graphqlHandler))
func Chain(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(final http.Handler) http.Handler {
		// Apply in reverse so the first middleware is outermost
		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}
		return final
	}
}

// ── Helpers ───────────────────────────────────────────────────────────────────

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
			return
		}

		// Read body
		body, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(body)) // restore

		// Allow introspection to pass without auth
		if strings.Contains(string(body), "IntrospectionQuery") ||
			strings.Contains(string(body), "__schema") {
			next.ServeHTTP(w, r)
			return
		}

		// Example: Get user from token
		token := r.Header.Get("Authorization")
		if token == "" {
			helpers.Error(fmt.Errorf("Authorization empty"))
			writeGraphQLError(w, constants.InvalidToken)
			return
		}
		grpcClient, cleanup := drivers.NewGrpcClient()
		defer cleanup() // conn.Close() called when handler is done

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		respUser, err := grpcClient.UserClient.Middleware(ctx, &grpc.MiddlewareInput{
			Token: token,
		})
		if err != nil {
			helpers.Error(err)
			writeGraphQLError(w, constants.InvalidToken)
			return
		}
		user := helpers.UserDetail{
			Id:    int(respUser.Id),
			Email: respUser.Email,
			Roles: respUser.Roles,
		}

		// Add to context
		ctx2 := context.WithValue(r.Context(), "userDetail", user)
		next.ServeHTTP(w, r.WithContext(ctx2))
	})

}

// writeGraphQLError returns a proper GraphQL error response format
func writeGraphQLError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // GraphQL always returns 200
	json.NewEncoder(w).Encode(map[string]any{
		"data": nil,
		"errors": []map[string]any{
			{
				"message": message,
				"extensions": map[string]any{
					"code": "UNAUTHENTICATED",
				},
			},
		},
	})
}
