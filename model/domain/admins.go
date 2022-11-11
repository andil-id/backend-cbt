package domain

import (
	"time"
)

type Admins struct {
	Id        string
	Name      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
