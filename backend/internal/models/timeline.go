package models

import "time"

// TimelineEvent is a derived event used by the frontend timeline view.
type TimelineEvent struct {
	ID             string    `json:"id"`
	OfficialID     string    `json:"official_id"`
	StatementID    string    `json:"statement_id,omitempty"`
	VerificationID string    `json:"verification_id,omitempty"`
	EventType      string    `json:"event_type"`
	Title          string    `json:"title"`
	Summary        string    `json:"summary"`
	OccurredAt     time.Time `json:"occurred_at"`
	Result         string    `json:"result,omitempty"`
}