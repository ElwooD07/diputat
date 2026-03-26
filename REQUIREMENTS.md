# Project Requirements

This file is the dedicated source for toolchain and runtime requirements.

## Core Requirements

- Go 1.21+
- Node.js 18+
- npm 9+

## Optional Requirements

- Docker and Docker Compose (for containerized development)
- MongoDB 7.0+ (for later project phases)

## Runtime Defaults

- Backend default port: `8080`
- Frontend default dev port: `4200`
- Backend health endpoint: `GET /health`

## Notes

- The backend reads local JSON data from `../data/samples` by default.
- You can override dataset location with `DATA_DIR`.
- You can override backend port with `PORT`.
