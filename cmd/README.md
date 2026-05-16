# Command Line Interface Documentation

This folder contains Cobra CLI command definitions for the application entrypoints.

## Overview

The root command is defined in `cmd/root.go` as `codebase`. It registers four subcommands:

- `api` — starts the REST API server
- `graphql` — starts the GraphQL server
- `grpc` — starts the gRPC server
- `migrate` — runs database migrations

The CLI is executed through `main.go`, which loads environment variables and initializes configuration before invoking `cmd.RootCmd.Execute()`.

## Commands

### `go run main.go api`

- Starts the REST API server
- Uses `app/api.NewApi()` and `Api.Run()`
- Loads routes in `app/api/router.go`
- Listens on `PORT_API`

### `go run main.go graphql`

- Starts the GraphQL server
- Uses `app/graphql.NewGraphQL()` and `GraphQL.Run()`
- Serves `/graphql`
- Enables GraphQL playground
- Starts `pprof` on `localhost:6060`

### `go run main.go grpc`

- Starts the gRPC server
- Uses `app/grpc.NewGrpc()` and `Grpc.Run()`
- Registers user gRPC service via `services/user/routes/grpc.go`
- Enables health checks and reflection
- Listens on `PORT_GRPC`

### `go run main.go migrate`

- Runs database migrations using `app/migrate/migrate.go`
- Loads PostgreSQL connection from `DB_URL`
- Uses `MIGRATIONS_PATH` and `DB_NAME`
- Applies pending `*.up.sql` migrations from the `migrations/` folder

## Environment Setup

Before running any command, ensure `main.go` has loaded environment variables via `godotenv.Load()`. The following vars are expected:

- `DB_URL`
- `DB_NAME`
- `DB_SCHEMA`
- `MIGRATIONS_PATH`
- `PORT_API`
- `PORT_GRAPHQL`
- `PORT_GRPC`
- `DEBUG`
- `GRPC_SERVER`

A sample `.env.example` is available in the repository root.

## Build and Run

To run a command directly:

```bash
go run main.go api
```

To build the CLI binary:

```bash
go build -o codebase .
```

Then run:

```bash
./codebase api
```

## Notes

- The `cmd/` package only defines command metadata and the command run behavior.
- Actual application behavior is implemented in the `app/` package.
- If you add new service entrypoints, register them in `cmd/root.go`.
