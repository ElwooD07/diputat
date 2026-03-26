# Diputat - Fact-Check Platform

Research project for monitoring, structuring, and verifying statements made by public officials.

## Recommended IDE

This project is optimized for VS Code.

## Project Structure

- **backend/** - Go API and local JSON-backed repository
- **frontend/** - React and Vite research dashboard
- **data/** - Schemas, sample documents, seeds, and migrations

## Working Principles

1. **Closed system** - local-first development and internal workflows
2. **Document-oriented data** - officials, statements, and verifications are stored as independent documents
3. **Deterministic MVP** - backend and frontend work without AI or MongoDB
4. **Review-first fact-checking** - evidence and verdicts are explicit data, not opaque automation

## Run Locally

```bash
# Backend
cd backend
go run ./cmd/server

# Frontend
cd frontend
npm install
npm run dev -- --host 0.0.0.0 --port 4200
```

The backend serves the API at `http://localhost:8080` and the frontend runs at `http://localhost:4200`.

## Requirements

- Go 1.21+
- Node.js 18+
- npm 9+
- Docker and Docker Compose (optional)
- MongoDB 7.0+ (optional for later phases)

## Current MVP

The repository now includes:

- A Go backend with `health`, `officials`, `statements`, `verifications`, and `timeline` endpoints

Requirements are tracked in [REQUIREMENTS.md](REQUIREMENTS.md).
- Local sample JSON datasets in `data/samples/`
- A React dashboard that consumes the backend API
- A migration example in `data/migrations/001_initial_setup.json`

## Project Standards

All code, documentation, and comments must be in English.

See [STANDARDS.md](STANDARDS.md) for conventions and [IMPLEMENTATION_PLAN.md](IMPLEMENTATION_PLAN.md) for the lowest-risk delivery sequence.

## Quick Links

- [Quick Start Guide](QUICKSTART.md) - Get started quickly
- [Implementation Plan](IMPLEMENTATION_PLAN.md) - Recommended delivery path
- [Contributing Guidelines](CONTRIBUTING.md) - How to contribute
- [Project Standards](STANDARDS.md) - Coding standards and conventions
- [Project Configuration](.diputatrc.json) - Project settings

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
