# Core Module

Core services and functionality used throughout the application.

## Structure

```
core/
├── services/       # Services (HTTP, State, Auth)
├── interceptors/   # HTTP interceptors
├── guards/         # Route guards
└── models/         # Base interfaces
```

## Services

- `data.service.ts` - Data operations
- `state.service.ts` - State management
- `logger.service.ts` - Logging

## Principles

Core module is imported once in AppModule.
