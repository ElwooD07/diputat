package config

import "os"

const (
	defaultPort    = "8080"
	defaultDataDir = "../data/samples"
)

// Config stores runtime settings for the local backend.
type Config struct {
	Port    string
	DataDir string
}

// Load builds configuration from environment variables with local defaults.
func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		dataDir = defaultDataDir
	}

	return Config{
		Port:    port,
		DataDir: dataDir,
	}
}