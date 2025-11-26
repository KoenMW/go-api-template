# group11-go-microservice-template

Minimal Go microservice template following a hexagonal architecture. The service exposes a single POST /hello endpoint.

## Architecture

The project follows a hexagonal (ports & adapters) architecture:

- Core (domain + use cases) contains business logic and interfaces (ports).
- Adapters implement inbound (HTTP) and outbound integrations.

This separation keeps business rules independent from transport and infrastructure concerns.

## Endpoint

POST /hello

- Description: returns a greeting for the provided name.
- Request (application/json):
  - body: { "message": string }
- Response (application/json):
  - body: { "greeting": string }

Example request:

```json
{ "message": "Hello" }
```

cURL example:

```bash
curl -X POST http://localhost:8080/hello \
  -H "Content-Type: application/json" \
  -d '{"message":"Hello"}'
```

## Prerequisites

- Go 1.25+ (or the version used by your project)
- Docker & Docker Compose (if using Docker)

## Run locally (with go)

1. Build & run directly:
   ```bash
   go run ./src/app
   ```

## Run with Docker

1. Build & run with Docker Compose:
   ```bash
   docker compose up --build
   ```

The service will be available at http://localhost:8080/hello
