package models

import "time"

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
	ID          string          `json:"_id"`
	OfficialID  string          `json:"official_id"`
	Content     string          `json:"content"`
	Source      StatementSource `json:"source"`
	ExtractedBy string          `json:"extracted_by,omitempty"`
	Status      string          `json:"status"`
	Topics      []string        `json:"topics,omitempty"`
	Sentiment   string          `json:"sentiment,omitempty"`
	Metadata    StatementMeta   `json:"metadata"`
}

// StatementMeta holds creation and language metadata for statements.
type StatementMeta struct {
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Language     string    `json:"language,omitempty"`
	AIConfidence float64   `json:"ai_confidence,omitempty"`
}