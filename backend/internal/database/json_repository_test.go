package database

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/diputat/diputat/backend/internal/models"
)

func TestNewJSONRepositoryInitializesMissingCollections(t *testing.T) {
	t.Parallel()

	dataDir := t.TempDir()
	repo, err := NewJSONRepository(dataDir)
	if err != nil {
		t.Fatalf("NewJSONRepository: %v", err)
	}

	officials, err := repo.ListOfficials(context.Background())
	if err != nil {
		t.Fatalf("ListOfficials: %v", err)
	}
	if len(officials) != 0 {
		t.Fatalf("expected empty officials collection, got %d items", len(officials))
	}
}

func TestJSONRepositoryFilteringAndSorting(t *testing.T) {
	t.Parallel()

	dataDir := t.TempDir()
	repo, err := NewJSONRepository(dataDir)
	if err != nil {
		t.Fatalf("NewJSONRepository: %v", err)
	}

	if err := repo.UpsertOfficial(context.Background(), models.Official{
		ID:                "official-1",
		Name:              "Official One",
		CurrentRole:       "Deputy",
		VerificationScore: 0.44,
	}); err != nil {
		t.Fatalf("UpsertOfficial official-1: %v", err)
	}
	if err := repo.UpsertOfficial(context.Background(), models.Official{
		ID:                "official-2",
		Name:              "Official Two",
		CurrentRole:       "Deputy",
		VerificationScore: 0.91,
	}); err != nil {
		t.Fatalf("UpsertOfficial official-2: %v", err)
	}

	sorted, err := repo.ListOfficialsByVerificationScore(context.Background())
	if err != nil {
		t.Fatalf("ListOfficialsByVerificationScore: %v", err)
	}
	if len(sorted) != 2 {
		t.Fatalf("expected 2 officials, got %d", len(sorted))
	}
	if sorted[0].ID != "official-2" || sorted[1].ID != "official-1" {
		t.Fatalf("unexpected sort order by VerificationScore: %+v", sorted)
	}

	if err := repo.UpsertStatement(context.Background(), models.Statement{
		ID:         "statement-1",
		OfficialID: "official-1",
		Text:       "Claim 1",
		SourceURL:  "https://example.org/1",
		Date:       time.Now().UTC(),
		Status:     models.StatementStatusInProgress,
	}); err != nil {
		t.Fatalf("UpsertStatement statement-1: %v", err)
	}
	if err := repo.UpsertStatement(context.Background(), models.Statement{
		ID:         "statement-2",
		OfficialID: "official-2",
		Text:       "Claim 2",
		SourceURL:  "https://example.org/2",
		Date:       time.Now().UTC(),
		Status:     models.StatementStatusDone,
	}); err != nil {
		t.Fatalf("UpsertStatement statement-2: %v", err)
	}

	filtered, err := repo.ListStatementsByOfficialID(context.Background(), "official-1")
	if err != nil {
		t.Fatalf("ListStatementsByOfficialID: %v", err)
	}
	if len(filtered) != 1 || filtered[0].ID != "statement-1" {
		t.Fatalf("unexpected filtered statements: %+v", filtered)
	}
}

func TestUpsertStatementRejectsInvalidStatus(t *testing.T) {
	t.Parallel()

	dataDir := t.TempDir()
	repo, err := NewJSONRepository(dataDir)
	if err != nil {
		t.Fatalf("NewJSONRepository: %v", err)
	}

	err = repo.UpsertStatement(context.Background(), models.Statement{
		ID:         "statement-invalid",
		OfficialID: "official-1",
		Text:       "Claim",
		Status:     "UnknownStatus",
	})
	if err == nil {
		t.Fatal("expected invalid status error")
	}
}

func TestJSONRepositoryConcurrentUpsertAndReads(t *testing.T) {
	t.Parallel()

	dataDir := t.TempDir()
	repo, err := NewJSONRepository(dataDir)
	if err != nil {
		t.Fatalf("NewJSONRepository: %v", err)
	}

	ctx := context.Background()
	const writers = 12
	var wg sync.WaitGroup
	wg.Add(writers * 2)
	errCh := make(chan error, writers)

	for i := 0; i < writers; i++ {
		i := i
		go func() {
			defer wg.Done()
			if err := repo.UpsertOfficial(ctx, models.Official{
				ID:                fmt.Sprintf("official-%d", i),
				Name:              "Concurrent Official",
				CurrentRole:       "Deputy",
				VerificationScore: float32(i) / 10,
			}); err != nil {
				errCh <- err
			}
		}()
		go func() {
			defer wg.Done()
			_, _ = repo.ListOfficialsByVerificationScore(ctx)
		}()
	}
	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			t.Fatalf("UpsertOfficial in concurrent flow: %v", err)
		}
	}

	officialsPath := filepath.Join(dataDir, officialsFileName)
	content, err := loadOrInitCollection[models.Official](officialsPath)
	if err != nil {
		t.Fatalf("load officials file: %v", err)
	}
	if len(content) == 0 {
		t.Fatal("expected officials to be persisted after concurrent writes")
	}

	raw, err := json.Marshal(content)
	if err != nil {
		t.Fatalf("marshal persisted content: %v", err)
	}
	if len(raw) == 0 {
		t.Fatal("persisted officials content must not be empty")
	}
}
