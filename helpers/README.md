# Helpers Documentation

This folder contains reusable utility functions and helper types used across the application.

## Overview

The `helpers` package includes utilities for:

- JWT generation and parsing
- request/response handling
- pagination calculation
- query execution abstraction
- input validation
- password hashing and verification
- string conversion
- GraphQL argument parsing
- user detail extraction for API and GraphQL flows
- logging helpers

There is also a nested helper package in `helpers/query_builder` for building SQL queries.

## Files

- `converter.go` — convert values, e.g. `StringToInt`
- `encryption.go` — password hashing, verification, and SHA512 hashing
- `graphql_helper.go` — GraphQL argument collection and user detail extraction helpers
- `jwt.go` — JWT creation and parsing helpers
- `logging.go` — application logging helpers with debug mode support
- `pagination.go` — pagination request parsing and page metadata creation
- `query.go` — `QueryHelper` implementation for executing SQL commands and scanning results
- `response.go` — response wrapper for Fiber handlers
- `strings.go` — string utilities such as `ToCamelCase`
- `user_detail.go` — helpers for extracting authenticated user details from context
- `validator.go` — request model validation using `go-playground/validator`
- `query_builder/` — nested query builder utilities and examples

## Key Helpers

### JWT Helpers (`jwt.go`)

- `GenerateJWT(data UserDetail, expiryTime *time.Duration) (string, error)`
  - Creates a signed JWT token using `JWT_SECRET` from environment variables.
  - Default expiry is 30 minutes when `expiryTime` is nil.

- `ParseUserJWT(tokenString string) (jwt.MapClaims, error)`
  - Parses and validates an incoming JWT token.
  - Returns `jwt.MapClaims` if valid.

### User Detail Helpers (`graphql_helper.go` / `user_detail.go`)

- `GetUserDetail(claim jwt.MapClaims) UserDetail`
  - Converts JWT claims into a typed `UserDetail` struct.

- `GetUserIdFromAPI(c fiber.Ctx) int`
  - Reads `userDetail` from Fiber context locals.

- `GetUserIdFromGraphql(c context.Context) int`
  - Reads `userDetail` from GraphQL request context.

### Response Helpers (`response.go`)

- `Response.SendResponse(c fiber.Ctx) error`
  - Sends a consistent JSON response with HTTP status code.

### Query Helpers (`query.go`)

`QueryHelper` provides a thin abstraction around `database/sql`:

- `Insert(ctx, sql, args, params...)`
- `Select(ctx, sql, args, scan)`
- `SelectRow(ctx, sql, args, params...)`
- `Update(ctx, sql, args, params...)`
- `Delete(ctx, sql, args)`

It supports SQL queries with `RETURNING` and scans the returned row into provided output parameters.

### Pagination Helpers (`pagination.go`)

- `Pagination.GetPagination(r *http.Request) Pagination`
  - Parses `page` and `limit` query params.
  - Supports `limit=all` to disable pagination.

- `Pagination.CreatePagination() Pagination`
  - Calculates visible row count and first/last page flags.

### Validation Helpers (`validator.go`)

- `Validate(model any) ([]string, bool)`
  - Uses `go-playground/validator` with English translations.
  - Returns validation errors and success status.

### Encryption Helpers (`encryption.go`)

- `HashPassword(password string) (string, error)`
  - Hashes passwords using Argon2id.

- `VerifyPassword(dataPass ValidatePassword) (bool, error)`
  - Compares a plaintext password against an Argon2id hash.

- `SHA512(text string) string`
  - Computes SHA512 hash of a string.

### Logging Helpers (`logging.go`)

- `Info(msg string)`
- `Warning(msg string)`
- `Debug(msg string)`
- `Error(err error)`

`Debug` logs only when `config.DEBUG == "true"`.

### String Helpers (`strings.go`)

- `ToCamelCase(text string) string`
  - Converts text to camelCase by removing non-alphanumeric characters.

### Converter Helpers (`converter.go`)

- `StringToInt(s string) int`
  - Converts string to integer, returning `0` and logging the error on failure.

## Usage Examples

### Validate a request model

```go
msgs, ok := helpers.Validate(input)
if !ok {
    return helpers.Response{Code: 400, Status: "failed", Message: strings.Join(msgs, ", ")}.SendResponse(c)
}
```

### Send a JSON response

```go
return helpers.Response{Code: 200, Status: "success", Message: "OK", Data: result}.SendResponse(c)
```

### Extract authenticated user ID in API handlers

```go
userId := helpers.GetUserIdFromAPI(c)
```

### Extract authenticated user ID in GraphQL handlers

```go
userId := helpers.GetUserIdFromGraphql(ctx)
```

## Notes

- This helper package is intended for cross-cutting concerns used by controllers, services, and middleware.
- Keep helper functions small and reusable.
- For SQL query building, prefer the nested `helpers/query_builder` package.
