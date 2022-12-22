package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/andil-id/api/model/domain"
	"github.com/segmentio/ksuid"
)

type EventRepositoryImpl struct {
}

func NewEventRepository() EventRepository {
	return &EventRepositoryImpl{}
}

func (r *EventRepositoryImpl) SaveEvent(ctx context.Context, tx *sql.Tx, event domain.Events) (string, error) {
	var err error
	id := ksuid.New().String()
	switch event.Type {
	case "paid":
		SQL := "INSERT INTO events (id, title, description, banner, certificate, price, type, bank_account_num, bank_account_name, recipient_name, location, start_at, end_at, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
		_, err = tx.ExecContext(ctx, SQL, id, event.Title, event.Description, event.Banner, event.Certificate, event.Price, event.Type, event.BankAccountNum, event.BackAccountName, event.RecipientName, event.Location, event.StartAt, event.EndAt, time.Now(), time.Now())
	case "free":
		SQL := "INSERT INTO events (id, title, description, banner, certificate, price, type, location, start_at, end_at, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
		_, err = tx.ExecContext(ctx, SQL, id, event.Title, event.Description, event.Banner, event.Certificate, event.Price, event.Type, event.Location, event.StartAt, event.EndAt, time.Now(), time.Now())
	default:
		return "", errors.New("event type not supported")

	}
	if err != nil {
		panic(err)
	}
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *EventRepositoryImpl) GetAllEvents(ctx context.Context, db *sql.DB) ([]domain.Events, error) {
	SQL := "SELECT * FROM events"
	rows, err := db.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var events []domain.Events
	for rows.Next() {
		var event domain.Events
		var bankAccountNum sql.NullString
		var bankAccountName sql.NullString
		var recipientName sql.NullString
		err := rows.Scan(&event.Id, &event.Title, &event.Description, &event.Banner, &event.Certificate, &event.Price, &event.Type, &bankAccountNum, &bankAccountName, &recipientName, &event.Location, &event.StartAt, &event.EndAt, &event.CreatedAt, &event.UpdatedAt)
		if err != nil {
			panic(err)
		}
		if bankAccountNum.Valid {
			event.BankAccountNum = bankAccountNum.String
		}
		if bankAccountName.Valid {
			event.BackAccountName = bankAccountName.String
		}
		if recipientName.Valid {
			event.RecipientName = recipientName.String
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *EventRepositoryImpl) GetEventById(ctx context.Context, db *sql.DB, id string) (domain.Events, error) {
	SQL := "SELECT * FROM events WHERE id = ?"
	rows, err := db.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	var event domain.Events
	if rows.Next() {
		var bankAccountNum sql.NullString
		var bankAccountName sql.NullString
		var recipientName sql.NullString
		err := rows.Scan(&event.Id, &event.Title, &event.Description, &event.Banner, &event.Certificate, &event.Price, &event.Type, &bankAccountNum, &bankAccountName, &recipientName, &event.Location, &event.StartAt, &event.EndAt, &event.CreatedAt, &event.UpdatedAt)
		if err != nil {
			panic(err)
		}
		if bankAccountNum.Valid {
			event.BankAccountNum = bankAccountNum.String
		}
		if bankAccountName.Valid {
			event.BackAccountName = bankAccountName.String
		}
		if recipientName.Valid {
			event.RecipientName = recipientName.String
		}
		return event, nil
	}
	return event, errors.New("event not found")
}
