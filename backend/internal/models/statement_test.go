package models

import "testing"

func TestStatementValidateAcceptsNormalizedLegacyStatus(t *testing.T) {
	statement := Statement{
		ID:         "statement-1",
		OfficialID: "official-1",
		Text:       "Claim",
		Status:     "verified",
	}

	statement.Normalize()

	if statement.Status != StatementStatusDone {
		t.Fatalf("expected normalized status %q, got %q", StatementStatusDone, statement.Status)
	}
	if err := statement.Validate(); err != nil {
		t.Fatalf("expected statement to validate, got error: %v", err)
	}
}

func TestStatementValidateRejectsUnknownStatus(t *testing.T) {
	statement := Statement{
		ID:         "statement-1",
		OfficialID: "official-1",
		Text:       "Claim",
		Status:     "NotAStatus",
	}

	statement.Normalize()

	if err := statement.Validate(); err == nil {
		t.Fatal("expected validation error for unknown status")
	}
}

func TestStatementValidateRejectsMissingRequiredFields(t *testing.T) {
	statement := Statement{
		ID:     "statement-1",
		Status: StatementStatusInProgress,
	}

	statement.Normalize()

	if err := statement.Validate(); err == nil {
		t.Fatal("expected validation error for missing required fields")
	}
}
