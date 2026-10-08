package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// RelationshipType defines the direction and intent of a link between statements
type RelationshipType string

const (
	Refutes    RelationshipType = "refutes"
	Supports   RelationshipType = "supports"
	References RelationshipType = "references"
)

// Author represents a public official as an independent graph node
type Author struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	BirthDate time.Time `json:"birth_date"`
	Notes     string    `json:"notes"`
}

// ExternalSource represents an outside tracking vector or media publisher
type ExternalSource struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Publisher   string  `json:"publisher"`
	TrustRating float64 `json:"trust_rating"`
}

// StatementSource maps out a concrete physical publication record
type StatementSource struct {
	SourceURL   string    `json:"source_url"`
	PublishedAt time.Time `json:"published_at"`
}

// StatementLink structures directional links between separate claims
type StatementLink struct {
	TargetStatementID string           `json:"target_statement_id"`
	Type              RelationshipType `json:"type"`
	Context           string           `json:"context"`
}

// Statement serves as the foundational data node of our ecosystem
type Statement struct {
	ID                  string            `json:"id"`
	AuthorID            string            `json:"author_id"`
	Content             string            `json:"content"`
	Sources             []StatementSource `json:"sources"`
	RelatedStatements   []StatementLink   `json:"related_statements"`
	ExternalEvidenceIDs []string          `json:"external_evidence_ids"`
	CreatedAt           time.Time         `json:"created_at"`
}

// FileStatementRepository handles pure disk-based storage mappings
type FileStatementRepository struct {
	OutputDir string
}

func NewFileStatementRepository(dir string) *FileStatementRepository {
	return &FileStatementRepository{OutputDir: dir}
}

// SaveStatement commits a statement node to disk
func (r *FileStatementRepository) SaveStatement(ctx context.Context, s *Statement) error {
	if s.ID == "" {
		s.ID = fmt.Sprintf("stmt_%d", time.Now().UnixNano())
	}
	s.CreatedAt = time.Now()
	return r.saveToFile(filepath.Join(r.OutputDir, "statements"), s.ID, s)
}

// GetStatementByID loads a strict statement layout from disk
func (r *FileStatementRepository) GetStatementByID(ctx context.Context, id string) (*Statement, error) {
	filePath := filepath.Join(r.OutputDir, "statements", id+".json")
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("statement node %s not found: %w", id, err)
	}

	var s Statement
	if err := json.Unmarshal(fileBytes, &s); err != nil {
		return nil, fmt.Errorf("failed to parse statement JSON: %w", err)
	}
	return &s, nil
}

// GetAuthorByID implements our Approach 2 (Lazy Resolution / Placeholder Generation)
func (r *FileStatementRepository) GetAuthorByID(ctx context.Context, id string) (*Author, error) {
	filePath := filepath.Join(r.OutputDir, "authors", id+".json")
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Clean Lazy Resolution: Node is missing, provision an on-the-fly human profile
			return &Author{
				ID:        id,
				FullName:  "Pending Verification / New Actor",
				BirthDate: time.Time{},
				Notes:     "Auto-provisioned on missing graph node resolution.",
			}, nil
		}
		return nil, fmt.Errorf("failed to access author node file: %w", err)
	}

	var a Author
	if err := json.Unmarshal(fileBytes, &a); err != nil {
		return nil, fmt.Errorf("failed to parse author JSON: %w", err)
	}
	return &a, nil
}

// Low-level helper method handling isolated object output formatting
func (r *FileStatementRepository) saveToFile(subDir string, id string, data interface{}) error {
	if err := os.MkdirAll(subDir, 0755); err != nil {
		return err
	}
	fileBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(subDir, fmt.Sprintf("%s.json", id)), fileBytes, 0644)
}
