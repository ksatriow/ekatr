# Ekatr

Ekatr is a electronic kasir trenggalek project that implements Domain-Driven Design (DDD) architecture. It is designed to provide robust, scalable, and maintainable code by dividing the application into layers and contexts.

## Project Structure

Ekatr's project structure is organized based on DDD principles to ensure a clear separation of concerns. Here is an overview of the structure:

```sh
ekatr/
├── cmd/
│ └── main.go
├── internal/
│ ├── application/
│ │ └── user/
│ │ ├── dto.go
│ │ └── service.go
│ ├── domain/
│ │ └── user/
│ │ ├── user.go
│ │ └── user_repository.go
│ ├── infrastructure/
│ │ └── persistence/
│ │ └── postgresql/
│ │ └── user_repository.go
│ ├── interfaces/
│ │ └── http/
│ │ ├── handler.go
│ │ └── router.go
│ └── utils/
│ └── jwt.go
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```


## Domain-Driven Design (DDD) Architecture

DDD is a software design approach focusing on modeling the business domain and its logic. In Ekatr, we have implemented the following layers:

- **Domain Layer**: Contains the core business logic and domain objects.
- **Application Layer**: Contains the application services which coordinate the use cases.
- **Infrastructure Layer**: Contains the implementations of repository interfaces and other external integrations.
- **Interface Layer**: Contains the HTTP handlers and routing logic.

### Domain Layer

The domain layer contains the core business logic of the application. For example, the `user` package defines the `User` entity and the `UserRepository` interface.

### Application Layer

The application layer contains the service classes that handle the use cases of the application. For example, the `UserService` class handles user registration and login logic.

### Infrastructure Layer

The infrastructure layer contains the implementations of the repository interfaces defined in the domain layer. For example, the `postgresql` package contains the `UserRepository` implementation that interacts with a PostgreSQL database.

### Interface Layer

The interface layer contains the HTTP handlers and routing logic. For example, the `http` package defines the `UserHandler` class which handles the HTTP requests related to user operations.

## Setting Up the Project

### Prerequisites

- Docker
- Docker Compose
- Make

### Environment Variables

Create a `.env` file in the root of the project and set the following env.sample
