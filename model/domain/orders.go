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

type OrderEventByUser struct {
	Id        string
	UserId    string
	EventId   string
	Amount    int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Banner    string
	Location  string
	StartAt   time.Time
	EndAt     time.Time
}
