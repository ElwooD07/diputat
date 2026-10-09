# DIPUTAT: GLOBAL DEVELOPMENT MASTERPLAN

## 1. COMPILER & RUNTIME REFACTOR FACT-BASE

- **OS Environment:** Windows 10 x64 | `CGO_ENABLED=0` (Pure Go compiler forced)
- **Module Architecture:** Single Go module initialized. `go.mod` is physically located inside the `backend/` directory.
- **Source Component State:**
  - `backend/internal/storage/model.go` [STABLE] - Defines raw domain structures (`Author`, `Party`, `SourcePoint`, `Link`, `Statement`).
  - `backend/internal/storage/repository.go` [STABLE] - Handles pure disk-based storage mappings and SHA-256 cryptographic content addressing (`SaveStatement`, `SaveLink`, `GetAuthorByID`).
  - `backend/internal/storage/repository_test.go` [STABLE] - Verifies content-addressed token serialization locally via `go test -count=1 ./...`.

## 2. API PROTOCOL CONSTRAINTS (CONTRACT-FIRST)

- **Protocol:** Strict Flat REST API verified via OpenAPI 3.0 specification.
- **Data Transfer Architecture:** Pure decoupled layout. The backend serves flat node models directly from disk without performing expensive recursive graph-nesting loops. Node compilation and visual edge assembly are fully delegated to the lightweight React layer.
- **Cryptographic Continuity:** IDs are passed strictly as string hashes (`stmt_<sha256>`, `link_<sha256>`), allowing polymorphic runtime routing on both ends of the generator pipeline.

## 3. ACTIVE CODEGENERATION & INGESTION BACKLOG

- Task 3.1: Generate type-safe Go server stubs from the OpenAPI spec using `oapi-codegen`.
- Task 3.2: Generate type-safe TypeScript interfaces from the OpenAPI spec for the React workspace.
- Task 3.3: Implement the raw JSON stream mapping structures inside `scripts/downloader.go` to match Verkhovna Rada's actual log layout.
