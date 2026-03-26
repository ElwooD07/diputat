# Project Standards

This document defines the standards and conventions for the Diputat project.

## Language Standard

**All project content MUST be in English:**

### ✅ English Required For:
- All documentation (README files, guides, etc.)
- Code comments
- Commit messages
- Variable names, function names, class names
- API documentation
- Configuration files comments
- Error messages (user-facing can be localized)

### 📝 Example:

**✅ Good:**
```go
// ProcessStatement extracts and analyzes statement from media source
func ProcessStatement(statement Statement) error {
    // Validate input
    if statement.Content == "" {
        return errors.New("statement content is required")
    }
    // ...
}
```

**❌ Bad:**
```go
// Обробити заяву
func ProcessStatement(zayava Statement) error {
    // Перевірка вхідних даних
    if zayava.Content == "" {
        return errors.New("вміст заяви обов'язковий")
    }
    // ...
}
```

## Naming Conventions

### Go (Backend)
- **Files**: `snake_case.go` (e.g., `statement_processor.go`)
- **Packages**: `lowercase` (e.g., `scraper`, `verifier`)
- **Variables/Functions**: `camelCase` (e.g., `processStatement`, `officialID`)
- **Types/Structs**: `PascalCase` (e.g., `Statement`, `OfficialData`)
- **Constants**: `PascalCase` or `UPPER_CASE` (e.g., `MaxRetries`, `API_TIMEOUT`)

### TypeScript/React (Frontend)
- **Files**: `kebab-case.ts` or `kebab-case.tsx` (e.g., `statement-list.tsx`)
- **Components/Interfaces**: `PascalCase` (e.g., `StatementList`, `OfficialData`)
- **Variables/Functions**: `camelCase` (e.g., `processStatement`, `officialId`)
- **Constants**: `UPPER_CASE` (e.g., `API_URL`, `MAX_ITEMS`)

### JSON (Data)
- **Files**: `snake_case.json` (e.g., `officials.seed.json`)
- **Fields**: `snake_case` (e.g., `official_id`, `created_at`)

## Documentation Standards

### README Files
Every directory MUST have a `README.md` with:
1. **Title** - Clear description of the directory purpose
2. **Structure** - Directory structure overview (if applicable)
3. **Usage** - How to use/implement
4. **Examples** - Code examples (if applicable)

### Code Comments
- Use comments to explain **why**, not **what**
- Complex algorithms need detailed explanations
- Public functions/methods need documentation comments
- Keep comments concise and clear

## Commit Message Format

```
<type>(<scope>): <subject>

[optional body]

[optional footer]
```

### Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

### Examples:
```
feat(scraper): add support for Instagram posts
fix(verifier): correct confidence calculation
docs(readme): update installation instructions
refactor(database): simplify query builder
```

## Code Style

### Go
- Use `gofmt` for formatting
- Use `go vet` for static analysis
- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use tabs for indentation (defined in `.editorconfig`)

### TypeScript/React
- Use `npm run typecheck` or equivalent CI checks before merge
- Keep components small and data contracts explicit
- Use 2 spaces for indentation (defined in `.editorconfig`)
- Strict TypeScript mode enabled

## Testing Standards

### Backend (Go)
- Test files: `*_test.go`
- Table-driven tests preferred
- Aim for >70% coverage
- Mock external dependencies

### Frontend (React)
- Test files: `*.test.ts` or `*.test.tsx`
- Unit tests for components and utilities
- Integration tests for feature flows
- E2E tests for critical paths when the product surface expands

## Configuration Files

### Required Files:
- `.editorconfig` - Code style settings
- `.gitignore` - Git ignore rules
- `.diputatrc.json` - Project configuration

### Optional Files:
- `.env` - Environment variables (never commit!)
- `docker-compose.yml` - Local development setup
- `Makefile` - Convenience commands

## Enforcement

These standards are enforced through:
1. **Code Review** - All PRs must follow standards
2. **Linters** - Automated checks (`gofmt`, frontend type checks, and repository-specific validation)
3. **CI/CD** - Automated validation (when implemented)
4. **Documentation** - This file and `.diputatrc.json`

## Exceptions

Standards can be relaxed for:
- **User-facing content** - Can be localized to Ukrainian/other languages
- **External APIs** - Follow their naming conventions
- **Third-party libraries** - Use their conventions

## Updates

This document is versioned with the project. Updates require:
1. Team discussion
2. Update `.diputatrc.json`
3. Update this document
4. Notify all contributors

---

**Version**: 1.0.0  
**Last Updated**: 2026-03-26  
**Maintained By**: Diputat Team

For questions or suggestions, create an issue or contact the team.
