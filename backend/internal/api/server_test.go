package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/diputat/diputat/backend/internal/config"
	"github.com/diputat/diputat/backend/internal/models"
)

type stubRepository struct {
	officials     []models.Official
	statements    []models.Statement
	verifications []models.Verification
}

func (s *stubRepository) ListOfficials(context.Context) ([]models.Official, error) {
	return append([]models.Official(nil), s.officials...), nil
}

func (s *stubRepository) ListOfficialsByVerificationScore(context.Context) ([]models.Official, error) {
	return append([]models.Official(nil), s.officials...), nil
}

func (s *stubRepository) GetOfficial(_ context.Context, id string) (models.Official, bool, error) {
	for _, official := range s.officials {
		if official.ID == id {
			return official, true, nil
		}
	}
	return models.Official{}, false, nil
}

func (s *stubRepository) UpsertOfficial(_ context.Context, official models.Official) error {
	s.officials = append(s.officials, official)
	return nil
}

func (s *stubRepository) ListStatements(context.Context) ([]models.Statement, error) {
	return append([]models.Statement(nil), s.statements...), nil
}

func (s *stubRepository) ListStatementsByOfficialID(_ context.Context, officialID string) ([]models.Statement, error) {
	filtered := make([]models.Statement, 0)
	for _, statement := range s.statements {
		if statement.OfficialID == officialID {
			filtered = append(filtered, statement)
		}
	}
	return filtered, nil
}

func (s *stubRepository) GetStatement(_ context.Context, id string) (models.Statement, bool, error) {
	for _, statement := range s.statements {
		if statement.ID == id {
			return statement, true, nil
		}
	}
	return models.Statement{}, false, nil
}

func (s *stubRepository) UpsertStatement(_ context.Context, statement models.Statement) error {
	s.statements = append(s.statements, statement)
	return nil
}

func (s *stubRepository) ListVerifications(context.Context) ([]models.Verification, error) {
	return append([]models.Verification(nil), s.verifications...), nil
}

func (s *stubRepository) GetVerification(_ context.Context, id string) (models.Verification, bool, error) {
	for _, verification := range s.verifications {
		if verification.ID == id {
			return verification, true, nil
		}
	}
	return models.Verification{}, false, nil
}

func (s *stubRepository) UpsertVerification(_ context.Context, verification models.Verification) error {
	s.verifications = append(s.verifications, verification)
	return nil
}

func TestHealthRoute(t *testing.T) {
	server := NewServerWithRepository(config.Config{}, &stubRepository{})
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestOfficialsRouteReturnsUnifiedIDShape(t *testing.T) {
	server := NewServerWithRepository(config.Config{}, &stubRepository{
		officials: []models.Official{
			{
				ID:                "official-1",
				Name:              "Official One",
				CurrentRole:       "Deputy",
				VerificationScore: 77,
			},
		},
	})
	req := httptest.NewRequest(http.MethodGet, "/api/v1/officials", nil)
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var payload struct {
		Items []map[string]any `json:"items"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(payload.Items) != 1 {
		t.Fatalf("expected 1 official, got %d", len(payload.Items))
	}
	if _, ok := payload.Items[0]["id"]; !ok {
		t.Fatalf("expected field id in response item, got %+v", payload.Items[0])
	}
	if _, ok := payload.Items[0]["_id"]; ok {
		t.Fatalf("did not expect legacy _id field in response item, got %+v", payload.Items[0])
	}
}

func TestStatementsRouteSupportsOfficialFilter(t *testing.T) {
	server := NewServerWithRepository(config.Config{}, &stubRepository{
		statements: []models.Statement{
			{ID: "s-1", OfficialID: "o-1", Text: "A", Status: models.StatementStatusInProgress},
			{ID: "s-2", OfficialID: "o-2", Text: "B", Status: models.StatementStatusDone},
		},
	})
	req := httptest.NewRequest(http.MethodGet, "/api/v1/statements?official_id=o-2", nil)
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var payload struct {
		Items []models.Statement `json:"items"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(payload.Items) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(payload.Items))
	}
	if payload.Items[0].OfficialID != "o-2" {
		t.Fatalf("expected statement filtered by official_id=o-2, got %+v", payload.Items[0])
	}
}

func TestUnknownRoute(t *testing.T) {
	server := NewServerWithRepository(config.Config{}, &stubRepository{})
	req := httptest.NewRequest(http.MethodGet, "/missing", nil)
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", rec.Code)
	}
}
