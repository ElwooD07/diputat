# Database Layer

Repository abstractions and local data access for the backend MVP.

## Current Support

- JSON files in `data/samples/`
- MongoDB planned for a later phase

## Active Interface

The current `Repository` interface exposes read and write operations for:

- officials
- statements
- verifications
- statement filtering by `official_id`
- official sorting by `verification_score`

## Implementation

`json_repository.go` provides a concurrent-safe local JSON storage engine:

- `sync.RWMutex` for high-read/controlled-write access patterns
- one JSON collection file per document type (`officials.json`, `statements.json`, `verifications.json`)
- atomic write-through persistence on every upsert
- backward-compatible normalization of legacy sample fields/statuses
