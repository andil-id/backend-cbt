package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/andil-id/api/helper"
	"github.com/andil-id/api/model/domain"
	"github.com/andil-id/api/model/web"
	"github.com/andil-id/api/repository"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
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
	proofPaymentPath, err := helper.ImageUploader(ctx, s.Cld, proofPayment, "payment")
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
