# Quick Start Guide

Quick start for the current Diputat MVP.

## Prerequisites

- Docker and Docker Compose
- Go 1.21+ for local backend development
- Node.js 18+ and npm 9+ for local frontend development
- Make, optional but convenient

## Quick Start with Docker

```bash
# 1. Clone repository
git clone <repo-url>
cd diputat

# 2. Start all services
make start
# or
docker-compose up -d --build

# 3. Open in browser
# Frontend: http://localhost:4200
# Backend: http://localhost:8080
# Mongo Express: http://localhost:8081 (admin/admin)
```

## Local Development

### Backend

```bash
cd backend
go mod download
go run ./cmd/server
```

### Frontend

```bash
cd frontend
npm install
npm run dev -- --host 0.0.0.0 --port 4200
```

### Optional MongoDB

```bash
docker-compose up -d mongodb mongo-express
```

## Useful Commands

```bash
make help
make backend-dev
make frontend-dev
make test-backend
make test-frontend
make build-backend
make build-frontend
make clean
```

## Sample Data

The backend reads starter documents from `data/samples/` by default:

- `officials.json`
- `statements.json`
- `verifications.json`

## Next Steps

1. Read `README.md` for the current stack and MVP boundaries.
2. Review `IMPLEMENTATION_PLAN.md` for the phased roadmap.
3. Extend the backend contracts before adding new frontend flows.
