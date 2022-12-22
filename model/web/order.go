package web

import (
	"mime/multipart"
	"time"
)

type Order struct {
	Id           string    `json:"id,omitempty"`
	UserId       string    `json:"user_id,omitempty"`
	EventId      string    `json:"event_id,omitempty"`
	Amount       int       `json:"amount,omitempty"`
	ProofPayment string    `json:"proof_of_payment,omitempty"`
	Status       string    `json:"status,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

type CreateOrderRequest struct {
	EventId      string                `form:"event_id" binding:"required"`
	Amount       int                   `form:"amount" binding:"required"`
	ProofPayment *multipart.FileHeader `form:"proof_of_payment" binding:"required"`
}

type OrderByUserId struct {
	Id        string    `json:"id,omitempty"`
	UserId    string    `json:"user_id,omitempty"`
	EventId   string    `json:"event_id,omitempty"`
	Amount    int       `json:"amount,omitempty"`
	Status    string    `json:"status,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Event     Event     `json:"event"`
}

type OrderByEventIdData struct {
	UserId string `json:"id"`
	Name   string `json:"name"`
}
type OrderByEventId struct {
	Id           string             `json:"id"`
	EventId      string             `json:"event_id"`
	Amount       int                `json:"amount"`
	ProofPayment string             `json:"proof_of_payment"`
	Status       string             `json:"status"`
	User         OrderByEventIdData `json:"user"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}
