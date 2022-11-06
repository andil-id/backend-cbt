package domain

import (
	"time"
)

type Users struct {
	Id                string
	Username          string
	Name              string
	ParentName        string
	Email             string
	Password          string
	PhoneNumber       string
	ParentPhoneNumber string
	SchoolAddress     string
	Address           string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
