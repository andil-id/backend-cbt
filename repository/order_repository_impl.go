package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/andil-id/api/model/domain"
	"github.com/segmentio/ksuid"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (r *OrderRepositoryImpl) GetAllOrder(ctx context.Context, db *sql.DB) ([]domain.Orders, error) {
	SQL := "SELECT * FROM orders"
	rows, err := db.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var orders []domain.Orders
	for rows.Next() {
		var order domain.Orders
		var proofPayment *sql.NullString
		err := rows.Scan(&order.Id, &order.UserId, &order.EventId, &order.Amount, &proofPayment, &order.Status, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			panic(err)
		}
		if proofPayment.Valid {
			order.ProofPayment = proofPayment.String
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrderRepositoryImpl) SaveOrder(ctx context.Context, tx *sql.Tx, data domain.Orders) (string, error) {
	id := ksuid.New().String()
	SQL := "INSERT INTO orders (id, user_id, event_id, amount, proof_payment, status, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, id, data.UserId, data.EventId, data.Amount, data.ProofPayment, "WAITING", data.CreatedAt, data.UpdatedAt)
	if err != nil {
		panic(err)
	}
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *OrderRepositoryImpl) GetOrderById(ctx context.Context, db *sql.DB, id string) (domain.OrderById, error) {
	SQL := "SELECT orders.id, orders.user_id, orders.event_id, orders.amount, orders.proof_payment, orders.status, orders.created_at, orders.updated_at, events.location FROM orders LEFT JOIN events ON orders.event_id=events.id WHERE orders.id = ?"
	rows, err := db.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	var order domain.OrderById
	var proofPayment *sql.NullString
	if rows.Next() {
		err := rows.Scan(&order.Id, &order.UserId, &order.EventId, &order.Amount, &proofPayment, &order.Status, &order.CreatedAt, &order.UpdatedAt, &order.Location)
		if err != nil {
			panic(err)
		}
		if proofPayment.Valid {
			order.ProofPayment = proofPayment.String
		}
		return order, nil
	}
	return order, errors.New("event not found")
}

func (r *OrderRepositoryImpl) GetOrderByUserIdAndEventId(ctx context.Context, db *sql.DB, userId string, eventId string) (domain.Orders, error) {
	SQL := "SELECT * FROM orders WHERE user_id = ? AND 	event_id = ?"
	rows, err := db.QueryContext(ctx, SQL, userId, eventId)
	if err != nil {
		panic(err)
	}
	var order domain.Orders
	var proofPayment *sql.NullString
	if rows.Next() {
		err := rows.Scan(&order.Id, &order.UserId, &order.EventId, &order.Amount, &proofPayment, &order.Status, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			panic(err)
		}
		if proofPayment.Valid {
			order.ProofPayment = proofPayment.String
		}
		return order, nil
	}
	return order, errors.New("order not found")
}

func (r *OrderRepositoryImpl) UpdateOrderStatus(ctx context.Context, tx *sql.Tx, status string, id string) error {
	SQL := "UPDATE orders SET status = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, status, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepositoryImpl) GetOrderByUserId(ctx context.Context, db *sql.DB, id string) ([]domain.OrderEventByUser, error) {
	SQL := "SELECT orders.id, orders.user_id, orders.event_id, orders.amount, orders.status, orders.created_at, orders.updated_at, events.title, events.banner, events.location, events.start_at, events.end_at FROM orders LEFT JOIN events ON orders.event_id=events.id WHERE orders.user_id = ?"
	rows, err := db.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var orders []domain.OrderEventByUser
	for rows.Next() {
		var order domain.OrderEventByUser
		err := rows.Scan(&order.Id, &order.UserId, &order.EventId, &order.Amount, &order.Status, &order.CreatedAt, &order.UpdatedAt, &order.Title, &order.Banner, &order.Location, &order.StartAt, &order.EndAt)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrderRepositoryImpl) GetOrderByEventId(ctx context.Context, db *sql.DB, id string) ([]domain.OrderByEventId, error) {
	SQL := "SELECT orders.id, orders.user_id, orders.event_id, orders.amount, orders.proof_payment, orders.status, orders.created_at, orders.updated_at, users.name FROM orders LEFT JOIN users ON orders.user_id = users.id WHERE orders.event_id = ?"
	rows, err := db.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var orders []domain.OrderByEventId
	for rows.Next() {
		var order domain.OrderByEventId
		var proofPayment *sql.NullString
		err := rows.Scan(&order.Id, &order.UserId, &order.EventId, &order.Amount, &proofPayment, &order.Status, &order.CreatedAt, &order.UpdatedAt, &order.Name)
		if err != nil {
			panic(err)
		}
		if proofPayment.Valid {
			order.ProofPayment = proofPayment.String
		}
		orders = append(orders, order)
	}
	return orders, nil
}
