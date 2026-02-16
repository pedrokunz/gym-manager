# Gym Manager

A monorepo for a Gym Management system with a Go backend and Vue 3 frontend.
## Project Structure


- `backend/`: Go monolith application.
- `frontend/`: Vue 3 / Vite application.
- `.github/`: CI/CD workflows.

## Features

### Backend
- **Hybrid API**: Support for both REST (Members, Plans) and GraphQL (Classes) endpoints.
- **Persistence**: Integrated SQLite database with automatic schema initialization.
- **Data Consistency**: Automated data seeding and state reset on startup.
- **Middleware**: Built-in CORS support and basic error handling.

### Frontend
- **Member Management**: List, add, and remove gym members with real-time validation.
- **Membership Plans**: Interactive viewing of plans with automatic discount logic.
- **Class Schedules**: Dynamic fetching of gym classes using GraphQL.
- **Navigation**: smooth routing between management modules.
- **State Management**: Localized state handling for high-performance interactions.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (1.21+)
- [Node.js](https://nodejs.org/) (18+)
- [npm](https://www.npmjs.com/get-npm)

### Installation

1. Clone the repository.
2. Install dependencies and test browsers:
   ```bash
   npm run install:all
   ```

### Running the Application

#### Backend

```bash
npm run dev:backend
```

The API serves at `http://localhost:8080`.

#### Frontend

```bash
npm run dev:frontend
```

The application will be available at `http://localhost:5173`.

## CI/CD Infrastructure

Continuous Integration is handled via GitHub Actions. Workflows are defined in `.github/workflows/ci.yml`.
