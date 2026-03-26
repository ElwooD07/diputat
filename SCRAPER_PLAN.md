# Scraper & Verification Implementation Plan

## Overview

Phase 2 of diputat extends beyond the JSON-backed MVP to build automated pipelines for scraping statements from Ukrainian news sources and verifying them against fact-checking databases.

---

## Phase 2a: Web Scraper (`internal/scraper/`)

### Target Media Sources

Primary sources (descending by reliability & coverage):
1. **1+1.ua** - Major news outlet, 1+1 TV coverage
2. **TSN.ua** - Ukrainian news agency of record
3. **UNIAN.ua** - Ukrainian Independent News Agency
4. **Suspilne.media** - Public broadcaster
5. **RBC Ukraine** - Business news
6. **Polityka.ua** - Political analysis

### Data Extraction Strategy

**Document Structure**:
```
Found headline: "Deputy Mudry says IT sector at 12% of GDP"
↓
Extract official name (NER: Named Entity Recognition)
↓
Extract statement (OCR or text parsing from image/video transcript)
↓
Identify source (origin URL, date published, media outlet)
↓
Link to sample official (fuzz match against `officials.json` names)
↓
Generate JSON statement with `extracted_by: "scraper_ai"`
```

**Tools & Libraries**:
- `colly` (Go web scraper) - HTTP client with built-in queue & rate limiting
- `goquery` - jQuery-style DOM traversal
- `mozilla/readability` - Article text extraction from HTML
- `jio` - Media source detection (1+1, TSN, etc. patterns)

**Implementation**:
```go
// backend/internal/scraper/scraper.go
type Scraper struct {
  sources []Source
  client  *colly.Collector
}

func (s *Scraper) ScrapeLatestStatements() ([]models.Statement, error) {
  // For each source:
  //   1. Fetch latest articles
  //   2. Parse article text
  //   3. Extract official names and claims
  //   4. Deduplicate against existing statements
  //   5. Return new statements for verification
}
```

### Scheduling

**Local MVP** (when ready):
```bash
# Run scraper on-demand
PORT=8080 SCRAPER_SOURCES="1plus1,tsn" go run ./cmd/server

# In future: cron or Kubernetes CronJob
0 */6 * * * /usr/bin/scraper --sources 1plus1,tsn --output /data/new-statements.json
```

---

## Phase 2b: NLP Statement Extraction

### Problem
Raw article text may contain multiple claims from multiple officials. We need to:
1. Identify official names (Named Entity Recognition)
2. Extract direct quotes or attributable statements
3. Assign confidence score to extraction

### Approach

**Option 1: Rule-Based Extraction** (MVP-ready, low cost)
```
Patterns:
- "Deputy [NAME] said: [QUOTE]"
- "[NAME] stated that [CLAIM]"
- "According to [NAME], [STATEMENT]"
- Regex on transcript: "^[NAME]:.*" (from debate/speech transcripts)
```

**Option 2: AI-Based (Future, higher confidence)**
- Fine-tune Hugging Face transformer on Ukrainian political speech
- Input: Article text
- Output: (official_name, statement, extraction_confidence)
- Use existing models:
  - `bert-base-multilingual` for NER
  - `xlm-roberta` for sentence understanding
  - Custom fine-tune on labeled diputat corpus

### Implementation Roadmap

**Week 1-2**: Rule-based extractor
- Regex patterns for common phrasing
- Test against sample 1+1.ua articles
- Target: 70% precision on quote attribution

**Week 3-4**: AI confidence scoring
- Upload annotated statements to Hugging Face
- Fine-tune transformer on Ukrainian political speech
- Target: 85% confidence on real statements

---

## Phase 2c: AI Verifier (`internal/verifier/`)

### High-Level Flow

```
New Statement (from scraper)
  ↓
AI Analyzer (natural language + knowledge base)
  ├─ Search for contradicting evidence
  ├─ Cross-reference with known facts
  ├─ Assign preliminary verdict (true/false/partial/misleading/unverifiable)
  ├─ Generate confidence score (0.0-1.0)
  └─ Identify evidence URLs
  ↓
Store as Verification record with `verified_by: "hybrid_ai_human"`
  ├─ Verdict is AI suggestion
  ├─ Confidence reflects AI certainty
  └─ Awaits human review for finalization
  ↓
Human Curator Review
  ├─ Accept AI verdict (auto-publish)
  ├─ Refine verdict (publish with curator notes)
  └─ Reject and flag for research (keep as unverified)
```

### AI Verifier Components

#### 1. Fact Database Connection
```go
// backend/internal/verifier/fact_db.go
type FactDatabase interface {
  GetFacts(query string) []Fact        // Search by keyword
  GetOfficialFacts(id string) []Fact   // Historical statements by official
  GetMetricHistory(metric string) []HistoricalValue
}

// Implementations:
// - OpenAI GPT embeddings (search semantic similarity)
// - Wikidata/DBpedia (structured facts about Ukraine)
// - Ukrainian statistics portal (StatData.gov.ua)
```

