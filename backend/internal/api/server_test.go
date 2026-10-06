package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/diputat/diputat/backend/internal/config"
)

func TestHealthRoute(t *testing.T) {
	server := NewServer(config.Config{})
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestUnknownRoute(t *testing.T) {
	server := NewServer(config.Config{})
	req := httptest.NewRequest(http.MethodGet, "/missing", nil)
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", rec.Code)
	}
}
