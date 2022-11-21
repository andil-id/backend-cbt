package domain

import "time"

type Orders struct {
	Id           string
	UserId       string
	EventId      string
	Amount       int
	ProofPayment string
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
