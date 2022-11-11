package repository

import (
	"context"
	"database/sql"

	"github.com/andil-id/api/model/domain"
)

type EventRepository interface {
	SaveEvent(ctx context.Context, tx *sql.Tx, event domain.Events) (string, error)
	GetEventById(ctx context.Context, db *sql.DB, id string) (domain.Events, error)
	GetAllEvents(ctx context.Context, db *sql.DB) ([]domain.Events, error)
}
