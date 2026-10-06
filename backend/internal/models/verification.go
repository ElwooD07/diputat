package models

import "time"

// VerificationVerdict defines the review outcome.
type VerificationVerdict string

const (
	VerificationVerdictTruth        VerificationVerdict = "Truth"
	VerificationVerdictHalfTruth    VerificationVerdict = "Half-Truth"
	VerificationVerdictLie          VerificationVerdict = "Lie"
	VerificationVerdictManipulation VerificationVerdict = "Manipulation"
)

var legacyVerificationResultMap = map[string]VerificationVerdict{
	"true":       VerificationVerdictTruth,
	"partially":  VerificationVerdictHalfTruth,
	"false":      VerificationVerdictLie,
	"misleading": VerificationVerdictManipulation,
}

// Evidence captures a source used to support or dispute a statement.
type Evidence struct {
	Type      string    `json:"type"`
	Source    string    `json:"source"`
	URL       string    `json:"url,omitempty"`
	Content   string    `json:"content"`
	Date      time.Time `json:"date,omitempty"`
	Relevance float64   `json:"relevance,omitempty"`
}

// TimelineHistoryEvent describes a related event in a verification history.
type TimelineHistoryEvent struct {
	Date               time.Time `json:"date"`
	Event              string    `json:"event"`
	RelatedStatementID string    `json:"related_statement_id,omitempty"`
}

// VerificationTimeline stores verification execution history.
type VerificationTimeline struct {
	CheckedAt  time.Time              `json:"checked_at"`
	VerifiedBy string                 `json:"verified_by"`
	History    []TimelineHistoryEvent `json:"history,omitempty"`
}

// Contradiction points to a conflicting past statement.
type Contradiction struct {
	StatementID string    `json:"statement_id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

// Verification stores the outcome of a fact-checking workflow.
type Verification struct {
	ID             string               `json:"_id"`
	StatementID    string               `json:"statement_id"`
	EvidenceText   string               `json:"evidence_text,omitempty"`
	EvaluatorID    string               `json:"evaluator_id,omitempty"`
	Citations      []string             `json:"citations,omitempty"`
	Verdict        VerificationVerdict  `json:"verdict,omitempty"`
	Result         string               `json:"result"`
	Confidence     float64              `json:"confidence,omitempty"`
	Evidence       []Evidence           `json:"evidence,omitempty"`
	Timeline       VerificationTimeline `json:"timeline"`
	Contradictions []Contradiction      `json:"contradictions,omitempty"`
	Notes          string               `json:"notes,omitempty"`
	Metadata       VerificationMeta     `json:"metadata"`
}

// VerificationMeta holds operational metadata for a verification document.
type VerificationMeta struct {
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
	VerificationDurationMS int       `json:"verification_duration_ms,omitempty"`
}

// Normalize keeps verdict/evidence aliases consistent.
func (v *Verification) Normalize() {
	if v.EvidenceText == "" && len(v.Evidence) > 0 {
		v.EvidenceText = v.Evidence[0].Content
	}
	if len(v.Citations) == 0 && len(v.Evidence) > 0 {
		v.Citations = make([]string, 0, len(v.Evidence))
		for _, evidence := range v.Evidence {
			if evidence.URL == "" {
				continue
			}
			v.Citations = append(v.Citations, evidence.URL)
		}
	}
	v.Verdict = normalizeVerificationVerdict(v.Verdict, v.Result)
}

func normalizeVerificationVerdict(verdict VerificationVerdict, result string) VerificationVerdict {
	switch verdict {
	case VerificationVerdictTruth, VerificationVerdictHalfTruth, VerificationVerdictLie, VerificationVerdictManipulation:
		return verdict
	}

	mapped, ok := legacyVerificationResultMap[result]
	if !ok {
		return verdict
	}
	return mapped
}
