# Data

Document-oriented project data for the Diputat MVP.

## Structure

```
data/
├── schemas/           # JSON schemas for the document model
├── samples/           # Local datasets used by the Go backend MVP
├── migrations/        # Migration examples and future migration assets
└── seeds/             # Optional MongoDB seed files
```

## Active MVP Documents

### officials
```json
{
  "_id": "official-1",
  "name": "Full Name",
  "position": "Position",
  "party": "Party",
  "metadata": {
    "created_at": "2026-03-26T09:00:00Z",
    "updated_at": "2026-03-26T09:00:00Z",
    "source": "seed"
  }
}
```

### statements
```json
{
  "_id": "statement-1",
  "official_id": "official-1",
  "content": "Statement text",
  "source": {
    "type": "interview",
    "title": "Source title",
    "date": "2026-03-26T09:00:00Z"
  },
  "status": "new",
  "topics": ["economy"],
  "metadata": {
    "created_at": "2026-03-26T09:10:00Z",
    "updated_at": "2026-03-26T09:10:00Z",
    "language": "en"
  }
}
```

### verifications
```json
{
  "_id": "verification-1",
  "statement_id": "statement-1",
  "result": "partially",
  "confidence": 0.82,
  "verdict": "Short human-readable verdict.",
  "evidence": [],
  "timeline": {
    "checked_at": "2026-03-26T10:00:00Z",
    "verified_by": "manual"
  },
  "metadata": {
    "created_at": "2026-03-26T10:00:00Z",
    "updated_at": "2026-03-26T10:00:00Z"
  }
}
```

## Usage

- The backend reads `data/samples/*.json` by default.
- `officials.json`, `statements.json`, and `verifications.json` form the current runnable dataset.
- `timeline` is a derived API response built from statements and verifications.

## Notes

The repository remains JSON-first for the MVP. MongoDB support is still a later phase.
