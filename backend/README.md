# Backend (Go)

The backend provides the local-first API for officials, statements, verifications, and derived timeline events.

## Structure

```
backend/
├── cmd/
│   └── server/          # HTTP entry point
├── internal/
│   ├── api/             # HTTP handlers and timeline builder
│   ├── config/          # Runtime configuration
│   ├── database/        # Repository interfaces and JSON-backed storage
│   └── models/          # Domain models for fact-checking data
└── go.mod
```

## Current API

### Endpoints
- `GET /health`
- `GET /api/v1/officials`
- `GET /api/v1/officials/{id}`
- `GET /api/v1/statements`
- `GET /api/v1/statements/{id}`
- `GET /api/v1/verifications`
- `GET /api/v1/verifications/{id}`
- `GET /api/v1/timeline`

### Data Source
- Local JSON files in `../data/samples` by default
- Optional `DATA_DIR` override for Docker or alternate datasets

## Run

```bash
go mod download
go run ./cmd/server
```

## Test

```bash
go test ./...
```

## Configuration

Environment variables:

- `PORT` - HTTP port, defaults to `8080`
- `DATA_DIR` - path to local JSON datasets, defaults to `../data/samples`

MongoDB, ingestion automation, and AI-assisted workflows remain later-phase work.
