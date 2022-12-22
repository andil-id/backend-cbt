package service

import (
	"context"

	"github.com/andil-id/api/model/web"
)

type OrderService interface {
	// GetAllOrder(ctx context.Context) (error, []web.Order)
	GetOrderById(ctx context.Context, id string) (web.Order, error)
	CreateOrder(ctx context.Context, data web.CreateOrderRequest, userId string) (web.Order, error)
	// GetOrderUserByEventId(ctx context.Context, userId string, eventId string) (error, web.Order)
	UpdateOrderStatus(ctx context.Context, status string, id string) error
	GetOrderEventByUserId(ctx context.Context, id string) ([]web.OrderByUserId, error)
	GetOrderByEventId(ctx context.Context, id string) ([]web.OrderByEventId, error)
}
