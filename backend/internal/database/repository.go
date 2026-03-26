package database

import (
	"context"

	"github.com/diputat/diputat/backend/internal/models"
)

// Repository exposes the read operations needed by the MVP API.
type Repository interface {
	ListOfficials(ctx context.Context) ([]models.Official, error)
	GetOfficial(ctx context.Context, id string) (models.Official, bool, error)
	ListStatements(ctx context.Context) ([]models.Statement, error)
	GetStatement(ctx context.Context, id string) (models.Statement, bool, error)
	ListVerifications(ctx context.Context) ([]models.Verification, error)
	GetVerification(ctx context.Context, id string) (models.Verification, bool, error)
}