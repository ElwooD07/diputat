package storage

import (
	"context"
	"time"
)

// Statement представляє структуру заяви, адаптовану під реальні дані
type Statement struct {
	ID          string    `json:"id"`
	Author      string    `json:"author"`
	Content     string    `json:"content"`
	SourceURL   string    `json:"source_url"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
}

// StatementRepository описує контракт для роботи з базою даних або файлами
type StatementRepository interface {
	Save(ctx context.Context, statement *Statement) error
	GetByID(ctx context.Context, id string) (*Statement, error)
	List(ctx context.Context, limit int, offset int) ([]*Statement, error)
}
