# Migrations

Migration examples for future database-backed phases of the project.

## Files

Current migration assets:
- `001_initial_setup.json` - initial collection setup for `officials`, `statements`, and `verifications`

## Format

```json
{
  "version": "001",
  "description": "Create initial collections for officials, statements, and verifications.",
  "up": {
    "collections": ["officials", "statements", "verifications"]
  },
  "down": {
    "collections": []
  }
}
```

## Notes

- The current MVP is JSON-first and does not execute migrations yet.
- These files document the intended collection lifecycle for the later MongoDB phase.
