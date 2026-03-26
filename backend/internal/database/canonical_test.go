package database

import (
	"testing"

	"github.com/diputat/diputat/backend/internal/models"
)

var testCanonical = map[string]canonicalRecord{
	"tracked_handle": {
		Handle:   "tracked_handle",
		Position: "Deputy (Verkhovna Rada)",
	},
}

func TestCanonicalValidationAcceptsCorrectPositions(t *testing.T) {
	officials := []models.Official{{
		Handle:   "tracked_handle",
		Position: "Deputy (Verkhovna Rada)",
	}}

	if err := validateOfficials(officials, testCanonical); err != nil {
		t.Fatalf("validateOfficials rejected correct data: %v", err)
	}
}

func TestCanonicalValidationRejectsBadPosition(t *testing.T) {
	officials := []models.Official{{
		Handle:   "tracked_handle",
		Position: "Prime Minister",
	}}

	if err := validateOfficials(officials, testCanonical); err == nil {
		t.Fatal("validateOfficials accepted mismatched canonical position")
	}
}

func TestCanonicalValidationIgnoresUnknownHandles(t *testing.T) {
	officials := []models.Official{{
		Handle:   "unknown_handle",
		Position: "President",
	}}

	if err := validateOfficials(officials, testCanonical); err != nil {
		t.Fatalf("validateOfficials rejected unknown handle: %v", err)
	}
}

func TestCanonicalSeedFileLoads(t *testing.T) {
	canonical, err := loadCanonical("../../../data/seeds/canonical_officials.json")
	if err != nil {
		t.Fatalf("loadCanonical: %v", err)
	}
	if canonical == nil {
		t.Fatal("canonical map must not be nil")
	}
}
