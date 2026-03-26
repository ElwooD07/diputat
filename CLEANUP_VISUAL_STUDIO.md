# Clean Visual Studio Legacy Files

This script removes old Visual Studio C++ project files that are no longer needed.

## Files to Remove

The following files are from an old Visual Studio C++ project template and are NOT used by the Diputat platform:

- `diputat.sln` - Visual Studio solution file
- `diputat.vcxproj` - Visual Studio C++ project file
- `diputat.vcxproj.filters` - Visual Studio filters
- `diputat.vcxproj.user` - Visual Studio user settings
- `diputat.cpp` - Empty C++ template file (already removed)

## Why Remove?

1. **Not used**: The Diputat platform uses Go + React + Docker
2. **Confusing**: These files suggest C++ is part of the project
3. **Clutter**: Unnecessary files in repository

## How to Remove

### Option 1: PowerShell (Windows)
```powershell
Remove-Item -Path "diputat.sln" -ErrorAction SilentlyContinue
Remove-Item -Path "diputat.vcxproj" -ErrorAction SilentlyContinue
Remove-Item -Path "diputat.vcxproj.filters" -ErrorAction SilentlyContinue
Remove-Item -Path "diputat.vcxproj.user" -ErrorAction SilentlyContinue
```

### Option 2: Git Bash / WSL
```bash
rm -f diputat.sln diputat.vcxproj diputat.vcxproj.filters diputat.vcxproj.user
```

### Option 3: Manual
Close Visual Studio and delete these files manually:
- `diputat.sln`
- `diputat.vcxproj`
- `diputat.vcxproj.filters`
- `diputat.vcxproj.user`

## After Removal

The Diputat project will continue to work normally using:
- Go backend (use `go run ./cmd/server` or `make backend-dev`)
- React frontend (use `npm run dev` or `make frontend-dev`)
- Docker Compose (use `docker-compose up --build` or `make start`)

## Confirmation

✅ VS Code configuration is ready in `.vscode/`  
✅ EditorConfig is configured in `.editorconfig`  
✅ All README files use English  
✅ No dependencies on Visual Studio

You can safely migrate to VS Code!
