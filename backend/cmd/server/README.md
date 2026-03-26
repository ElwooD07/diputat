# Server Entry Point

This directory contains the backend executable entry point.

## Current Responsibility

`main.go` loads runtime configuration, opens the local JSON-backed repository, and starts the HTTP API server.

## Run

```bash
go run ./cmd/server
```

## Notes

The current MVP starts only the API server. Scraping, automation, and verification jobs are later-phase work.
