# Middlewares Documentation

This folder contains reusable middleware components used by the REST API and GraphQL server.

## Files

- `api_middleware.go` — request validation for REST API routes
- `auth_middleware.go` — simple bearer token authorization for REST API
- `graphql_middleware.go` — GraphQL middleware utilities for logging, recovery, CORS, and auth

## REST API Middleware

### `ApiMiddleware`

Defined in `api_middleware.go`.

Purpose:

- Validates that routes requiring an `:id` parameter (PUT, PATCH, DELETE) receive it.
- Returns HTTP 400 when the `id` parameter is missing.

Usage:

This middleware is attached to the REST API route groups in `app/api/router.go`:

- `v1.Use(middlewares.ApiMiddleware)`
- `auth.Use(middlewares.ApiMiddleware)`

### `AuthMiddleware`

Defined in `auth_middleware.go`.

Purpose:

- Verifies that the `Authorization` header exists
- Ensures the header contains a `Bearer` token
- Parses the JWT token using `helpers.ParseUserJWT`
- Stores parsed user claims in `c.Locals("userDetail", claims)` for downstream handlers

Behavior:

- Returns HTTP 401 if the header is missing
- Returns HTTP 401 if the token is invalid

## GraphQL Middleware

Defined in `graphql_middleware.go`.

This package provides middleware functions designed for the GraphQL HTTP handler.

### `Logger`

- Logs request method, path, status, and duration
- Skips logging for GraphQL introspection queries

### `Recovery`

- Recovers from panics inside GraphQL handlers
- Returns a JSON error response with HTTP 500 instead of crashing the server

### `CORS`

- Adds `Access-Control-Allow-Origin`, `Access-Control-Allow-Methods`, and `Access-Control-Allow-Headers`
- Responds to `OPTIONS` preflight requests with `204 No Content`

### `Auth`

- Allows `GET` requests and introspection queries through without auth
- Checks the `Authorization` header for a bearer token
- Calls `drivers.NewGrpcClient()` and the user gRPC middleware service
- Converts gRPC response into a local `userDetail` context value
- Writes a GraphQL-style error response when auth fails

### `Chain`

- Composes middleware functions into a single pipeline
- Applies middleware left-to-right

Example:

```go
chain := middlewares.Chain(
    middlewares.Recovery,
    middlewares.Logger,
    middlewares.CORS("*"),
    middlewares.Auth,
)
http.Handle("/graphql", chain(g.handler))
```

## How to Use These Middlewares

- REST API middleware is used in `app/api/router.go`.
- GraphQL middleware is used in `app/graphql/graphql.go`.

### Adding a new middleware

1. Add a new function in this folder with the appropriate signature:
   - `func(c fiber.Ctx) error` for Fiber REST middleware
   - `func(http.Handler) http.Handler` for GraphQL middleware
2. Import and use it in the relevant application entrypoint.

## Notes

- `ApiMiddleware` only validates existence of `:id` for modifying REST routes.
- `AuthMiddleware` only validates token format and JWT claims for REST API.
- GraphQL auth uses a gRPC middleware call, so it depends on the user gRPC service.
- GraphQL errors are returned in the official GraphQL error format.
