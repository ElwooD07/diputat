package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/diputat/diputat/backend/internal/config"
)

type Official struct {
	ID                string  `json:"id"`
	Name              string  `json:"name"`
	CurrentRole       string  `json:"current_role"`
	VerificationScore float32 `json:"verification_score"`
}

type Statement struct {
	ID         string `json:"id"`
	OfficialID string `json:"official_id"`
	Text       string `json:"text"`
	Status     string `json:"status"` // "InProgress", "Done", "Blocked", "Toxic"
	SourceURL  string `json:"source_url"`
}

type Server struct {
	config config.Config
}

func NewServer(cfg config.Config) *Server {
	return &Server{config: cfg}
}

func loadJSON(filename string, target interface{}) error {
	content, err := os.ReadFile(filepath.Join("..", "data", "samples", filename))
	if err != nil {
		return err
	}
	return json.Unmarshal(content, target)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "" {
		path = "/"
	}

	var payload any
	switch {
	case r.Method == http.MethodGet && path == "/health":
		payload = map[string]string{"status": "ok"}
	case r.Method == http.MethodGet && path == "/api/v1/officials":
		var officials []Official
		if err := loadJSON("officials.json", &officials); err != nil {
			http.Error(w, "failed to load officials", http.StatusInternalServerError)
			return
		}
		payload = map[string]any{"items": officials}
	case r.Method == http.MethodGet && path == "/api/v1/statements":
		var statements []Statement
		if err := loadJSON("statements.json", &statements); err != nil {
			http.Error(w, "failed to load statements", http.StatusInternalServerError)
			return
		}
		payload = map[string]any{"items": statements}
	default:
		http.Error(w, "route not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(payload)
}
