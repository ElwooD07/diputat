# Models

Backend domain models for the fact-checking workflow.

## Files

- `official.go` - public official records
- `statement.go` - tracked statements and source metadata
- `verification.go` - verdicts, evidence, and contradictions
- `timeline.go` - derived timeline response model
- `common.go` - shared metadata types

## Purpose

These models stay close to the JSON document structure so the backend, frontend, and sample data remain aligned.
