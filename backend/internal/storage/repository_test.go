package storage

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestGraphArchitectureIngestion(t *testing.T) {
	// 1. Setup a clean, predictable generated directory for testing
	// Paths are relative to the 'backend/internal/storage' execution folder
	outputDir := filepath.Join("..", "..", "..", "data", "samples", "generated")
	repo := NewFileStatementRepository(outputDir)

	// Clean up any old test runs
	_ = os.RemoveAll(outputDir)

	ctx := context.Background()

	// Step A: Instantiate and save an Author node
	author := &Author{
		ID:        "auth_ivanov",
		FullName:  "Ivan Ivanov",
		BirthDate: time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC),
		Notes:     "MP since 2019",
	}
	// Our saveToFile method wants: (subDir, id, data)
	err := repo.saveToFile("authors", author.ID, author)
	if err != nil {
		t.Fatalf("Failed to save author node: %v", err)
	}

	// Step B: Instantiate and save an ExternalSource node
	source := &ExternalSource{
		ID:          "src_bihus",
		Title:       "Investigative Report on Budgets",
		Publisher:   "Bihus.Info",
		TrustRating: 0.95,
	}
	err = repo.saveToFile("sources", source.ID, source)
	if err != nil {
		t.Fatalf("Failed to save external source node: %v", err)
	}

	// Step C: Create Statement A (Base Claim by Ivanov) with multiple sources
	stmtA := &Statement{
		ID:       "stmt_base_001",
		AuthorID: author.ID,
		Content:  "No funds were taken from the defense budget.",
		Sources: []StatementSource{
			{
				SourceURL:   "https://t.me",
				PublishedAt: time.Now().Add(-2 * time.Hour),
			},
			{
				SourceURL:   "https://rada.gov.ua",
				PublishedAt: time.Now().Add(-1 * time.Hour),
			},
		},
		ExternalEvidenceIDs: []string{source.ID},
	}
	err = repo.SaveStatement(ctx, stmtA)
	if err != nil {
		t.Fatalf("Failed to save base statement A: %v", err)
	}

	// Step D: Create Statement B (Refuting Claim) referencing Statement A's ID
	stmtB := &Statement{
		ID:       "stmt_refute_002",
		AuthorID: "auth_petrov", // Another author
		Content:  "Ivanov's claims are false, road construction stole defense money.",
		Sources: []StatementSource{
			{
				SourceURL:   "https://facebook.com",
				PublishedAt: time.Now(),
			},
		},
		RelatedStatements: []StatementLink{
			{
				TargetStatementID: stmtA.ID,
				Type:              Refutes,
				Context:           "Direct contradiction based on treasury reports.",
			},
		},
		ExternalEvidenceIDs: []string{source.ID},
	}
	err = repo.SaveStatement(ctx, stmtB)
	if err != nil {
		t.Fatalf("Failed to save refuting statement B: %v", err)
	}

	// 2. Core Assertions: Read back files and verify structural fields match perfectly
	stmtBPath := filepath.Join(outputDir, "statements", stmtB.ID+".json")
	fileBytes, err := os.ReadFile(stmtBPath)
	if err != nil {
		t.Fatalf("Expected saved file %s to exist, but could not read: %v", stmtBPath, err)
	}

	var savedStmtB Statement
	if err := json.Unmarshal(fileBytes, &savedStmtB); err != nil {
		t.Fatalf("Failed to unmarshal saved Statement B file: %v", err)
	}

	// Assert relations held perfectly in the serialized JSON
	if len(savedStmtB.RelatedStatements) != 1 {
		t.Fatalf("Expected 1 related statement edge, found %d", len(savedStmtB.RelatedStatements))
	}

	if savedStmtB.RelatedStatements[0].TargetStatementID != "stmt_base_001" || savedStmtB.RelatedStatements[0].Type != Refutes {
		t.Errorf("Graph relation corrupted. Got Target: %s, Type: %s",
			savedStmtB.RelatedStatements[0].TargetStatementID, savedStmtB.RelatedStatements[0].Type)
	}
}
