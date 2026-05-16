# Services Folder Documentation

## Overview

The `services` folder contains the core business logic for the application. It is organized using a clean architecture approach, separating request handling, business rules, and data access into distinct layers.

This documentation reflects the current implementation in the repository, including REST API routes, GraphQL support, and gRPC service registration.

---

## Table of Contents

1. [Architecture Pattern](#architecture-pattern)
2. [File Structure](#file-structure)
3. [Service Responsibilities](#service-responsibilities)
4. [REST API Services](#rest-api-services)
5. [GraphQL Service](#graphql-service)
6. [gRPC Service](#grpc-service)
7. [Layer-by-Layer Breakdown](#layer-by-layer-breakdown)
8. [Service Implementation Notes](#service-implementation-notes)

---

## Architecture Pattern

The services layer follows Clean Architecture principles:

```
Request → Controller → Usecase → Repository → Database
```

- **Controller**: Converts transport requests into usecase calls.
- **Usecase**: Contains business logic and service orchestration.
- **Repository**: Handles data persistence and database access.

### Benefits

- Clear separation of concerns
- Easier testing and mocking
- Reusable business logic across transport adapters
- Explicit dependencies via interfaces

---

## File Structure

```
services/
├── common/
│   └── models/
│       └── model.go                # Shared models and GraphQL object helpers
│
├── product/
│   ├── controller/
│   │   ├── controller_api.go       # Product REST API controller
│   │   ├── controller_graphql.go   # Product GraphQL controller
│   │   └── controller_interface.go # Product controller interface
│   │
│   ├── models/
│   │   └── model.go                # Product data model definitions
│   │
│   ├── repository/
│   │   ├── repository.go           # Product repository implementation
│   │   └── repository_interface.go # Product repository interface
│   │
│   ├── routes/
│   │   ├── api.go                  # Product REST route definitions
│   │   └── graphql.go              # Product GraphQL definition builder
│   │
│   └── usecase/
│       ├── usecase.go              # Product business logic
│       └── usecase_interface.go    # Product usecase interface
│
└── user/
    ├── controller/
    │   ├── controller.go           # User REST controller
    │   └── controller_interface.go # User controller interface
    │
    ├── grpc/
    │   └── grpc_controller.go      # gRPC controller interface implementation
    │
    ├── models/
    │   └── model.go                # User data model definitions
    │
    ├── repository/
    │   ├── repository.go           # User repository implementation
    │   └── repository_interface.go # User repository interface
    │
    ├── routes/
    │   ├── api.go                  # User REST route definitions
    │   └── grpc.go                 # User gRPC route registration
    │
    └── usecase/
        ├── usecase.go              # User business logic
        ├── usecase_grpc.go         # User gRPC usecase implementation
        └── usecase_interface.go    # User usecase interface
```

---

## Service Responsibilities

### Product Service

- Provides REST API operations for product resources.
- Exposes GraphQL query and mutation definitions for products.
- Implements controllers, usecases, and repository access for product data.

### User Service

- Provides authentication-related REST operations (`/register`, `/login`).
- Exposes a gRPC controller implementation for user service methods.
- Implements controllers, usecases, and repository access for user data.

---

## REST API Services

The REST API routes are mounted in `app/api/router.go`.

### API routing setup

- Base path: `/api/v1`
- Auth path: `/api/auth/v1`

User routes are registered under the auth group:

- `POST /api/auth/v1/register`
- `POST /api/auth/v1/login`

Product routes are registered under the main API group:

- `GET /api/v1/products`
- `GET /api/v1/products/:id`
- `POST /api/v1/products`
- `PATCH /api/v1/products/:id`
- `PUT /api/v1/products/:id`
- `DELETE /api/v1/products/:id`

### Product REST routes definition

Defined in `services/product/routes/api.go`, the product route builder wires:

- `repository.NewProductRepository`
- `usecase.NewProductUsecase`
- `controller.NewProductControllerAPI`

That controller implements GET, POST, PATCH, PUT, and DELETE handlers.

### User REST routes definition

Defined in `services/user/routes/api.go`, the user route builder wires:

- `repository.NewUserRepository`
- `usecase.NewUserUsecase`
- `controller.NewUserController`

Then it exposes authentication endpoints for registration and login.

---

## GraphQL Service

The GraphQL server is initialized in `app/graphql/graphql.go` and uses the route definitions built by `services/product/routes/graphql.go`.

### GraphQL endpoint

- `POST /graphql`
- GraphQL playground is enabled

### Product GraphQL schema

The product schema exposes:

- `products` → list of products
- `product` → product query by `id`
- `createProduct` → create a new product
- `updateProduct` → update an existing product
- `deleteProduct` → delete a product

### GraphQL route builder

`services/product/routes/graphql.go` returns:

- `productQuery []models.GraphQLObjectModel`
- `productMutation []models.GraphQLObjectModel`

These are merged into the GraphQL schema by `app/graphql/root.go`.

### Shared GraphQL model helper

`services/common/models/model.go` defines reusable GraphQL input/object helpers such as `GraphQLObjectModel` used by GraphQL route builders.

---

## gRPC Service

The gRPC server is initialized in `app/grpc/grpc.go`.

### gRPC service registration

`services/user/routes/grpc.go` returns a gRPC service implementation:

- `UserGrpcRoute() grpc.UserControllerServer`

This user gRPC service is registered in `app/grpc/grpc.go` via:

- `ug.RegisterUserControllerServer(g.App, userService)`

### gRPC middleware and features

The gRPC server includes:

- Unary interceptors for logging and panic recovery
- Stream interceptor logging
- Health checking via `grpc_health_v1.RegisterHealthServer`
- Reflection support via `reflection.Register`

---

## Layer-by-Layer Breakdown

### 1. Routes Layer (`services/*/routes`)

**Purpose:** Wire dependencies and expose transport-specific endpoints.

**Responsibilities:**

- Construct repositories, usecases, and controllers
- Define HTTP and gRPC endpoints
- Register route handlers with the application router

### 2. Controller Layer (`services/*/controller`)

**Purpose:** Handle request parsing, validation, response formatting, and delegation to usecases.

**Responsibilities:**

- Parse HTTP/GraphQL request payloads
- Validate user input
- Call usecase methods
- Return structured responses

**Key detail:** Controller code should not contain repository or SQL logic.

### 3. Usecase Layer (`services/*/usecase`)

**Purpose:** Implement core business rules and data orchestration.

**Responsibilities:**

- Execute business flows
- Validate domain constraints
- Transform data between models and repository formats
- Use repository interfaces for persistence

**Example:** `services/product/usecase/usecase.go` orchestrates product creation, update, retrieval, and deletion.

### 4. Repository Layer (`services/*/repository`)

**Purpose:** Encapsulate database access logic.

**Responsibilities:**

- Build and execute SQL queries
- Map query results to domain models
- Return persistence errors to usecases

**Example:** `services/product/repository/repository.go` implements product CRUD using the query builder.

---

## Service Implementation Notes

- **Product service** supports both REST and GraphQL transport adapters.
- **User service** supports REST for authentication and gRPC for service-to-service calls.
- The GraphQL service currently exposes product operations only.
- The REST API is mounted on `/api/v1` with authentication routes under `/api/auth/v1`.
- gRPC registration is centralized in `app/grpc/grpc.go`.
- Shared models and GraphQL helper types are placed in `services/common/models`.

---

## How to Extend Services

1. Add domain models to `services/<domain>/models`.
2. Implement repository methods in `services/<domain>/repository`.
3. Add business logic in `services/<domain>/usecase`.
4. Add transport handlers in `services/<domain>/controller`.
5. Register new REST routes in `services/<domain>/routes/api.go`.
6. Register new GraphQL definitions in `services/<domain>/routes/graphql.go`.
7. Register new gRPC routes in `services/<domain>/routes/grpc.go` and `app/grpc/grpc.go`.

---

## Contact

If you need to update or extend service behavior, the authoritative implementation entrypoints are:

- `app/api/router.go`
- `app/graphql/root.go`
- `app/grpc/grpc.go`

Use these files as the starting point for transport-level integration with the `services` layer.
