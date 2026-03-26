---
applyTo: "**/data/**/*.json,**/samples/officials.json"
description: "Use when editing officials.json or related data files — keep repository free of static personal officials roster and validate schema consistency."
---

# Officials Data Rules

Before modifying `officials.json` or related seed files, apply these checks.

## 0. No Static Personal Roster in Repository

- `data/samples/officials.json` must stay as an empty array (`[]`).
- Do not commit hardcoded personal names or per-person roster records to sample data files.
- If officials data is required, load it from the runtime database layer, not from repository JSON fixtures.

## 1. Canonical Seed Format

- `data/seeds/canonical_officials.json` may be empty.
- If used, entries must be machine keys only (`handle`, `position`) and must avoid personal name fields.
- Keep compatibility with `backend/internal/database/canonical.go` loader.

## 2. Common Mistakes to Avoid

- Do not reintroduce `name`, `bio`, `contacts`, or `photo_url` into `officials.json` sample fixtures.
- Do not add person-specific validation rules in code comments or tests.
- Keep sample data deterministic and non-personal.

## 3. Validation Checklist

- [ ] `data/samples/officials.json` is exactly `[]`
- [ ] `data/seeds/canonical_officials.json` contains no personal name fields
- [ ] `go test ./backend/internal/database/...` passes after change
