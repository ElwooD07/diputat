package models

import "time"

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
	Result         string               `json:"result"`
	Confidence     float64              `json:"confidence,omitempty"`
	Verdict        string               `json:"verdict,omitempty"`
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