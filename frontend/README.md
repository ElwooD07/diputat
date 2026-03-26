# Frontend (React)

The frontend is a Vite and React dashboard for browsing officials, statements, verification results, and timeline events.

## Structure

```
frontend/
├── src/
│   ├── api.ts           # Backend API client
│   ├── App.tsx          # Dashboard composition
│   ├── main.tsx         # React entry point
│   ├── models.ts        # Shared TypeScript interfaces
│   └── styles.css       # Application styling
├── index.html
├── package.json
├── tsconfig.app.json
├── tsconfig.node.json
└── vite.config.ts
```

## Current Views

- **Officials** - tracked public figures and positions
- **Statements** - claims, sources, topics, and statuses
- **Verifications** - verdicts with confidence levels
- **Timeline** - combined statement and verification events

## Run

```bash
npm install
npm run dev -- --host 0.0.0.0 --port 4200
```

The dashboard is available at `http://localhost:4200`.

## API Configuration

The frontend reads `VITE_API_URL` and defaults to `http://localhost:8080/api/v1`.

## Notes

Some legacy scaffold README files still exist under `src/app/` from the earlier placeholder structure. The implemented MVP uses the files in `src/` listed above.
