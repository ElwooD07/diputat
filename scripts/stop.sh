#!/usr/bin/env bash
set -euo pipefail

BACKEND_PORT="${BACKEND_PORT:-8080}"
FRONTEND_PORT="${FRONTEND_PORT:-4200}"

kill_by_port() {
  local port="$1"
  local label="$2"
  local pids=""

  if command -v lsof >/dev/null 2>&1; then
    pids="$(lsof -t -iTCP:"$port" -sTCP:LISTEN 2>/dev/null || true)"
  elif command -v ss >/dev/null 2>&1; then
    pids="$(ss -ltnp "( sport = :$port )" 2>/dev/null | awk -F'pid=' 'NR>1 {print $2}' | awk -F',' '{print $1}' | sort -u || true)"
  else
    echo "[stop] neither lsof nor ss found; cannot inspect port $port"
    return 1
  fi

  if [[ -z "$pids" ]]; then
    echo "[stop] $label not running on port $port"
    return 0
  fi

  echo "[stop] stopping $label on port $port (pid: $pids)"
  kill $pids 2>/dev/null || true
}

kill_by_port "$BACKEND_PORT" "backend"
kill_by_port "$FRONTEND_PORT" "frontend"

echo "[stop] done"