#### 2. Evidence Search
```go
// backend/internal/verifier/evidence_searcher.go
type EvidenceSearcher struct {
  googleAPI     *http.Client
  wikipediaAPI  *http.Client
  statportals   []string  // Ukrainian, EU, World Bank data
}

func (es *EvidenceSearcher) FindEvidence(statement string) []Evidence {
  // 1. Generate search queries from statement keywords
  // 2. Search Google Scholar, Wikipedia, government sources
  // 3. Return top N results with relevance scores
}
```

#### 3. Verdict Assignment
```go
// backend/internal/verifier/verdict.go
type VerdictAssigner struct {
  model *OpenAIClient  // or local fine-tuned model
}

func (va *VerdictAssigner) AssignVerdict(statement string, evidence []Evidence) (*Verification, error) {
  // Prompt: "Given the statement: '{}' and evidence [{}], assign a verdict: true/false/partially/misleading/unverifiable"
  // Return structured JSON: { verdict, confidence, reasoning }
}
```

### Knowledge Sources

1. **Structured Data**:
   - World Bank Open Data (GDP, HDI metrics)
   - UN Stats (social indicators)
   - EU Open Data Portal
   - Ukrainian State Statistics Service (StatData.gov.ua)

2. **News Archives**:
   - Wikipedia articles on Ukraine policies
   - News agency archives (UNIAN, RBC)
   - Government press releases

3. **Specialist Databases**:
   - Transparency International (corruption indices)
   - Global Peace Index (security metrics)
   - Environmental data (AirQuality, satelliteio)

---

## Phased Rollout

### Phase 1 (Complete - Current)
✅ JSON-backed MVP with 10 officials, 20 statements, 8 verifications
- Serves as baseline data model
- Frontend dashboard displays sample data
- Backend API functional with HTTP handlers

### Phase 2a (Next - Scraper)
⏳ Build web scraper targeting 1+1.ua
- Extract latest political statements
- Link to official database
- Queue for verification
- **Effort**: 2-3 weeks
- **Deliverable**: New statements in `pending` status flowing into dashboard

### Phase 2b (After 2a - NLP Extraction)
⏳ Improve statement extraction confidence
- Move from regex to AI/transformer-based extraction
- Fine-tune on Ukrainian political speech corpus
- **Effort**: 2-4 weeks
- **Deliverable**: >85% extraction accuracy

### Phase 2c (Parallel with 2a/2b - AI Verifier)
⏳ Implement automated fact-checking
- Connect to multiple evidence sources
- Assign preliminary verdicts with confidence
- Build human review interface for curators
- **Effort**: 4-6 weeks (core logic + evidence integration)
- **Deliverable**: Verifications auto-generated with curator sign-off

### Phase 3 (Post-Phase 2)
⏳ MongoDB Migration
- Replace JSON files with MongoDB Atlas
- Implement data persistence layer
- Enable scalable multi-user curation

---

## Data Flow Diagram

```
1+1.ua / TSN / UNIAN
    ↓
Scraper (colly + goquery)
    ↓
Statement Extractor (regex → NLP)
    ↓
new_statements.json
    ↓
Backend → /api/v1/statements (status="pending")
    ↓
Dashboard displays "Pending Statements" list
    ↓
AI Verifier (async job)
  ├─ Search evidence sources
  ├─ Assign preliminary verdict
  └─ Create Verification record (verified_by="hybrid_ai_human")
    ↓
Dashboard displays "Marked for Review"
    ↓
Human Curator (optional)
  ├─ Accept AI verdict
  ├─ Refine with expert notes
  └─ Publish or flag for research
    ↓
Backend updates Verification (verified_by="human")
    ↓
Dashboard displays final verdict badge
    ↓
Citizens see fact-check result
```

---

## Local Development Commands (Future)

```bash
# Run scraper for specific sources
make scrape-sources SOURCES="1plus1,tsn"

# Run AI verifier on pending statements
make verify-pending

# Review verification queue
make review-queue

# Migrate JSON to MongoDB
make migrate-to-mongodb

# Start full stack with scraper jobs
make up-with-jobs
```

---

## Success Metrics

### Phase 2a (Scraper)
- ✅ 50+ new statements captured per week from 1+1.ua
- ✅ 90% valid statement extraction (linked to real officials)
- ✅ <1% duplicate rate against existing statements

### Phase 2b (NLP)
- ✅ 85%+ extraction confidence on political speech
- ✅ Named entity recognition: 92%+ accuracy on official names

### Phase 2c (Verifier)
- ✅ 70%+ agreement between AI verdict and human curator review
- ✅ Average verdict generation: <30 seconds per statement
- ✅ 95%+ of fact-checks completed within 24 hours of statement capture

---

## Current Status

**MVP (Phase 1)**: ✅ COMPLETE
- 10 officials, 20 statements, 8 verifications in JSON
- Frontend dashboard displays all data
- Backend API serves all endpoints
- TypeScript/Go type safety ensures data consistency

**Next Action**: Begin Phase 2a (scraper implementation)
- Set up colly scraper client
- Target 1+1.ua homepage and article pages
- Extract 5-10 test statements and validate schema
