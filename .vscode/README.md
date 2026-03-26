# VSCode Configuration

This directory contains Visual Studio Code workspace settings and recommendations.

## Files

- `settings.json` - Workspace settings for code formatting, linting, and editor behavior
- `extensions.json` - Recommended extensions for development

## Settings Highlights

### Code Formatting
- **Format on save** enabled for all languages
- **Go**: Uses tabs (4 spaces width)
- **TypeScript/JavaScript/HTML/CSS**: Uses 2 spaces
- **JSON**: Uses 2 spaces
- **Markdown**: Uses 2 spaces, preserves trailing whitespace

### Language Servers
- **Go**: Uses official Go language server with `golangci-lint`
- **TypeScript**: Uses the built-in TypeScript language server
- **React**: Uses TypeScript and ESLint tooling for `.tsx` development

### Spell Checking
- **Language**: English only (enforces project standard)
- **Enabled for**: Go, TypeScript, JavaScript, Markdown, JSON
- **Custom words**: Project-specific terms (diputat, nosql, etc.)

## Recommended Extensions

### Required
- **Go** (`golang.go`) - Go language support
- **ESLint** (`dbaeumer.vscode-eslint`) - JavaScript and TypeScript linting
- **Prettier** (`esbenp.prettier-vscode`) - Code formatter
- **EditorConfig** (`editorconfig.editorconfig`) - Maintains coding styles

### Recommended
- **Code Spell Checker** (`streetsidesoftware.code-spell-checker`) - Spell checking
- **GitLens** (`eamodio.gitlens`) - Git integration
- **Docker** (`ms-azuretools.vscode-docker`) - Docker support
- **Markdown All in One** (`yzhang.markdown-all-in-one`) - Markdown tools

## Installation

VSCode will automatically prompt you to install recommended extensions when you open this workspace.

Or install manually:
1. Open Command Palette (`Ctrl+Shift+P` or `Cmd+Shift+P`)
2. Type: `Extensions: Show Recommended Extensions`
3. Click "Install" on each extension

## Notes

- These settings enforce project standards defined in [STANDARDS.md](../STANDARDS.md)
- Settings apply only to this workspace, not globally
- You can override settings in your user settings if needed
