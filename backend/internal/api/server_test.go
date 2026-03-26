package api

import (
	"testing"
	"time"

	"github.com/diputat/diputat/backend/internal/models"
)

func TestBuildTimelineSortsAndFilters(t *testing.T) {
	statements := []models.Statement{
		{
			ID:         "statement-1",
			OfficialID: "official-1",
			Content:    "Budget spending increased by 12%.",
			Source: models.StatementSource{
				Title: "Budget interview",
				Date:  time.Date(2026, 3, 20, 9, 0, 0, 0, time.UTC),
			},
		},
		{
			ID:         "statement-2",
			OfficialID: "official-2",
			Content:    "Road repairs are complete.",
			Source: models.StatementSource{
				Title: "Press conference",
				Date:  time.Date(2026, 3, 19, 9, 0, 0, 0, time.UTC),
			},
		},
	}

	verifications := []models.Verification{
		{
			ID:          "verification-1",
			StatementID: "statement-1",
			Result:      "partially",
			Verdict:     "Budget spending rose, but not by the claimed amount.",
			Timeline: models.VerificationTimeline{
				CheckedAt: time.Date(2026, 3, 21, 10, 0, 0, 0, time.UTC),
			},
		},
	}

	events := buildTimeline(statements, verifications, "official-1")
	if len(events) != 2 {
		t.Fatalf("expected 2 events, got %d", len(events))
	}

	if events[0].EventType != "verification_completed" {
		t.Fatalf("expected latest event to be verification, got %s", events[0].EventType)
	}

	if events[1].StatementID != "statement-1" {
		t.Fatalf("expected filtered statement to remain, got %s", events[1].StatementID)
	}
}
