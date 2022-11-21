package domain

import "time"

type Events struct {
	Id             string
	Title          string
	Description    string
	Banner         string
	Certificate    string
	Price          int
	Type           string
	BankAccountNum string
	Location       string
	StartAt        time.Time
	EndAt          time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
