package storage

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// FileStatementRepository handles pure disk-based storage mappings
type FileStatementRepository struct {
	OutputDir string
}

func NewFileStatementRepository(dir string) *FileStatementRepository {
	return &FileStatementRepository{OutputDir: dir}
}

// SaveStatement commits a statement node to disk under the strict prefix scheme
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

// GetAuthorByID implements clean Lazy Resolution with custom placeholder provisioning
func (r *FileStatementRepository) GetAuthorByID(ctx context.Context, id string) (*Author, error) {
	filePath := filepath.Join(r.OutputDir, "authors", id+".json")
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Author{
				ID:        id,
				FullName:  "Pending Verification / New Actor",
				BirthDate: time.Time{},
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

// SaveLink commits a standalone link edge to disk under a content-addressed schema
func (r *FileStatementRepository) SaveLink(ctx context.Context, l *Link) error {
	if l.SourceID == "" || l.TargetID == "" || l.Type == "" {
		return fmt.Errorf("invalid link parameters: source, target, and type are mandatory")
	}

	// Compute link cryptographic ID based on composite unique attributes
	compositeKey := fmt.Sprintf("%s_%s_%s", l.SourceID, l.TargetID, string(l.Type))
	hash := sha256.Sum256([]byte(compositeKey))
	l.ID = fmt.Sprintf("link_%x", hash)

	return r.saveToFile(filepath.Join(r.OutputDir, "links"), l.ID, l)
}
