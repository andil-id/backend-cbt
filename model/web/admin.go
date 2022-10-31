package web

import "time"

type RegisterAdminRequest struct {
	NamaAdmin     string    `json:"nama_admin"`
	UsernameAdmin string    `json:"username_admin"`
	PasswordAdmin string    `json:"password_admin"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type GetAdminResponse struct {
	IdAdmin       string    `json:"id_admin"`
	NamaAdmin     string    `json:"nama_admin"`
	UsernameAdmin string    `json:"username_admin"`
	PasswordAdmin string    `json:"password_admin"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
type UpdateProfileAdminRequest struct {
	NamaAdmin     string `json:"nama_admin"`
	UsernameAdmin string `json:"username_admin"`
	PasswordAdmin string `json:"password_admin"`
}
