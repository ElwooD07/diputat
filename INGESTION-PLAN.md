# Diputat Ingestion Plan (MVP -> 6-Month Growth)

## Goal
Build a simple, auditable, scalable pipeline that automatically collects public political statements and government actions, then prepares them for human review and public analysis.

Core principles:
- Deterministic first, AI second.
- Every record must be traceable to original evidence.
- Keep schema strict and explicit.
- Prefer simple systems that survive 6+ months of iteration.

---

## 1) Scope of Data Collection

### 1.1 Source Classes
1. **Government official sources**
   - Parliament websites, vote records, committee pages, legal portals.
2. **Public statement channels**
   - YouTube, Telegram, X, Facebook, interviews, TV transcripts.
3. **Independent reference sources**
   - Reputable media and watchdog publications used for verification.

### 1.2 Source Registry
Maintain a `sources` collection/table (not embedded in statements).

Minimum fields:
- `id`
- `url`
- `source_class` (`Parliament`, `PresidentialOffice`, `Cabinet`, `Ministry`, `Party`, `Media`, `Social`, `Other`)
- `platform` (`Parliament`, `YouTube`, `Telegram`, `X`, `Facebook`, `TV`, `Interview`, `Other`)
- `type` (`video`, `post`, `article`, `transcript`, `archive`, `official_record`)
- `publisher`
- `reliability` (`low`, `medium`, `high`)
- `created_at`
- `updated_at`

Notes:
- `published_at` belongs to statement/event context; keep source metadata minimal unless needed for specific source lifecycle handling.
- `PresidentialOffice` (ОП) is a first-class `source_class`, not a media subtype.

---

## 2) Target Domain Model (Unified)

Use one domain model for backend + API contracts (no temporary frontend-only model branch).

## 2.1 Officials
Stable entity for person profile and role history.

## 2.1.1 Institutions
Add an explicit institution layer for power centers.

Minimum fields:
- `id`
- `name` (e.g. `Office of the President of Ukraine`)
- `short_name` (e.g. `OP`)
- `type` (`Parliament`, `PresidentialOffice`, `Cabinet`, `Ministry`, `Party`, `Committee`, `Court`, `LocalGov`, `Other`)
- `country` (default `UA`)
- `active`

## 2.2 Statements
Atomic unit of claim/expression with strict fields.

Minimum fields:
- `id`
- `official_id`
- `institution_id` (required when statement is made on behalf of institution/account)
- `text`
- `status` (`InProgress`, `Done`, `Blocked`, `Toxic`)
- `platform` (enum from source classes)
- `language` (`ua`, `ru`, `en`)
- `statement_kind` (`promise`, `position`, `justification`, `accusation`, `denial`, `report`, `call_to_action`, `other`)
- `policy_domain` (`defense`, `economy`, `tax`, `energy`, `healthcare`, `justice`, `anti_corruption`, `education`, `other`)
- `published_at`
- `source_ids` (`string[]`, weak references)
- `reaction_ids` (`string[]`, optional weak references)
- `related_statement_ids` (`string[]`, optional weak references)
- `created_at`
- `updated_at`

## 2.3 Reactions
Separate, reusable collection for public/party/institution responses.

Minimum fields:
- `id`
- `target_type` (`statement`, `official`, `topic`)
- `target_id`
- `actor`
- `stance` (`support`, `reject`, `neutral`, `mixed`)
- `source_ids`
- `created_at`

## 2.4 Verifications
Keep evidence-centric verification data separate from statements.

## 2.5 Statement Relations
Store volatility/history explicitly.

Minimum fields:
- `id`
- `from_statement_id`
- `to_statement_id`
- `relation_type` (`supports`, `contradicts`, `retracts`, `clarifies`)
- `confidence` (`low`, `medium`, `high`)
- `created_at`

---

## 3) Ingestion Pipeline (Deterministic First)

Pipeline stages:
1. **Fetch**
   - Pull raw content from source URLs/APIs.
2. **Parse**
   - Extract structured text + metadata.
3. **Normalize**
   - Map to strict enums and canonical field names.
