package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/andil-id/api/exception"
	"github.com/andil-id/api/helper"
	"github.com/andil-id/api/model/domain"
	"github.com/andil-id/api/model/web"
	"github.com/andil-id/api/repository"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	e "github.com/pkg/errors"
)

type OrderServiceImpl struct {
	DB              *sql.DB
	Validate        *validator.Validate
	OrderRepository repository.OrderRepository
	Cld             *cloudinary.Cloudinary
}

func NewOrderService(db *sql.DB, validate *validator.Validate, orderRepository repository.OrderRepository, cld *cloudinary.Cloudinary) OrderService {
	return &OrderServiceImpl{
		DB:              db,
		Validate:        validate,
		OrderRepository: orderRepository,
		Cld:             cld,
	}
}

func (s *OrderServiceImpl) UpdateOrderStatus(ctx context.Context, status string, id string) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	switch status {
	case "confirm":
		err := s.OrderRepository.UpdateOrderStatus(ctx, tx, "CONFIRM", id)
		if err != nil {
			return err
		}
	case "reject":
		err := s.OrderRepository.UpdateOrderStatus(ctx, tx, "REJECT", id)
		if err != nil {
			return err
		}
	default:
		return errors.New("set order status not allowed: only accept confirm, reject, and waiting")
	}
	return nil
}

func (s *OrderServiceImpl) CreateOrder(ctx context.Context, data web.CreateOrderRequest, userId string) (web.Order, error) {
	res := web.Order{}
	now := time.Now()

	err := s.Validate.Struct(data)
	if err != nil {
		return res, err
	}
	tx, err := s.DB.Begin()
	if err != nil {
		return res, err
	}
	defer helper.CommitOrRollback(tx)

	proofPayment, err := data.ProofPayment.Open()
	if err != nil {
		return res, err
	}
	proofPaymentPath, err := helper.UploadFileToFirebaseStorageAndGetURL(ctx, proofPayment)
	if err != nil {
		return res, err
	}

	order := domain.Orders{
		UserId:       userId,
		EventId:      data.EventId,
		Amount:       data.Amount,
		ProofPayment: proofPaymentPath,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	id, err := s.OrderRepository.SaveOrder(ctx, tx, order)
	if err != nil {
		return res, err
	}

	res = web.Order{
		Id:           id,
		UserId:       userId,
		EventId:      data.EventId,
		Amount:       data.Amount,
		ProofPayment: proofPaymentPath,
		Status:       "WAITING",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	return res, nil
}

func (s *OrderServiceImpl) GetOrderEventByUserId(ctx context.Context, id string) ([]web.OrderByUserId, error) {
	var res []web.OrderByUserId
	orders, err := s.OrderRepository.GetOrderByUserId(ctx, s.DB, id)
	if err != nil {
		return res, err
	}
	for _, order := range orders {
		res = append(res, web.OrderByUserId{
			Id:        order.Id,
			UserId:    order.UserId,
			EventId:   order.EventId,
			Amount:    order.Amount,
			Status:    order.Status,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
			Event: web.Event{
				Id:       order.EventId,
				Title:    order.Title,
				Banner:   order.Banner,
				Location: order.Location,
				StartAt:  order.StartAt,
				EndAt:    order.EndAt,
			},
		})
	}
	return res, nil
}

func (s *OrderServiceImpl) GetOrderByEventId(ctx context.Context, id string) ([]web.Order, error) {
	var res []web.Order
	orders, err := s.OrderRepository.GetOrderByEventId(ctx, s.DB, id)
	if err != nil {
		return res, err
	}
	for _, order := range orders {
		res = append(res, web.Order{
			Id:           order.Id,
			UserId:       order.UserId,
			EventId:      order.EventId,
			Amount:       order.Amount,
			ProofPayment: order.ProofPayment,
			Status:       order.Status,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
		})
	}
	return res, nil
}

func (s *OrderServiceImpl) GetOrderById(ctx context.Context, id string) (web.Order, error) {
	var res web.Order
	order, err := s.OrderRepository.GetOrderById(ctx, s.DB, id)
	if err != nil {
		return res, e.Wrap(exception.ErrNotFound, err.Error())
	}
	res = web.Order{
		Id:           order.Id,
		UserId:       order.UserId,
		EventId:      order.EventId,
		Amount:       order.Amount,
		ProofPayment: order.ProofPayment,
		Status:       order.Status,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}
	return res, nil
}
