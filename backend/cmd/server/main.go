package main

import (
	"log"
	"net/http"

	"github.com/diputat/diputat/backend/internal/api"
	"github.com/diputat/diputat/backend/internal/config"
	"github.com/diputat/diputat/backend/internal/database"
)

func main() {
	cfg := config.Load()

	repository, err := database.NewJSONRepository(cfg.DataDir)
	if err != nil {
		log.Fatalf("load repository: %v", err)
	}

	server := api.NewServer(cfg, repository)

	log.Printf("%s backend listening on :%s using data directory %s", config.PROJECT_NAME, cfg.Port, cfg.DataDir)
	if err := http.ListenAndServe(":"+cfg.Port, server); err != nil {
		log.Fatalf("run server: %v", err)
	}

}
