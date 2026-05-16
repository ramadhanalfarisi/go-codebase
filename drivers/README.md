# Drivers Documentation

This folder contains infrastructure drivers used by the application.

## Overview

The `drivers` package provides reusable connection helpers for external systems:

- PostgreSQL database (`connect_db.go`)
- Redis cache (`redis.go`)
- gRPC client connection (`grpc.go`)

These helpers are used by the application entrypoints and middleware to access external services.

## Files

- `connect_db.go` — PostgreSQL database connection helper
- `redis.go` — Redis client helpers for cache operations
- `grpc.go` — gRPC client creation for the user service

## PostgreSQL Driver

### `ConnectDB()`

Defined in `connect_db.go`.

Behavior:

- Opens a PostgreSQL connection using `config.DB_URL`
- Configures connection pooling:
  - max lifetime: 3 minutes
  - max open connections: 10
  - max idle connections: 10

Usage example:

```go
db := drivers.ConnectDB()
```

This function is used throughout the app for database access.

## Redis Driver

Defined in `redis.go`.

### `RedisConnection()`

- Creates a Redis client using `config.REDIS_URL` and `config.REDIS_PASSWORD`
- Uses database `0`
- Pings Redis on initialization

### Helper functions

- `SetRedisValue(key string, value string) bool`
- `GetRedisValue(key string) string`
- `DeleteRedisValue(keys []string) bool`
- `SearchRedisValue(keys string) []string`

Each helper opens a connection, performs the operation, and closes the client.

Usage examples:

```go
ok := drivers.SetRedisValue("user123", "cached-data")
value := drivers.GetRedisValue("user123")
```

## gRPC Driver

Defined in `grpc.go`.

### `NewGrpcClient()`

- Connects to the gRPC server at `config.GRPC_SERVER`
- Uses insecure transport credentials for local development
- Applies a unary client interceptor for logging
- Returns a `GrpcClient` struct and a cleanup function

Usage example:

```go
client, cleanup := drivers.NewGrpcClient()
defer cleanup()
resp, err := client.UserClient.Middleware(ctx, &grpc.MiddlewareInput{Token: token})
```

### `GrpcClient`

The returned struct contains:

- `UserClient` — a generated gRPC client for `UserController`
- `conn` — the underlying gRPC connection object

### Logging Interceptor

`clientLoggingInterceptor` logs outgoing RPC call details including method name, duration, and error.

## Required Environment Variables

The driver helpers depend on these config values:

- `DB_URL`
- `REDIS_URL`
- `REDIS_PASSWORD`
- `GRPC_SERVER`

These are loaded from the application environment and initialized in `config.InitializeDatabaseConfig()` and `config.InitializeCacheConfig()`.

## Notes

- `ConnectDB()` returns a raw `*sql.DB`, so callers should manage query helpers and transactions.
- Redis helpers currently open and close the client on every call.
- The gRPC client is designed for local/insecure usage; update `grpc.WithTransportCredentials` for production.
