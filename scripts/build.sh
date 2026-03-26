#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
BACKEND_DIR="$ROOT_DIR/backend"
FRONTEND_DIR="$ROOT_DIR/frontend"
BIN_DIR="$ROOT_DIR/bin"

mkdir -p "$BIN_DIR"

echo "[build] backend: go test ./..."
(
  cd "$BACKEND_DIR"
  go test ./...
)

echo "[build] backend: go build -> $BIN_DIR/diputat-backend"
(
  cd "$BACKEND_DIR"
  go build -o "$BIN_DIR/diputat-backend" ./cmd/server
)

echo "[build] frontend: npm install"
(
  cd "$FRONTEND_DIR"
  npm install
)

echo "[build] frontend: npm run build"
(
  cd "$FRONTEND_DIR"
  npm run build
)

echo "[build] done"
