package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// RelationshipType визначає характер зв'язку між двома заявами
type RelationshipType string

const (
	Refutes    RelationshipType = "refutes"
	Supports   RelationshipType = "supports"
	References RelationshipType = "references"
)

// Author представляє можновладця як окремий вузол графа
type Author struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	BirthDate time.Time `json:"birth_date"`
	Notes     string    `json:"notes"` // Вільна форма для біографії/приміток
}

// ExternalSource представляє медіа або ресурс з динамічним рейтингом довіри
type ExternalSource struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Publisher   string  `json:"publisher"` // Наприклад: "Bihus.Info", "VoxUkraine"
	TrustRating float64 `json:"trust_rating"`
}

// StatementSource описує конкретне місце і час публікації заяви
type StatementSource struct {
	SourceURL   string    `json:"source_url"`
	PublishedAt time.Time `json:"published_at"`
}

// StatementLink описує рекурсивний зв'язок між заявами
type StatementLink struct {
	TargetStatementID string           `json:"target_statement_id"`
	Type              RelationshipType `json:"type"`
	Context           string           `json:"context"` // Додатковий коментар до зв'язку
}

// Statement є центральним вузлом графа заяв
type Statement struct {
	ID                  string            `json:"id"`
	AuthorID            string            `json:"author_id"` // Зв'язок із конкретним Author.ID
	Content             string            `json:"content"`
	Sources             []StatementSource `json:"sources"`               // Масив першоджерел (де і коли сказано)
	RelatedStatements   []StatementLink   `json:"related_statements"`    // Зв'язки "заява -> заява"
	ExternalEvidenceIDs []string          `json:"external_evidence_ids"` // Зв'язки "заява -> ExternalSource.ID"
	CreatedAt           time.Time         `json:"created_at"`
}

// FileStatementRepository зберігає всі сутності як окремі файли для MVP
type FileStatementRepository struct {
	OutputDir string
}

func NewFileStatementRepository(dir string) *FileStatementRepository {
	return &FileStatementRepository{OutputDir: dir}
}

// SaveStatement зберігає заяву
func (r *FileStatementRepository) SaveStatement(ctx context.Context, s *Statement) error {
	if s.ID == "" {
		s.ID = fmt.Sprintf("stmt_%d", time.Now().UnixNano())
	}
	s.CreatedAt = time.Now()
	return r.saveToFile(filepath.Join(r.OutputDir, "statements"), s.ID, s)
}

// Вдосконалений універсальний допоміжний метод для збереження JSON
func (r *FileStatementRepository) saveToFile(subDir string, id string, data interface{}) error {
	fullPath := filepath.Join(subDir)
	if err := os.MkdirAll(fullPath, 0755); err != nil {
		return err
	}
	fileBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(fullPath, fmt.Sprintf("%s.json", id)), fileBytes, 0644)
}
