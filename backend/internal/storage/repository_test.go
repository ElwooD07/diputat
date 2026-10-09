package storage

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestContentAddressedGraphIngestion(t *testing.T) {
	outputDir := filepath.Join("..", "..", "..", "data", "samples", "generated")
	repo := NewFileStatementRepository(outputDir)

	_ = os.RemoveAll(outputDir)
	ctx := context.Background()

	// Test Case A: Validate Cryptographic Statement Generation
	stmt := &Statement{
		AuthorID:       "auth_poroshenko",
		CurrentPartyID: "party_es_2019",
		SessionType:    "plenary",
		Content:        "The European Solidarity faction demands full tracking transparency.",
		Records: []StatementRecord{
			{
				URL:         "https://rada.gov.ua",
				PublishedAt: time.Now(),
			},
		},
	}

	err := repo.SaveStatement(ctx, stmt)
	if err != nil {
		t.Fatalf("Failed to save content-addressed statement node: %v", err)
	}

	if !strings.HasPrefix(stmt.ID, "stmt_") {
		t.Errorf("Expected statement ID prefix 'stmt_', got raw token: %s", stmt.ID)
	}

	// Test Case B: Validate Standalone Link Edge Integration
	link := &Link{
		SourceID: stmt.ID,
		TargetID: "psrc_voxukraine",
		Type:     Supports,
		Context:  "Validated baseline speech match from official portal transcripts.",
	}

	// This call relies strictly on your production method in repository.go
	err = repo.SaveLink(ctx, link)
	if err != nil {
		t.Fatalf("Failed to save standalone link edge: %v", err)
	}

	if !strings.HasPrefix(link.ID, "link_") {
		t.Errorf("Expected link ID prefix 'link_', got raw token: %s", link.ID)
	}

	// Test Case C: Validate Lazy Resolution Profile Fallbacks
	ghostAuthorID := "auth_ghost_actor"
	author, err := repo.GetAuthorByID(ctx, ghostAuthorID)
	if err != nil {
		t.Fatalf("GetAuthorByID failed on unmapped actor resolution: %v", err)
	}

	if author.ID != ghostAuthorID {
		t.Errorf("Expected provisioned author ID to match target %s, got %s", ghostAuthorID, author.ID)
	}
}
