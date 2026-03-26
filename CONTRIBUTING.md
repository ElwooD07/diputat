# Contributing to Diputat

Thank you for contributing. This document describes the current development rules and expectations.

## Language Requirement

All contributions must be in English:
- Code comments
- Documentation
- Commit messages
- Variable and function names
- Pull request descriptions

See [STANDARDS.md](STANDARDS.md) for the full language and formatting conventions.

## Project Structure

The project currently has three active areas:
- **backend/** - Go API and local JSON-backed data access
- **frontend/** - React and Vite dashboard
- **data/** - Schemas, samples, seeds, and migrations

## Principles

1. **Closed system** - no public API exposure is required for the MVP
2. **Mental model** - NoSQL documents represent concepts, not rigid tables
3. **Deterministic first** - backend and frontend should work before AI automation is added
4. **Review-first fact-checking** - explicit evidence and verdicts matter more than automation speed

## Code Standards

### Go (Backend)

```go
// Use tabs for indentation.
// Keep comments in English.
// Exported functions need Go documentation comments.
```

Conventions:
- File names: `snake_case.go`
- Package names: `lowercase`
- Variable names: `camelCase`
- Constant names: `PascalCase` or `UPPER_CASE`
- Run `gofmt` and `go vet` before commit

### TypeScript and React (Frontend)

```typescript
export interface Statement {
  _id: string;
  content: string;
}
```

Conventions:
- File names: `kebab-case.ts` or `kebab-case.tsx`
- Component and interface names: `PascalCase`
- Variable and function names: `camelCase`
- Constants: `UPPER_CASE`
- Run `npm run typecheck` before commit

### JSON (Data)

```json
{
  "field_name": "value"
}
```

Conventions:
- Use 2 spaces for indentation
- Field names use `snake_case`
- Keep sample data aligned with the schemas and backend models

## Git

### Commit Messages

```text
<type>(<scope>): <subject>
```

Types:
- `feat`
- `fix`
- `docs`
- `style`
- `refactor`
- `test`
- `chore`

Examples:
```text
feat(api): add timeline endpoint
fix(frontend): handle empty verification list
docs(readme): update local run commands
```

### Branches

- `main` - main branch
- `feature/*` - new features
- `fix/*` - bug fixes
- `docs/*` - documentation changes

## Testing

### Backend

```bash
go test ./...
go test -cover ./...
```

### Frontend

```bash
npm test
npm run build
```

## Documentation

- Keep the root documentation aligned with the implemented stack.
- Update API notes when endpoints or payloads change.
- Preserve README coverage for active directories.

## Code Review

Before merge:
- Code follows standards
- Checks pass
- Documentation is updated
- Changes stay scoped to the task

## Questions

Create an issue or open a discussion in the repository.
