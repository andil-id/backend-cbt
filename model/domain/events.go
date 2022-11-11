package domain

import "time"

type Events struct {
	Id          string
	Title       string
	Description string
	Banner      string
	StartAt     time.Time
	EndAt       time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
