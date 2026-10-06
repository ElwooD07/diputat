package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/diputat/diputat/backend/internal/models"
)

const (
	officialsFileName     = "officials.json"
	statementsFileName    = "statements.json"
	verificationsFileName = "verifications.json"
)

// seedsDir derives the seeds directory from a dataDir path.
// Assumes dataDir is one level below the data root (e.g. data/samples).
func seedsDir(dataDir string) string {
	return filepath.Join(dataDir, "..", "seeds")
}

// JSONRepository stores and persists MVP data in local JSON collections.
type JSONRepository struct {
	mu sync.RWMutex

	officialsPath     string
	statementsPath    string
	verificationsPath string

	officials     []models.Official
	statements    []models.Statement
	verifications []models.Verification
}

// NewJSONRepository loads the datasets from local JSON collections.
func NewJSONRepository(dataDir string) (*JSONRepository, error) {
	if strings.TrimSpace(dataDir) == "" {
		return nil, errors.New("data directory must not be empty")
	}

	repo := &JSONRepository{
		officialsPath:     filepath.Join(dataDir, officialsFileName),
		statementsPath:    filepath.Join(dataDir, statementsFileName),
		verificationsPath: filepath.Join(dataDir, verificationsFileName),
	}

	officials, err := loadOrInitCollection[models.Official](repo.officialsPath)
	if err != nil {
		return nil, err
	}
	for i := range officials {
		officials[i].Normalize()
	}

	if err := validateCanonicalIfPresent(dataDir, officials); err != nil {
		return nil, err
	}

	statements, err := loadOrInitCollection[models.Statement](repo.statementsPath)
	if err != nil {
		return nil, err
	}
	for i := range statements {
		statements[i].Normalize()
	}

	verifications, err := loadOrInitCollection[models.Verification](repo.verificationsPath)
	if err != nil {
		return nil, err
	}
	for i := range verifications {
		verifications[i].Normalize()
	}

	repo.officials = officials
	repo.statements = statements
	repo.verifications = verifications

	return repo, nil
}

func validateCanonicalIfPresent(dataDir string, officials []models.Official) error {
	canonical, err := loadCanonical(filepath.Join(seedsDir(dataDir), "canonical_officials.json"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("load canonical seed: %w", err)
	}
	if err := validateOfficials(officials, canonical); err != nil {
		return fmt.Errorf("officials data failed canonical check: %w", err)
	}
	return nil
}

func loadOrInitCollection[T any](path string) ([]T, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("read %s: %w", path, err)
		}
		empty := make([]T, 0)
		if err := writeCollection(path, empty); err != nil {
			return nil, err
		}
		return empty, nil
	}

	if strings.TrimSpace(string(content)) == "" {
		return make([]T, 0), nil
	}

	var items []T
	if err := json.Unmarshal(content, &items); err != nil {
		return nil, fmt.Errorf("decode %s: %w", path, err)
	}

	return items, nil
}

func writeCollection[T any](path string, items []T) error {
	payload, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return fmt.Errorf("encode %s: %w", path, err)
	}
	payload = append(payload, '\n')

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create directory for %s: %w", path, err)
	}

	tmp, err := os.CreateTemp(filepath.Dir(path), filepath.Base(path)+".tmp-*")
	if err != nil {
		return fmt.Errorf("create temp file for %s: %w", path, err)
	}
	tmpPath := tmp.Name()
	defer func() {
		_ = os.Remove(tmpPath)
	}()

	if _, err := tmp.Write(payload); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("write temp file for %s: %w", path, err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("close temp file for %s: %w", path, err)
	}

	if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("replace %s: %w", path, err)
	}
	if err := os.Rename(tmpPath, path); err != nil {
		return fmt.Errorf("rename temp file for %s: %w", path, err)
	}

	return nil
}

// ListOfficials returns all officials in local storage.
func (r *JSONRepository) ListOfficials(context.Context) ([]models.Official, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return append([]models.Official(nil), r.officials...), nil
}

// ListOfficialsByVerificationScore returns officials sorted by score in descending order.
func (r *JSONRepository) ListOfficialsByVerificationScore(context.Context) ([]models.Official, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := append([]models.Official(nil), r.officials...)
	sort.Slice(out, func(i, j int) bool {
		return out[i].VerificationScore > out[j].VerificationScore
	})
	return out, nil
}

// GetOfficial returns a single official by identifier.
func (r *JSONRepository) GetOfficial(_ context.Context, id string) (models.Official, bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, official := range r.officials {
		if official.ID == id {
			return official, true, nil
		}
	}
	return models.Official{}, false, nil
}

// UpsertOfficial creates or updates an official and persists the collection.
func (r *JSONRepository) UpsertOfficial(_ context.Context, official models.Official) error {
	if strings.TrimSpace(official.ID) == "" {
		return errors.New("official id must not be empty")
	}
	official.Normalize()

	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.officials {
		if r.officials[i].ID != official.ID {
			continue
		}
		r.officials[i] = official
		return writeCollection(r.officialsPath, r.officials)
	}

	r.officials = append(r.officials, official)
	return writeCollection(r.officialsPath, r.officials)
}

// ListStatements returns all statements in local storage.
func (r *JSONRepository) ListStatements(context.Context) ([]models.Statement, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return append([]models.Statement(nil), r.statements...), nil
}

// ListStatementsByOfficialID returns statements for a single official.
func (r *JSONRepository) ListStatementsByOfficialID(_ context.Context, officialID string) ([]models.Statement, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	filtered := make([]models.Statement, 0)
	for _, statement := range r.statements {
		if statement.OfficialID == officialID {
			filtered = append(filtered, statement)
		}
	}
	return filtered, nil
}

// GetStatement returns a single statement by identifier.
func (r *JSONRepository) GetStatement(_ context.Context, id string) (models.Statement, bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, statement := range r.statements {
		if statement.ID == id {
			return statement, true, nil
		}
	}
	return models.Statement{}, false, nil
}

// UpsertStatement creates or updates a statement and persists the collection.
func (r *JSONRepository) UpsertStatement(_ context.Context, statement models.Statement) error {
	if strings.TrimSpace(statement.ID) == "" {
		return errors.New("statement id must not be empty")
	}
	statement.Normalize()

	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.statements {
		if r.statements[i].ID != statement.ID {
			continue
		}
		r.statements[i] = statement
		return writeCollection(r.statementsPath, r.statements)
	}

	r.statements = append(r.statements, statement)
	return writeCollection(r.statementsPath, r.statements)
}

// ListVerifications returns all verifications in local storage.
func (r *JSONRepository) ListVerifications(context.Context) ([]models.Verification, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return append([]models.Verification(nil), r.verifications...), nil
}

// GetVerification returns a single verification by identifier.
func (r *JSONRepository) GetVerification(_ context.Context, id string) (models.Verification, bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, verification := range r.verifications {
		if verification.ID == id {
			return verification, true, nil
		}
	}
	return models.Verification{}, false, nil
}

// UpsertVerification creates or updates a verification and persists the collection.
func (r *JSONRepository) UpsertVerification(_ context.Context, verification models.Verification) error {
	if strings.TrimSpace(verification.ID) == "" {
		return errors.New("verification id must not be empty")
	}
	verification.Normalize()

	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.verifications {
		if r.verifications[i].ID != verification.ID {
			continue
		}
		r.verifications[i] = verification
		return writeCollection(r.verificationsPath, r.verifications)
	}

	r.verifications = append(r.verifications, verification)
	return writeCollection(r.verificationsPath, r.verifications)
}
