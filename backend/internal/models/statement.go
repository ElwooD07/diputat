package models

import "time"

// StatementStatus defines the workflow state of a tracked statement/task.
type StatementStatus string

const (
	StatementStatusInProgress StatementStatus = "InProgress"
	StatementStatusDone       StatementStatus = "Done"
	StatementStatusBlocked    StatementStatus = "Blocked"
	StatementStatusToxic      StatementStatus = "Toxic"
)

var legacyStatementStatusMap = map[StatementStatus]StatementStatus{
	"pending":  StatementStatusInProgress,
	"new":      StatementStatusInProgress,
	"":         StatementStatusInProgress,
	"verified": StatementStatusDone,
}

// StatementSource describes where a statement came from.
type StatementSource struct {
	Type        string    `json:"type"`
	URL         string    `json:"url,omitempty"`
	Title       string    `json:"title,omitempty"`
	Date        time.Time `json:"date"`
	MediaOutlet string    `json:"media_outlet,omitempty"`
}

// Statement is the claim unit used for verification.
type Statement struct {
	ID              string          `json:"_id"`
	OfficialID      string          `json:"official_id"`
	SourceURL       string          `json:"source_url,omitempty"`
	Text            string          `json:"text,omitempty"`
	Date            time.Time       `json:"date,omitempty"`
	Status          StatementStatus `json:"status"`
	BacklogPriority int             `json:"backlog_priority,omitempty"`
	Content         string          `json:"content"`
	Source          StatementSource `json:"source"`
	ExtractedBy     string          `json:"extracted_by,omitempty"`
	Topics          []string        `json:"topics,omitempty"`
	Sentiment       string          `json:"sentiment,omitempty"`
	Metadata        StatementMeta   `json:"metadata"`
}

// StatementMeta holds creation and language metadata for statements.
type StatementMeta struct {
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Language     string    `json:"language,omitempty"`
	AIConfidence float64   `json:"ai_confidence,omitempty"`
}

// Normalize keeps alias fields consistent and maps legacy statuses.
func (s *Statement) Normalize() {
	if s.Text == "" {
		s.Text = s.Content
	}
	if s.Content == "" {
		s.Content = s.Text
	}
	if s.SourceURL == "" {
		s.SourceURL = s.Source.URL
	}
	if s.Source.URL == "" {
		s.Source.URL = s.SourceURL
	}
	if s.Date.IsZero() {
		s.Date = s.Source.Date
	}
	if s.Source.Date.IsZero() {
		s.Source.Date = s.Date
	}
	s.Status = normalizeStatementStatus(s.Status)
	if s.BacklogPriority < 0 {
		s.BacklogPriority = 0
	}
}

func normalizeStatementStatus(status StatementStatus) StatementStatus {
	switch status {
	case StatementStatusInProgress, StatementStatusDone, StatementStatusBlocked, StatementStatusToxic:
		return status
	default:
		if mapped, ok := legacyStatementStatusMap[status]; ok {
			return mapped
		}
		return status
	}
}
