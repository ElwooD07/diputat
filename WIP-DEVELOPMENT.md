# Diputat - Project Status & State Log

_Last Updated: 2026-10-08_

## 1. System Context & Active Rules

- **Repository Layout:** Multi-module architecture. The Go module (`go.mod`) is located inside the `backend/` directory.
- **Build Constraints:** Windows 10 x64, CGO is disabled globally (`CGO_ENABLED=0`).
- **Commands:** Run `go build ./...` and `go test ./...` strictly from inside the `backend/` folder.
- **AI Stack:**
  - Remote Architect: Cloud Gemini (High-reasoning, architectural decisions).
  - Local Executor: Qwen 2.5 Coder 14B via VS Code Continue (Inline edits via `Ctrl+I`).

## 2. Verified File Architecture

### `backend/internal/storage/repository.go`

- Fully normalized graph architecture (Nodes: Authors, ExternalSources, Statements).
- Contains structs: `Author`, `ExternalSource`, `StatementSource`, `StatementLink`, `Statement`.
- Contains `FileStatementRepository` with `SaveStatement` and `saveToFile` helper.

### `backend/internal/storage/repository_test.go`

- Target test case: `TestGraphArchitectureIngestion`.
- Successfully writes and cross-links graph nodes into `data/samples/generated/`.
- **Status:** COMPILING & PASSING.

## 3. Immediate Implementation Backlog

- [ ] Task 1: Implement Graph Reader (`GetStatementByID` & `GetAuthorByID`).
- [ ] Task 2: Implement Graph Crawler (`GetStatementTimeline` recursively resolving `StatementLink`).
- [ ] Task 3: Implement Trust Aggregator (Calculating weights based on `ExternalSource.TrustRating`).

## 4. Execution Rules for Remote AI

1. Provide highly strict, explicit English prompts for the local model to execute using `Ctrl+I`.
2. Do not attempt to guess or fix build errors speculatively unless the user explicitly copies a terminal failure block.
3. Assume the codebase matches this file state exactly.

## 5. Established Architecture Decisions (Frozen & Approved)

### A. Graph Traversal & Infinite Loop Mitigation

- **Mechanism:** Choice B (Visited Nodes Map Tracker).
- **Implementation Rule:** The recursive timeline graph crawler must maintain a thread-safe `map[string]bool` tracker of visited `StatementID` entries during any depth traversal.
- **Human-Centric Rule:** If a cycle is detected, the traversal breaks to prevent a stack overflow, but the cycle edge must be explicitly logged and flagged in the output data structure so human investigators can audit the self-referential rhetorical loop.

### B. Ingestion Fault-Tolerance & Lazy Node Resolution

- **Mechanism:** Approach 2 (Dynamic Placeholder Provisioning).
- **Implementation Rule:** Entry-point ingestion must never drop or fail a statement due to missing foreign key profiles (`AuthorID` or `ExternalEvidenceIDs`).
- **Resolution Strategy:**
  1. If a statement references a non-existent `AuthorID`, the system automatically provisions a placeholder profile under that ID with `FullName: "Pending Verification / New Actor"` and records the statement.
  2. The data structure is treated as an organic, growing graph. Cleanups, biochemical deduplication, and manual biography updates are delegated downstream to human auditors. The pipeline remains unblocked.

### C. Production-Grade Ingestion Constraints (2026-10-09)

- **Data Retention Rule:** No text-length filters, character gates, or structural omissions are permitted during the download pass.
- **Upgradability Law:** All chronological sentences delivered by target MPs must be saved intact to preserve historical fidelity. Text cleaning, optimization, and analytical noise-filtering are entirely decoupled from Phase 1 and deferred downstream.