4. **Deduplicate**
   - Exact hash and near-duplicate checks.
5. **Store**
   - Save raw artifact + normalized entities.
6. **Queue for review**
   - Route uncertain/high-risk items to human validation.

Each stage writes logs and status records.

---

## 4) AI Usage Strategy (Small AIs)

Use AI only where it adds clear value:
- Claim extraction from long text/transcript.
- Topic classification.
- Contradiction candidate detection.
- Optional language hinting.

Rules:
- AI output is a **proposal**, not final truth.
- Persist `model`, `prompt_version`, `confidence`.
- Require review when confidence is below threshold or item is high impact.

---

## 5) Trust, Auditability, and Safety

Every statement must have:
- original source link(s),
- fetch timestamp,
- raw artifact hash,
- normalization trace,
- review decision trail.

Safety controls:
- source allowlist/registry,
- request rate limits,
- archive snapshots,
- tamper-evident ingestion logs,
- explicit handling for manipulated media risk.

---

## 6) Implementation Blueprint in This Repo

## 6.1 Backend Structure (proposed extension)
- `backend/internal/ingest/`
  - `fetch/`
  - `parse/`
  - `normalize/`
  - `dedupe/`
  - `review/`
- `backend/internal/models/`
  - extend unified `Statement`, add `Source`, add `Reaction`.
- `backend/internal/database/`
  - repository methods for sources/reactions.
- `backend/cmd/ingest/`
  - CLI entry for ingestion jobs.

## 6.2 Storage Approach
For MVP:
- Continue JSON collections, add:
  - `data/samples/sources.json`
  - `data/samples/reactions.json`
- Keep read/write atomicity and schema validation in repository layer.

---

## 7) Validation Rules (Hard Constraints)

1. `language` must be exactly one of: `ua`, `ru`, `en`.
2. `status` must be one of workflow enums.
3. `platform` must be one of allowed enum values.
4. `source_class` must be one of allowed enum values.
5. `official_id` must reference existing official (or be rejected/flagged).
6. `institution_id`, if present, must reference existing institution.
7. `statement_kind` and `policy_domain` must be valid enum values.
8. `related_statement_ids` are weak refs with warnings if missing.
9. `source_ids` and `reaction_ids` are weak refs:
   - missing refs are allowed with warning + audit entry, not silent drop.
10. Records missing mandatory fields are rejected at ingest time.

---

## 7.1 Keep It Stupid UX Rules (Citizen-First)

Every record shown to a user must answer 3 questions only:
1. **Who said/did it?**
2. **What exactly was said/done?**
3. **Where is proof?**

UI-facing output should always include:
- display name (`official` or `institution`),
- 1-line statement text,
- status badge,
- source link button,
- simple context labels: platform + date + domain.

Avoid:
- internal jargon,
- hidden confidence-only decisions,
- unexplained score jumps.

---

## 8) Delivery Phases

## Phase A — Foundation
- Finalize unified schema for Statement/Source/Reaction.
- Add strict validation tests.
- Add OpenAPI update for new fields/entities.

## Phase B — Deterministic Ingestion
- Implement fetch + parse + normalize + dedupe pipeline.
- Persist raw + normalized records.
- Add ingestion run logs and failure recovery.

## Phase C — Human Review Workflow
- Add review queue model and API endpoints.
- Add decision states (`pending_review`, `approved`, `rejected`, `needs_more_evidence`).

## Phase D — AI Assist
- Add optional AI enrichment workers with confidence gates.
- Store model metadata and prompt versions.

## Phase E — Public Transparency Layer
- Expose explainable score components and change history.
- Show evidence-backed links for each displayed conclusion.

---

## 9) Definition of Done (for each ingestion feature)

A feature is complete only when:
1. It has strict schema validation.
2. It is covered by tests (happy path + invalid input).
3. It writes auditable records.
4. It fails explicitly with actionable error messages.
5. It is reflected in OpenAPI and repo docs.

---

## 10) Next Step

Decompose this plan into implementation tasks with dependencies:
- schema tasks,
- repository tasks,
- ingestion worker tasks,
- API tasks,
- tests and verification tasks.
