package domain

import (
	"time"
)

type User struct {
	IdUser          string
	NamaUser        string
	NamaOrtu        string
	EmailUser       string
	PasswordUser    string
	NoHandphoneUser string
	NoHandphoneOrtu string
	AlamatSekolah   string
	AlamatUser      string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
