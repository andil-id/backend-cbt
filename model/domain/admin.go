package domain

import (
	"time"
)

type Admin struct {
	IdAdmin       string
	NamaAdmin     string
	UsernameAdmin string
	PasswordAdmin string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
