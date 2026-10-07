package main

import (
	"log"
	"net/http"

	"github.com/diputat/diputat/backend/internal/api"
	"github.com/diputat/diputat/backend/internal/config"
)

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	cfg := config.Load()
	server, err := api.NewServer(cfg)
	if err != nil {
		log.Fatalf("initialize server: %v", err)
	}
	handler := withCORS(server)

	log.Printf("%s backend listening on :%s using data directory %s", config.PROJECT_NAME, cfg.Port, cfg.DataDir)
	if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
		log.Fatalf("run server: %v", err)
	}

}
