package storage

import (
	"time"
)

// RelationshipType tracks the explicit context of links between nodes
type RelationshipType string

const (
	Refutes    RelationshipType = "refutes"
	Supports   RelationshipType = "supports"
	References RelationshipType = "references"
)

// Link is a universal, lightweight standalone directed graph edge
type Link struct {
	ID       string           `json:"id"`        // Prefix: "link_<sha256_of_source_target_type>"
	SourceID string           `json:"source_id"` // Origin node ID (e.g., "stmt_xxx")
	TargetID string           `json:"target_id"` // Target node ID (e.g., "stmt_yyy", "psrc_zzz")
	Type     RelationshipType `json:"type"`
	Context  string           `json:"context"` // Analyst commentary explaining the link
}

// SourcePoint tracks public analytical vectors or trusted web platforms
type SourcePoint struct {
	ID          string  `json:"id"`           // Prefix: "psrc_<unique_slug>"
	Name        string  `json:"name"`         // E.g., "VoxUkraine"
	URL         string  `json:"url"`          // Strict digital root identity
	TrustRating float64 `json:"trust_rating"` // Metric ranging from 0.0 to 1.0
}

// Party maps political machines, their lifespans, and their true owners
type Party struct {
	ID             string    `json:"id"`              // Prefix: "party_<slug>"
	Name           string    `json:"name"`            // E.g., "Блок Петра Порошенка"
	OwnerAuthorID  string    `json:"owner_author_id"` // Ultimate owner link: "auth_poroshenko"
	CreatedAt      time.Time `json:"created_at"`
	EndedAt        time.Time `json:"ended_at,omitempty"`
	PredecessorIDs []string  `json:"predecessor_ids"` // Flat array of previous "party_xxx" node strings
}

// Author captures the human politician as a pristine tracking anchor
type Author struct {
	ID        string    `json:"id"` // Prefix: "auth_<lowercase_lastname>"
	FullName  string    `json:"full_name"`
	BirthDate time.Time `json:"birth_date"`
}

// StatementRecord tracks the physical location of the raw text capture
type StatementRecord struct {
	URL         string    `json:"url"`          // Exact granular deep link to the source text
	PublishedAt time.Time `json:"published_at"` // Timeline coordinate
}

// Statement represents a cryptographically un-tamperable transaction node
type Statement struct {
	ID             string            `json:"id"`               // Strict Cryptographic Prefix: "stmt_<sha256_hash_of_content>"
	AuthorID       string            `json:"author_id"`        // Actor pointer: "auth_xxx"
	CurrentPartyID string            `json:"current_party_id"` // Active party snapshot: "party_xxx"
	SessionType    string            `json:"session_type"`     // Parliament metadata: "plenary", "closed_plenary"
	Content        string            `json:"content"`          // The literal, unalterable speech text
	Records        []StatementRecord `json:"records"`          // Cross-references of publication instances
	CreatedAt      time.Time         `json:"created_at"`
}
