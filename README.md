# Go Codebase: Comprehensive Backend Software Engineering Portfolio

## Overview

This repository represents my comprehensive work experience in backend software engineering, showcasing a robust API application built with Go that incorporates modern design patterns, architectural principles, and a wide range of backend technologies. This project serves as a demonstration of best practices in backend development, from clean architecture to advanced communication protocols.

## Architecture

This codebase is built following **CLEAN Architecture** principles, which emphasizes separation of concerns, dependency inversion, and maintainability. The architecture is divided into layers:

- **Entities**: Core business logic and domain models
- **Use Cases**: Application-specific business rules
- **Interface Adapters**: Controllers, presenters, and gateways
- **Frameworks & Drivers**: External frameworks, databases, and external interfaces

## Key Features and Technologies

This project demonstrates expertise in various software engineering domains:

### Design Patterns
- **Repository Pattern**: Abstract data access layer for database operations
- **Dependency Injection**: Loose coupling through interfaces and injection
- **Middleware Pattern**: Request processing pipeline for authentication, logging, etc.
- **Factory Pattern**: Object creation abstractions

### API Development
- **RESTful APIs**: Standard HTTP-based API endpoints
- **OpenAPI Specification**: API documentation and contract definition
- **GraphQL**: Flexible query language for APIs
- **RPC (Remote Procedure Call)**: Direct method invocation over network

### Real-time Communication
- **WebSocket**: Bidirectional communication for real-time features
- **Pub/Sub (Publish-Subscribe)**: Event-driven architecture for decoupled services

### Performance and Scalability
- **Caching**: Redis-based caching for improved performance
- **Database Optimization**: Efficient query building and connection management

### Event-Driven Architecture
- **Webhook Events**: HTTP callbacks for external integrations
- **Message Queues**: Asynchronous processing and service communication

## Project Structure

```
├── app/                    # Application layer
│   ├── api/               # API handlers and routing
│   ├── graphql/           # GraphQL handlers and routing
│   └── migrate/           # Database migration logic
├── cmd/                   # Command-line interfaces
├── config/                # Configuration management
├── drivers/               # Driver configurations and helpers
├── helpers/               # Utility functions and helpers
├── middlewares/           # HTTP middleware components
├── migrations/            # Database migration files
├── services/              # Business logic services
│   ├── common/            # Shared models and utilities
│   └── user/              # User service with MVC-like structure
│       ├── controller/    # Request handlers
│       ├── models/        # Data models
│       ├── repository/    # Data access layer
│       ├── routes/        # Route definitions
│       └── usecase/       # Business logic
└── test/                  # Test files
```

## Getting Started

### Prerequisites
- Go 1.19+
- Docker and Docker Compose
- Redis (for caching)
- PostgreSQL (or your preferred database)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/ramadhanalfarisi/go-codebase.git
cd go-codebase
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables (create a `.env` based on `.env.example` file)

4. Run database migrations:
```bash
go run main.go migrate
```

5. Start the application:
```bash
go run main.go api
```

### Using Docker

```bash
docker-compose up -d
```

## API Documentation

The API is documented using OpenAPI specification. Access the documentation at `/swagger` when the server is running.

### Key Endpoints

- `API /api` - API
- `WebSocket /ws` - Real-time communication
- `POST /graphql` - GraphQL queries
- `OPEN API /open-api` - OPEN API

## Technologies Used

- **Language**: Go
- **Framework**: Fiber (HTTP framework)
- **Database**: PostgreSQL
- **Cache**: Redis
- **Message Queue**: (Implementation depends on specific Pub/Sub needs)
- **API Documentation**: Swagger/OpenAPI
- **Testing**: Go testing framework
- **Containerization**: Docker

## Design Patterns Implementation

### Repository Pattern
Located in `services/*/repository/`, this pattern abstracts data access, allowing for easy testing and database switching.

### Clean Architecture Layers
- **Domain Layer**: Business entities and rules
- **Application Layer**: Use cases and application logic
- **Infrastructure Layer**: External concerns (database, web framework)

### Dependency Injection
Interfaces are defined for all major components, allowing for easy mocking and testing.

## Real-time Features

### WebSocket Implementation
Real-time communication is handled through WebSocket connections, enabling server-to-client communication for real-time data updates.

### Pub/Sub Pattern
Event-driven architecture allows for decoupled service communication, improving scalability and maintainability.

## Caching Strategy

Redis is used for caching frequently accessed data, reducing database load and improving response times.

## Webhook Integration

The application supports webhook events for integrating with external services, enabling event-driven workflows.

## GraphQL API

A GraphQL endpoint provides flexible data querying capabilities, allowing clients to request exactly the data they need.

## RPC Implementation

Remote Procedure Calls are implemented for direct service-to-service communication in distributed systems.

## Testing

The project includes comprehensive unit tests, especially for the query builder and core business logic.

## Conclusion

This codebase represents a culmination of my experience in backend software engineering, demonstrating proficiency in modern backend development practices, architectural patterns, and a wide range of server-side technologies. It serves as a reference implementation for building scalable, maintainable, and feature-rich backend applications and APIs.

## Contact

- **LinkedIn**: https://www.linkedin.com/in/ramadhan-salman-alfarisi-69520117a/
- **Email**: ramadhansalmanalfarisi8@gmail.com
- **Medium**: https://medium.com/@ramadhansalmanalfarisi8

