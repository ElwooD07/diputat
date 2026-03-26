# Implementation Plan

This document defines the lowest-risk implementation path for Diputat.

## Decision Summary

The project should follow this sequence:

1. Replace the planned Angular frontend with React before frontend implementation begins.
2. Build a stable backend and data contract before AI-assisted workflows.
3. Start with deterministic ingestion and verification flows.
4. Introduce AI only as an internal draft-generation tool with human review.

## Why React Instead of Angular

React is the lower-risk choice at the current stage of the repository because:

1. The frontend is not implemented yet, so migration cost is close to zero.
2. The product is closer to an exploratory dashboard than a large enterprise form-heavy application.
3. The interface will likely change quickly as the data model and verification workflow mature.
4. Phase 6 will require flexible review and evidence-analysis UX, which is easier to iterate in React.

Angular should only remain the default choice if the team already has strong Angular expertise and wants stricter framework conventions from the start.

## Delivery Principles

1. Build the system around explicit contracts, not around AI components.
2. Keep storage and service boundaries stable before adding automation.
3. Prefer human-reviewed workflows over fully autonomous pipelines.
4. Add MongoDB after the application works with local deterministic data.
5. Treat AI output as a draft artifact, never as a trusted final result.

## Phase Plan

### Phase 1: Foundation and Cleanup

Goal: make the repository internally consistent and ready for implementation.

Deliverables:

1. Align documentation with the chosen React frontend direction.
2. Remove or update references that assume Angular-specific tooling.
3. Standardize project language to English where standards require it.
4. Confirm the MVP scope for officials, statements, verification, and timeline.
5. Define API and data ownership boundaries between backend, frontend, and data.

Exit criteria:

1. Repository documentation reflects the real intended stack.
2. MVP entities and flows are defined.
3. There is one agreed execution plan for backend, frontend, and data.

### Phase 2: Backend MVP

Goal: create a runnable backend with a stable data contract.

Deliverables:

1. Go module and project bootstrap.
2. HTTP server entry point.
3. Configuration loading for local development.
4. Domain models for official, statement, and verification.
5. Validation rules derived from the existing JSON schemas.
6. Repository interfaces.
7. JSON-file repository implementation for deterministic local use.
8. Read-focused API endpoints:
   - health
   - officials list/detail
   - statements list/detail
   - verifications list/detail
   - timeline query

Exit criteria:

1. Backend runs locally without MongoDB.
2. Sample data can be loaded and returned through the API.
3. Tests cover core models, handlers, and repositories.

### Phase 3: React Frontend MVP

Goal: build the first usable research interface against the stable backend.

Recommended stack:

1. React with TypeScript.
2. Vite for build tooling.
3. React Router for navigation.
4. TanStack Query for server-state management.
5. A simple component library only if it does not constrain the domain UI.

Deliverables:

1. App shell and routing.
2. Officials view.
3. Statements view with filtering.
4. Verification results view.
5. Timeline view.
6. Shared API client and typed models.
7. Error, loading, and empty states.

Exit criteria:

1. The UI consumes live backend data.
2. Core entity views are usable on desktop and mobile.
3. Frontend tests cover key data loading and rendering paths.

### Phase 4: MongoDB Integration

Goal: add production-oriented persistence without destabilizing the app.

Deliverables:

1. Mongo-backed repository implementations.
2. Repository selection through configuration.
3. Seed loading strategy.
4. Migration or initialization strategy for required collections and indexes.
5. Docker and Make targets that actually work against the implemented app.

Exit criteria:

1. The application works with both JSON and Mongo-backed storage.
2. Local Docker startup is reliable.
3. Seed data loads consistently.

### Phase 5: Verification Workflow

Goal: make verification a real deterministic product workflow.

Deliverables:

1. Verification service layer.
2. Evidence data model and storage.
3. Timeline construction logic.
4. Contradiction detection between statements.
5. Manual verification creation and update flows.
6. Reviewable status transitions for statements and verifications.

Exit criteria:

1. Statements can be linked to evidence and verification outcomes.
2. Timeline and contradiction features are functional without AI.
3. Review workflows are explicit and testable.

### Phase 6: Safe AI-Assisted Automation

Goal: introduce automation without making the system opaque or unsafe.

This phase should start only after Phases 2 through 5 are stable.

Recommended scope:

1. Collector service
   Fetches and stores raw source material from selected inputs.
2. Extractor service
   Produces candidate statements from raw content using AI or NLP.
3. Review queue
   Allows human approval, rejection, or editing of extracted statements.
4. Draft verifier support
   Suggests evidence and verification drafts, but does not auto-publish verdicts.

Rules for implementation:

1. Store raw source content separately from extracted claims.
2. Persist provenance, model metadata, and confidence scores for each extraction.
3. Never let AI-generated statements bypass human review.
4. Never let AI-generated verification results become final without an explicit approval workflow.
5. Keep AI components behind internal service interfaces so they can be replaced later.

Exit criteria:

1. AI produces draft statements linked to original source content.
2. Reviewers can accept, edit, or reject drafts.
3. The system logs provenance and confidence for every AI-generated artifact.

## Immediate Next Steps

The next implementation steps should be:

1. Update repository documentation from Angular to React.
2. Bootstrap the Go backend MVP.
3. Define the REST contract for officials, statements, verifications, and timeline.
4. Add sample JSON data that matches the schemas.
5. Scaffold the React frontend once backend contracts are stable.

## What Not To Do Yet

1. Do not implement autonomous scraping first.
2. Do not make MongoDB mandatory for the first working build.
3. Do not let the frontend read schema files directly as its primary data source.
4. Do not couple handlers directly to AI providers.
5. Do not mix verification logic with transport or scraping code.