package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/diputat/diputat/backend/internal/config"
	"github.com/diputat/diputat/backend/internal/database"
	"github.com/diputat/diputat/backend/internal/models"
)

// Server serves the MVP fact-checking API.
type Server struct {
	config     config.Config
	repository database.Repository
}

// NewServer builds the HTTP handler for local development.
func NewServer(cfg config.Config, repository database.Repository) *Server {
	return &Server{
		config:     cfg,
		repository: repository,
	}
}

// ServeHTTP routes requests for the MVP API.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.applyCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "" {
		path = "/"
	}

	switch {
	case r.Method == http.MethodGet && path == "/health":
		s.writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	case r.Method == http.MethodGet && path == "/api/v1/officials":
		s.listOfficials(w, r)
	case r.Method == http.MethodGet && strings.HasPrefix(path, "/api/v1/officials/"):
		s.getOfficial(w, r, strings.TrimPrefix(path, "/api/v1/officials/"))
	case r.Method == http.MethodGet && path == "/api/v1/statements":
		s.listStatements(w, r)
	case r.Method == http.MethodGet && strings.HasPrefix(path, "/api/v1/statements/"):
		s.getStatement(w, r, strings.TrimPrefix(path, "/api/v1/statements/"))
	case r.Method == http.MethodGet && path == "/api/v1/verifications":
		s.listVerifications(w, r)
	case r.Method == http.MethodGet && strings.HasPrefix(path, "/api/v1/verifications/"):
		s.getVerification(w, r, strings.TrimPrefix(path, "/api/v1/verifications/"))
	case r.Method == http.MethodGet && path == "/api/v1/timeline":
		s.listTimeline(w, r)
	default:
		s.writeJSON(w, http.StatusNotFound, map[string]string{"error": "route not found"})
	}
}

func (s *Server) applyCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
}

func (s *Server) listOfficials(w http.ResponseWriter, r *http.Request) {
	officials, err := s.repository.ListOfficials(r.Context())
	if err != nil {
		s.handleError(w, err)
		return
	}

	s.writeJSON(w, http.StatusOK, map[string]any{"items": officials})
}

func (s *Server) getOfficial(w http.ResponseWriter, r *http.Request, id string) {
	official, found, err := s.repository.GetOfficial(r.Context(), id)
	if err != nil {
		s.handleError(w, err)
		return
	}
	if !found {
		s.writeJSON(w, http.StatusNotFound, map[string]string{"error": "official not found"})
		return
	}

	s.writeJSON(w, http.StatusOK, official)
}

func (s *Server) listStatements(w http.ResponseWriter, r *http.Request) {
	statements, err := s.repository.ListStatements(r.Context())
	if err != nil {
		s.handleError(w, err)
		return
	}

	if officialID := r.URL.Query().Get("official_id"); officialID != "" {
		filtered := make([]models.Statement, 0, len(statements))
		for _, statement := range statements {
			if statement.OfficialID == officialID {
				filtered = append(filtered, statement)
			}
		}
		statements = filtered
	}

	s.writeJSON(w, http.StatusOK, map[string]any{"items": statements})
}

func (s *Server) getStatement(w http.ResponseWriter, r *http.Request, id string) {
	statement, found, err := s.repository.GetStatement(r.Context(), id)
	if err != nil {
		s.handleError(w, err)
		return
	}
	if !found {
		s.writeJSON(w, http.StatusNotFound, map[string]string{"error": "statement not found"})
		return
	}

	s.writeJSON(w, http.StatusOK, statement)
}

func (s *Server) listVerifications(w http.ResponseWriter, r *http.Request) {
	verifications, err := s.repository.ListVerifications(r.Context())
	if err != nil {
		s.handleError(w, err)
		return
	}

	if statementID := r.URL.Query().Get("statement_id"); statementID != "" {
		filtered := make([]models.Verification, 0, len(verifications))
		for _, verification := range verifications {
			if verification.StatementID == statementID {
				filtered = append(filtered, verification)
			}
		}
		verifications = filtered
	}

	s.writeJSON(w, http.StatusOK, map[string]any{"items": verifications})
}

func (s *Server) getVerification(w http.ResponseWriter, r *http.Request, id string) {
	verification, found, err := s.repository.GetVerification(r.Context(), id)
	if err != nil {
		s.handleError(w, err)
		return
	}
	if !found {
		s.writeJSON(w, http.StatusNotFound, map[string]string{"error": "verification not found"})
		return
	}

	s.writeJSON(w, http.StatusOK, verification)
}

func (s *Server) listTimeline(w http.ResponseWriter, r *http.Request) {
	statements, err := s.repository.ListStatements(r.Context())
	if err != nil {
		s.handleError(w, err)
		return
	}

	verifications, err := s.repository.ListVerifications(r.Context())
	if err != nil {
		s.handleError(w, err)
		return
	}

	officialID := r.URL.Query().Get("official_id")
	events := buildTimeline(statements, verifications, officialID)
	s.writeJSON(w, http.StatusOK, map[string]any{"items": events})
}

func buildTimeline(statements []models.Statement, verifications []models.Verification, officialID string) []models.TimelineEvent {
	statementByID := make(map[string]models.Statement, len(statements))
	events := make([]models.TimelineEvent, 0, len(statements)+len(verifications))

	for _, statement := range statements {
		statementByID[statement.ID] = statement
		if officialID != "" && statement.OfficialID != officialID {
			continue
		}

		events = append(events, models.TimelineEvent{
			ID:          "statement-" + statement.ID,
			OfficialID:  statement.OfficialID,
			StatementID: statement.ID,
			EventType:   "statement_created",
			Title:       statement.Source.Title,
			Summary:     statement.Content,
			OccurredAt:  statement.Source.Date,
		})
	}

	for _, verification := range verifications {
		statement, found := statementByID[verification.StatementID]
		if !found {
			continue
		}
		if officialID != "" && statement.OfficialID != officialID {
			continue
		}

		events = append(events, models.TimelineEvent{
			ID:             "verification-" + verification.ID,
			OfficialID:     statement.OfficialID,
			StatementID:    statement.ID,
			VerificationID: verification.ID,
			EventType:      "verification_completed",
			Title:          "Verification completed",
			Summary:        verification.Verdict,
			OccurredAt:     verification.Timeline.CheckedAt,
			Result:         verification.Result,
		})
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].OccurredAt.After(events[j].OccurredAt)
	})

	return events
}

func (s *Server) handleError(w http.ResponseWriter, err error) {
	log.Printf("request failed: %v", err)
	s.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
}

func (s *Server) writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("encode response: %v", err)
	}
}
