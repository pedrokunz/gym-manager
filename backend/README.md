# Gym Manager Backend

Go backend for Gym Manager.

## Features

- **Hybrid API**: Mix of REST and GraphQL.
- **SQLite Database**: Persistent storage with automated table initialization.
- **Auto-Seeding**: Resets and populates data on every startup for consistent local dev.

## API Endpoints

### REST

- `GET /api/members`: List members (supports `?status=`, `limit=`, `offset=`).
- `GET /api/members/:id`: Get detailed member profile.
- `GET /api/members/:id/invoices`: List invoices for a specific member.
- `POST /api/members`: Add a new member.
- `DELETE /api/members/:id`: Remove a member.
- `GET /api/dashboard`: Get dashboard analytics and recent activity.
- `GET /api/plans/getall`: List all membership plans.
- `POST /api/plans_create`: Add a new plan.
- `GET /api/invoices`: List all invoices.
- `POST /api/invoices/pay/:id`: Mark an invoice as paid.

### GraphQL

- `POST /graphql`: Query gym classes.

## Quick Start


### Run the server

```bash
go run cmd/main.go
```

### Run tests

```bash
go test ./...
```

## Structure

- `cmd/`: Application entry point.
- `internal/`: Private application and library code.
- `tests/`: Integration and unit tests.
