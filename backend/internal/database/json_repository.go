package database

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/diputat/diputat/backend/internal/models"
)

// seedsDir derives the seeds directory from a dataDir path.
// Assumes dataDir is one level below the data root (e.g. data/samples).
func seedsDir(dataDir string) string {
	return filepath.Join(dataDir, "..", "seeds")
}

// JSONRepository loads stable local sample documents from JSON files.
type JSONRepository struct {
	officials     []models.Official
	statements    []models.Statement
	verifications []models.Verification
}

// NewJSONRepository loads the sample datasets required by the MVP.
func NewJSONRepository(dataDir string) (*JSONRepository, error) {
	officials, err := loadCollection[models.Official](filepath.Join(dataDir, "officials.json"))
	if err != nil {
		return nil, err
	}
	canonical, err := loadCanonical(filepath.Join(seedsDir(dataDir), "canonical_officials.json"))
	if err != nil {
		return nil, fmt.Errorf("load canonical seed: %w", err)
	}
	if err := validateOfficials(officials, canonical); err != nil {
		return nil, fmt.Errorf("officials data failed canonical check: %w", err)
	}

	statements, err := loadCollection[models.Statement](filepath.Join(dataDir, "statements.json"))
	if err != nil {
		return nil, err
	}

	verifications, err := loadCollection[models.Verification](filepath.Join(dataDir, "verifications.json"))
	if err != nil {
		return nil, err
	}

	return &JSONRepository{
		officials:     officials,
		statements:    statements,
		verifications: verifications,
	}, nil
}

func loadCollection[T any](path string) ([]T, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	var items []T
	if err := json.Unmarshal(content, &items); err != nil {
		return nil, fmt.Errorf("decode %s: %w", path, err)
	}

	return items, nil
}

// ListOfficials returns all officials in the local sample dataset.
func (r *JSONRepository) ListOfficials(context.Context) ([]models.Official, error) {
	return append([]models.Official(nil), r.officials...), nil
}

// GetOfficial returns a single official by identifier.
func (r *JSONRepository) GetOfficial(_ context.Context, id string) (models.Official, bool, error) {
	for _, official := range r.officials {
		if official.ID == id {
			return official, true, nil
		}
	}

	return models.Official{}, false, nil
}

// ListStatements returns all statements in the local sample dataset.
func (r *JSONRepository) ListStatements(context.Context) ([]models.Statement, error) {
	return append([]models.Statement(nil), r.statements...), nil
}

// GetStatement returns a single statement by identifier.
func (r *JSONRepository) GetStatement(_ context.Context, id string) (models.Statement, bool, error) {
	for _, statement := range r.statements {
		if statement.ID == id {
			return statement, true, nil
		}
	}

	return models.Statement{}, false, nil
}

// ListVerifications returns all verifications in the local sample dataset.
func (r *JSONRepository) ListVerifications(context.Context) ([]models.Verification, error) {
	return append([]models.Verification(nil), r.verifications...), nil
}

// GetVerification returns a single verification by identifier.
func (r *JSONRepository) GetVerification(_ context.Context, id string) (models.Verification, bool, error) {
	for _, verification := range r.verifications {
		if verification.ID == id {
			return verification, true, nil
		}
	}

	return models.Verification{}, false, nil
}
