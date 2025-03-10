package repository

import (
	"context"
	"database/sql"

	"github.com/andil-id/api/model/domain"
)

type OrderRepository interface {
	GetAllOrder(ctx context.Context, db *sql.DB) ([]domain.Orders, error)
	SaveOrder(ctx context.Context, tx *sql.Tx, data domain.Orders) (string, error)
	GetOrderById(ctx context.Context, db *sql.DB, id string) (domain.OrderById, error)
	UpdateOrderStatus(ctx context.Context, tx *sql.Tx, status string, id string) error
	GetOrderByUserId(ctx context.Context, db *sql.DB, id string) ([]domain.OrderEventByUser, error)
	GetOrderByEventId(ctx context.Context, db *sql.DB, id string) ([]domain.OrderByEventId, error)
}
