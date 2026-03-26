# Environments

Configurations for different environments.

## Files

- `environment.ts` - Development (local development)
- `environment.prod.ts` - Production (if needed)

## Example

```typescript
export const environment = {
  production: false,
  apiUrl: 'http://localhost:8080',
  enableDebug: true
};
```

## Usage

```typescript
import { environment } from '../environments/environment';

console.log(environment.apiUrl); // http://localhost:8080
```

## Security

Do not store secret keys in environment files!
