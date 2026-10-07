package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithCORSInjectsHeaders(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	handler := withCORS(next)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Header().Get("Access-Control-Allow-Origin") != "http://localhost:4200" {
		t.Fatalf("unexpected allow origin: %q", rec.Header().Get("Access-Control-Allow-Origin"))
	}
	if rec.Header().Get("Access-Control-Allow-Methods") != "GET, POST, OPTIONS, PUT, DELETE" {
		t.Fatalf("unexpected allow methods: %q", rec.Header().Get("Access-Control-Allow-Methods"))
	}
	if rec.Header().Get("Access-Control-Allow-Headers") != "Content-Type, Authorization" {
		t.Fatalf("unexpected allow headers: %q", rec.Header().Get("Access-Control-Allow-Headers"))
	}
}

func TestWithCORSHandlesPreflight(t *testing.T) {
	nextCalled := false
	next := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		nextCalled = true
	})
	handler := withCORS(next)

	req := httptest.NewRequest(http.MethodOptions, "/api/v1/officials", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if nextCalled {
		t.Fatal("next handler should not be called on preflight request")
	}
}
