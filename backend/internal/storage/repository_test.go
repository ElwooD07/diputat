package storage

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

func TestIngestRealData(t *testing.T) {
	// 1. Читаємо реальний сирий JSON, який ми підготували
	rawBytes, err := os.ReadFile("../../data/samples/raw_statement_test.json")
	if err != nil {
		t.Fatalf("Failed to read test data file: %v", err)
	}

	// 2. Десеріалізуємо у тимчасову структуру для імпорту
	var rawData map[string]interface{}
	if err := json.Unmarshal(rawBytes, &rawData); err != nil {
		t.Fatalf("Failed to unmarshal raw data: %v", err)
	}

	// 3. Мапимо сирі дані на нашу детерміновану схему Statement
	stmt := &Statement{
		ID:        "test-uuid-123", // В майбутньому тут буде генеритися хэш або UUID
		Author:    rawData["author_raw"].(string),
		Content:   rawData["text_content"].(string),
		SourceURL: rawData["source_url"].(string),
	}

	parsedTime, err := time.Parse(time.RFC3339, rawData["published_at"].(string))
	if err == nil {
		stmt.PublishedAt = parsedTime
	}

	// 4. Твоя перевірка (Assertion)
	if stmt.Author != "Народний депутат Іванов" {
		t.Errorf("Expected author 'Народний депутат Іванов', got '%s'", stmt.Author)
	}

	if stmt.Content == "" {
		t.Error("Statement content should not be empty")
	}
}
