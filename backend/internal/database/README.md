# Database Layer

Repository abstractions and local data access for the backend MVP.

## Current Support

- JSON files in `data/samples/`
- MongoDB planned for a later phase

## Active Interface

The current `Repository` interface exposes read operations for:

- officials
- statements
- verifications

## Implementation

`json_repository.go` loads the starter datasets into memory and serves them to the API handlers.
