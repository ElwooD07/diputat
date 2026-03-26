# Samples

Comprehensive sample datasets representing Ukrainian political fact-checking for MVP development and testing.

## Dataset Overview

### Officials (10 deputies/officials)
Currently includes:
- Volodymyr Hroysman - Former PM, Verkhovna Rada deputy
- Yaroslav Mudry - Digital transformation specialist
- Natalia Sheremetyeva - Healthcare policy focus
- Oleksiy Honcharenko - Pro-European, NATO advocate
- Maryana Bezuhla - Anti-corruption activist
- Andriy Zahorodnyuk - Defense policy advisor
- Inna Sovsun - Environment & education advocate
- Dmytro Lytvyn - Labor rights advocate
- Yulia Navalnaya - Human rights defender
- Mykola Tyshchenko - Mayor of Kharkiv

**Purpose**: Represents diverse Ukrainian political landscape (multiple parties, roles, positions)

### Statements (20 tracked claims)
**Topics**: Budget, economy, healthcare, environment, defense, NATO, corruption, infrastructure, labor

**Fields**:
- Content: Direct quote or paraphrase of claim
- Source: Type (interview, parliament speech, press conference), URL, date, media outlet
- Status: pending, verified, new
- Extracted_by: scraper_ai (AI-assisted extraction from 1+1.ua and similar sources)
- Topics: Multiple classifications (budget, economy, etc.)
- Sentiment: neutral, positive, negative
- AI_confidence: 0.38-0.96 (AI model confidence in extraction accuracy)

**Sample statement**: "The IT sector contributes 12% of Ukraine's GDP" (verified with confidence 0.92)

### Verifications (8 detailed analyses)
**Verdict types**:
- **True** (2): "The IT sector does contribute ~12% to GDP" (confidence 0.95)
- **Partially** (2): "Healthcare spending increased, but 200% is nominal inflation; real increase ~48%" (confidence 0.81)
- **Misleading** (1): "NATO by 2027 is unrealistic timeline per NATO process" (confidence 0.88) 
- **Unverifiable** (1): "Corruption index claim premature; TI publishes annual data May 2026" (confidence 0.62)

**Verification method**: Hybrid human-AI (AI suggests verdict, human expert confirms with evidence)

**Evidence types**: Official data, government reports, international organization statements, satellite imagery, statistics, legal documents

**Timeline tracking**: Each verification records when statement was captured, when analysis occurred, when verdict was finalized

## Usage

### For Local Development
Files are served by Go backend from `data/samples/` directory:
```bash
GET /api/v1/officials → returns officials.json array
GET /api/v1/statements?official_id=official-1 → filters statements
GET /api/v1/timeline?official_id=official-1 → derives timeline from statements + verifications
```

### For Frontend Testing
Dashboard displays:
- 10 officials in selector dropdown
- 20 statements organized by official
- 8 verifications with verdict badges and evidence panels
- Timeline view showing statement and verification events

### For Data Model Validation
Ensures end-to-end schema consistency:
- `officials.json` structure matches Go `Official` struct and TypeScript `Official` interface
- `statements.json` includes all required fields for both backend and frontend
- `verifications.json` verdict field uses controlled vocabulary (true, partially, false, misleading, unverifiable)

## Next Phase: Real Data

To connect to actual Ukrainian news sources:

1. **Scraper** (`internal/scraper/`): Extract new statements from 1+1.ua, TSN, Suspilne
2. **AI Analyzer** (`internal/verifier/ai_analyzer.go`): Auto-generate preliminary verdicts
3. **Manual Curation**: Human review and confidence adjustment
4. **Database Migration**: Move from JSON to MongoDB with `data/migrations/`

This sample data serves as reference format and development baseline.

