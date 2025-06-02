# URL Shortener

A simple, scalable URL shortening service built with Go, Redis, and Docker, following Clean Architecture principles. The application provides an API to create shortened URLs and redirect to the original URLs using a short code. It includes unit tests and a Docker Compose setup for easy deployment.

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Setup](#setup)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Clean Architecture](#clean-architecture)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Shorten URLs**: Create a shortened URL (8-character code) from a given original URL.
- **Redirect URLs**: Redirect from a short code to the original URL.
- **Redis Storage**: Uses Redis to store URL mappings for fast retrieval.
- **Clean Architecture**: Organized with clear separation of concerns (domain, use cases, infrastructure).
- **Unit Tests**: Includes unit tests for the service and repository layers using `redismock`.
- **Dockerized**: Deployable with Docker and Docker Compose for easy setup.

## Project Structure

```
url-shortener/
├── cmd/
│   └── api/
│       └── main.go              # Entry point for the API server
├── internal/
│   ├── domain/
│   │   └── url/
│   │       ├── entity.go        # URL entity definition
│   │       ├── repository.go    # Repository interface
│   │       └── service.go       # Business logic for URL operations
│   ├── infrastructure/
│   │   ├── redis/
│   │   │   └── redis_repository.go  # Redis implementation of the repository
│   │   └── http/
│   │       └── handler.go       # HTTP handlers for API endpoints
│   └── usecases/
│       └── shorten_url.go       # Use case layer for URL operations
├── tests/
│   ├── redis_repository_test.go # Unit tests for Redis repository
│   └── url_service_test.go      # Unit tests for URL service
├── docker-compose.yml           # Docker Compose configuration
├── Dockerfile                   # Dockerfile for building the Go application
├── go.mod                       # Go module dependencies
├── go.sum                       # Go module checksums
└── url-shortener-api.http       # HTTP test file for API endpoints
```

## Prerequisites

- **Go**: Version 1.20 or higher
- **Docker**: Latest version with Docker Compose
- **Git**: For cloning the repository
- **Redis**: Managed via Docker Compose (no separate installation needed)
- **VS Code** (optional): For using the `.http` file with the REST Client extension

## Setup

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/marcosvfn/url-shortener.git
   cd url-shortener
   ```

2. **Install Go Dependencies**:
   Ensure you have Go installed, then run:

   ```bash
   go mod tidy
   ```

   This will download the required dependencies (`github.com/go-redis/redis/v8`, `github.com/gorilla/mux`, and `github.com/go-redis/redismock/v8`).

3. **Verify Project Structure**:
   Ensure all files and directories listed in the [Project Structure](#project-structure) section are present. If any are missing, create them as per the provided code.

## Running the Application

1. **Using Docker Compose**:

   ```bash
   docker-compose up --build
   ```

   This builds and starts the Go application (on `http://localhost:8080`) and a Redis instance (on `redis:6379`).

2. **Manually (without Docker)**:

   - Start a Redis server locally (e.g., `redis-server` or use a managed Redis instance).
   - Update the Redis address in `cmd/api/main.go` if necessary (default is `redis:6379`).
   - Run the application:
     ```bash
     go run ./cmd/api
     ```

   The server will start on `http://localhost:8080`.

## API Endpoints

The application exposes two endpoints:

1. **POST `/shorten`**:

   - **Description**: Creates a shortened URL from an original URL.
   - **Request**:
     ```json
     {
       "url": "https://example.com"
     }
     ```
   - **Response**:
     ```json
     {
       "short_url": "abc12345"
     }
     ```
   - **Example**:
     ```bash
     curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "https://example.com"}'
     ```

2. **GET `/{shortCode}`**:
   - **Description**: Redirects to the original URL associated with the `shortCode`.
   - **Response**: HTTP 301 redirect to the original URL.
   - **Example**:
     ```bash
     curl -v http://localhost:8080/abc12345
     ```
     Expected response includes `Location: https://example.com` with status `301 Moved Permanently`.

You can also use the provided `url-shortener-api.http` file with the VS Code REST Client extension to test these endpoints.

## Testing

### Unit Tests

The project includes unit tests for the Redis repository and URL service, located in the `tests/` directory.

1. **Run Unit Tests**:

   ```bash
   go test ./tests -v
   ```

   This runs all tests in `tests/redis_repository_test.go` and `tests/url_service_test.go`, using `redismock` to mock Redis interactions.

2. **API Testing**:
   Use the `url-shortener-api.http` file to test the API endpoints:
   - Open the file in VS Code with the REST Client extension.
   - Click "Send Request" for each endpoint to execute the HTTP requests.
   - Alternatively, use `curl` or tools like Postman.

### Test Scenarios

- **Successful URL Shortening**: Send a valid URL to `POST /shorten` and verify a `short_url` is returned.
- **Invalid URL**: Send an invalid URL (e.g., `{"url": "invalid"}`) to `POST /shorten` and expect a `400 Bad Request` response.
- **Redirect Success**: Use the `shortCode` from a successful POST to test `GET /{shortCode}` and verify the redirect.
- **Nonexistent Short Code**: Send a `GET /nonexistent` request and expect a `404 Not Found` response.

## Clean Architecture

The project follows Clean Architecture principles:

- **Domain Layer** (`internal/domain/url`): Contains the core business logic, including the `URL` entity, repository interface, and service logic.
- **Use Cases** (`internal/usecases`): Encapsulates application-specific business rules, delegating to the domain service.
- **Infrastructure** (`internal/infrastructure`): Implements external systems (Redis for storage, HTTP handlers for the API).
- **Cmd** (`cmd/api`): Entry point for the application, wiring dependencies together.

This separation ensures maintainability and testability, with clear boundaries between layers.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Make your changes and commit (`git commit -m "Add your feature"`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a Pull Request.

Please ensure all tests pass and follow the project's coding standards.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
