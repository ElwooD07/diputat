package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/diputat/diputat/backend/internal/config"
	"github.com/diputat/diputat/backend/internal/database"
)

type Server struct {
	config config.Config
	repo   database.Repository
	mux    http.Handler
}

func NewServer(cfg config.Config) (*Server, error) {
	repo, err := database.NewJSONRepository(cfg.DataDir)
	if err != nil {
		return nil, err
	}

	return NewServerWithRepository(cfg, repo), nil
}

func NewServerWithRepository(cfg config.Config, repo database.Repository) *Server {
	server := &Server{config: cfg, repo: repo}
	server.mux = Handler(server)
	return server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) GetHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, map[string]string{"status": "ok"})
}

func (s *Server) ListOfficials(w http.ResponseWriter, r *http.Request) {
	payload, err := s.listOfficials(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, payload)
}

func (s *Server) ListStatements(w http.ResponseWriter, r *http.Request, params ListStatementsParams) {
	officialID := ""
	if params.OfficialId != nil {
		officialID = *params.OfficialId
	}
	payload, err := s.listStatements(r.Context(), officialID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, payload)
}

func (s *Server) listOfficials(ctx context.Context) (map[string]any, error) {
	officials, err := s.repo.ListOfficials(ctx)
	if err != nil {
		return nil, err
	}

	return map[string]any{"items": officials}, nil
}

func (s *Server) listStatements(ctx context.Context, officialID string) (map[string]any, error) {
	var (
		statements any
		err        error
	)

	if officialID == "" {
		statements, err = s.repo.ListStatements(ctx)
	} else {
		statements, err = s.repo.ListStatementsByOfficialID(ctx, officialID)
	}
	if err != nil {
		return nil, err
	}

	return map[string]any{"items": statements}, nil
}

func writeJSON(w http.ResponseWriter, payload any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(payload)
}
