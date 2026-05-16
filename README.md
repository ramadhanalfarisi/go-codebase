# Go Codebase: Comprehensive Backend Software Engineering Portfolio

## Overview

This repository is a Go backend portfolio project that demonstrates clean architecture, modular service design, and multiple runtime entrypoints for REST API, GraphQL, and gRPC.

## Architecture

The repository is organized by clean architecture principles, separating business logic, transport adapters, and infrastructure:

- **Entities & Models**: domain definitions in `services/*/models`
- **Usecases**: application business logic in `services/*/usecase`
- **Interface Adapters**: route handlers and controllers in `services/*/controller` and `services/*/routes`
- **Frameworks & Drivers**: network, database, and cache initialization in `app/`, `config/`, and `drivers/`

## Key Features

- REST API with Fiber
- GraphQL endpoint using `graphql-go`
- gRPC server with health checks and reflection
- PostgreSQL database support
- Redis caching support
- Database migrations
- Middleware for logging, recovery, rate limiting, CORS, and auth
- Runtime profiling support via `pprof` on the GraphQL server

## Project Structure

```
├── app/                    # Application entrypoints and transport adapters
│   ├── api/                # REST API application and routing
│   ├── graphql/            # GraphQL application and schema setup
│   ├── grpc/               # GRPC application and service setup
│   └── migrate/            # Database migration logic
├── cmd/                    # Cobra CLI commands for each runtime
├── config/                 # Environment configuration and constants
├── drivers/                # Database and cache connections
├── helpers/                # Utility helpers
├── middlewares/            # HTTP middleware definitions
├── migrations/             # SQL migration files
├── services/               # Business logic services per domain
│   ├── product/            # Product service implementation
│   └── user/               # User service implementation
└── containers/             # Docker Compose files for local services
```

## Getting Started

### Prerequisites
- Go 1.19+
- PostgreSQL
- Redis (optional)
- Docker and Docker Compose (optional)

### Setup

1. Clone the repository:
```bash
git clone https://github.com/ramadhanalfarisi/go-codebase.git
cd go-codebase
```

2. Install Go dependencies:
```bash
go mod download
```

3. Create a `.env` file and set required environment variables, including:
- `PORT_API`
- `PORT_GRAPHQL`
- `PORT_GRPC`
- `DEBUG`
- `MIGRATIONS_PATH`
- `ENVIRONMENT`
- `GRPC_SERVER`

4. Run database migrations:
```bash
go run main.go migrate
```

5. Start the REST API server:
```bash
go run main.go api
```

6. Start the GraphQL server:
```bash
go run main.go graphql
```

7. Start the gRPC server:
```bash
go run main.go grpc
```

### Using Docker Compose

Docker Compose files are available in `containers/`.

For PostgreSQL only:
```bash
docker-compose -f containers/docker-compose-postgresql.yml up -d
```

If you have Docker Compose files for API or GraphQL services, use the corresponding YAML file in `containers/`.

## CLI Commands

- `go run main.go migrate` — run database migrations
- `go run main.go api` — run the REST API server
- `go run main.go graphql` — run the GraphQL server
- `go run main.go grpc` — run the gRPC server

## Endpoints

### REST API
- Base path: `/api/v1`
- Auth path: `/api/auth/v1`

The exact route definitions are in `app/api/router.go`, `services/user/routes`, and `services/product/routes`.

### GraphQL
- Endpoint: `/graphql`
- The GraphQL runtime also starts a `pprof` listener on `localhost:6060`

### gRPC
- Port: configured by `PORT_GRPC`
- Health check and reflection are enabled for service discovery

## Profiling with pprof

The GraphQL server starts a profiling listener on `localhost:6060`.

To collect a 30-second CPU profile:
```bash
curl -o cpu.pprof "http://localhost:6060/debug/pprof/profile?seconds=30"
go tool pprof cpu.pprof
```

To collect a heap profile:
```bash
curl -o heap.pprof "http://localhost:6060/debug/pprof/heap"
go tool pprof heap.pprof
```

To use pprof directly against the live endpoint:
```bash
go tool pprof http://localhost:6060/debug/pprof/profile
```

Common interactive pprof commands:
- `top`
- `list <function>`
- `web`
- `pdf`

> Note: `pprof` is enabled by the API runtime because it starts `http.ListenAndServe(":6060", nil)`.

## Notes

- `main.go` loads environment variables and initializes configuration before invoking the Cobra root command.
- The GraphQL command uses GraphQL middleware and serves `/graphql`.
- The REST API command serves routes under `/api/v1` and `/api/auth/v1`.

## Contact

- **LinkedIn**: https://www.linkedin.com/in/ramadhan-salman-alfarisi-69520117a/
- **Email**: ramadhansalmanalfarisi8@gmail.com
- **Medium**: https://medium.com/@ramadhansalmanalfarisi8
