# DIPUTAT: COMPACT ARCHITECTURE MATRIX & INGESTION STATE

## A. SYSTEM CONFIGURATION BASELINE

- **Environment:** Windows 10 x64 | CGO_ENABLED=0 (Pure Go compiler forced)
- **Module Boundary:** Verified local module mapping (`go.mod` inside `backend/`).
- **Target Component Source:** `backend/internal/storage/repository.go` [STABLE]
- **Verification Vector:** `backend/internal/storage/repository_test.go` [STABLE]

## B. DATA INGESTION PROTOCOL (PHASE 1: HARVESTING)

- **Core Subject Vector:** European Solidarity (ЄС) party faction track.
- **Data Range:** Retrospective (2018) | Pre-War (Winter 2021/22) | Active-War (Spring-Autumn 2022).
- **Source Gateway:** Verkhovna Rada Open Data Hub (`data.rada.gov.ua/ogd/zal/stenogram/`).

## C. INFRASTRUCTURE & NETWORK LAWS

1. **Authentication Rule:** Access is anonymous using the implicit `OpenData` system token context. Bulk script triggers targeting `api/limits` are prohibited to prevent immediate IP fire wall blocks.
2. **Rate Limiting Guardrail:** Explicit download orchestration forced inside `downloader.go` — maximum 60 requests per minute with randomized interval offsets between 5 and 7 seconds.
3. **Data Loss Strategy:** Chronological continuity drops caused by wartime redacted blocks or closed plenary sessions must be treated as valid historical status markers, not engine errors.

## D. ACTIVE TESTING BACKLOG

- [ ] Implement the static structures inside `scripts/downloader.go` mapping to Rada's chronological log payload fields.
- [ ] Map out the multi-year speaker naming variations for strict dictionary resolution matching.
