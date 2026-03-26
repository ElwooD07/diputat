package database

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/diputat/diputat/backend/internal/models"
)

// canonicalRecord is one entry in the canonical_officials.json seed file.
type canonicalRecord struct {
	Handle   string `json:"handle"`
	Position string `json:"position"`
}

// loadCanonical reads the seed file and returns a handle-keyed lookup map.
// The seed file (data/seeds/canonical_officials.json) is the single source of
// truth for known positions; no positions are hardcoded in Go.
func loadCanonical(seedPath string) (map[string]canonicalRecord, error) {
	data, err := os.ReadFile(seedPath)
	if err != nil {
		return nil, fmt.Errorf("read canonical seed %s: %w", seedPath, err)
	}
	var records []canonicalRecord
	if err := json.Unmarshal(data, &records); err != nil {
		return nil, fmt.Errorf("decode canonical seed %s: %w", seedPath, err)
	}
	m := make(map[string]canonicalRecord, len(records))
	for _, r := range records {
		m[r.Handle] = r
	}
	return m, nil
}

// validateOfficials rejects any official whose handle appears in the canonical
// map but whose position disagrees with it.  The repository refuses to start if
// a data file, scraper output, or AI-generated patch assigns the wrong role to
// a known person.
func validateOfficials(officials []models.Official, canonical map[string]canonicalRecord) error {
	for _, o := range officials {
		entry, known := canonical[o.Handle]
		if !known {
			continue
		}
		if o.Position != entry.Position {
			return fmt.Errorf(
				"official %q (handle %q): data says position=%q but canonical position is %q — "+
					"correct officials.json or update canonical_officials.json if the role genuinely changed",
				o.Name, o.Handle, o.Position, entry.Position,
			)
		}
	}
	return nil
}
