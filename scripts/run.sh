#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
BACKEND_BIN="$ROOT_DIR/bin/diputat-backend"
FRONTEND_DIR="$ROOT_DIR/frontend"

if [[ ! -x "$BACKEND_BIN" ]]; then
  echo "[run] missing backend binary: $BACKEND_BIN"
  echo "[run] run ./scripts/build.sh first"
  exit 1
fi

cleanup() {
  local exit_code=$?
  if [[ -n "${BACKEND_PID:-}" ]]; then
    kill "$BACKEND_PID" 2>/dev/null || true
  fi
  if [[ -n "${FRONTEND_PID:-}" ]]; then
    kill "$FRONTEND_PID" 2>/dev/null || true
  fi
  wait 2>/dev/null || true
  exit "$exit_code"
}
trap cleanup INT TERM EXIT

echo "[run] starting backend on http://127.0.0.1:8080"
PORT=8080 DATA_DIR="$ROOT_DIR/data/samples" "$BACKEND_BIN" &
BACKEND_PID=$!

echo "[run] starting frontend preview on http://127.0.0.1:4200"
(
  cd "$FRONTEND_DIR"
  npm run preview -- --host 127.0.0.1 --port 4200
) &
FRONTEND_PID=$!

echo "[run] services are up"
echo "[run] frontend: http://127.0.0.1:4200"
echo "[run] backend:  http://127.0.0.1:8080"
echo "[run] health:   http://127.0.0.1:8080/health"

wait
