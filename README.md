# group11-go-api-template

Minimal Go api template following a hexagonal architecture. The service exposes user management endpoints.

## Architecture

The project follows a hexagonal (ports & adapters) architecture:

- Core (domain + use cases) contains business logic and interfaces (ports).
- Adapters implement inbound (HTTP) and outbound integrations.

This separation keeps business rules independent from transport and infrastructure concerns.

## Endpoints

POST /users

- Description: create a new user.
- Request (application/json):
  - body: { "name": string, "email": string }
- Response (application/json):
  - body: { "id": string, "name": string, "email": string }

GET /users

- Description: retrieve all users with pagination.
- Query parameters:
  - page: page number (default: 1)
  - per_page: number of users per page (default: 10)
- Response (application/json):
  - body: array of { "id": string, "name": string, "email": string }

GET /users/{id}

- Description: retrieve a user by UUID.
- Response (application/json):
  - body: { "id": string, "name": string, "email": string }

PUT /users/{id}

- Description: update an existing user.
- Request (application/json):
  - body: { "name": string, "email": string }
- Response (application/json):
  - body: { "id": string, "name": string, "email": string }

DELETE /users/{id}

- Description: delete a user by UUID.
- Response (application/json):
  - body: { "id": string }

cURL examples:

Create user:

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com"}'
```

List users:

```bash
curl http://localhost:8080/users?page=1&per_page=10
```

Get user:

```bash
curl http://localhost:8080/users/<user-uuid>
```

Update user:

```bash
curl -X PUT http://localhost:8080/users/<user-uuid> \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Updated","email":"alice.updated@example.com"}'
```

Delete user:

```bash
curl -X DELETE http://localhost:8080/users/<user-uuid>
```

## ORM

This project includes an ORM (bun) for database access. Database models use bun struct tags and UUIDs (github.com/google/uuid). Database connection and migrations are handled in the adaptors/db package.

## Warning

This template does not include authentication, authorization, input validation beyond basic JSON decoding, encryption, rate limiting, or other security controls. Do NOT use this code as-is in production for systems that handle important or sensitive data. Treat this project as a starter template and add proper security measures before deploying.

## Prerequisites

- Go 1.25+ (or the version used by your project)
- Docker & Docker Compose (if using Docker)

## Run locally (with go)

1. Build & run directly:

   ```bash
   cd src
   go run ./app
   ```

## Run with Docker

1. Build & run with Docker Compose:

   ```bash
   docker compose up --build
   ```

The service will be available at <http://localhost:8080>
