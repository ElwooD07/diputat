# Verifier (Fact Checker + Timeline Tracker)

**Single entity**: verification + timeline.

## Functionality

1. Statement fact-checking
2. Timeline building
3. Credibility assessment

## Triggers

- New statement in DB
- Cron (periodic checks)
- Manual launch

## Files

- `verifier.go` - Main logic
- `timeline.go` - Timeline building
- `evidence.go` - Evidence handling

## Importance

Factual presence of the component is critical for architecture.
