package web

import (
	"time"
)

type RegisterUserRequest struct {
	NamaUser        string `json:"nama_user" binding:"required"`
	NamaOrtu        string `json:"nama_ortu" binding:"required"`
	EmailUser       string `json:"email_user" binding:"required"`
	PasswordUser    string `json:"password_user" binding:"required"`
	NoHandphoneUser string `json:"no_hp" binding:"required"`
	NoHandphoneOrtu string `json:"no_hp_ortu" binding:"required"`
	AlamatSekolah   string `json:"alamat_sekolah" binding:"required"`
	AlamatUser      string `json:"alamat_user" binding:"required"`
}

type UserResponse struct {
	IdUser          string    `json:"id_user"`
	NamaUser        string    `json:"nama_user"`
	NamaOrtu        string    `json:"nama_ortu"`
	EmailUser       string    `json:"email_user"`
	NoHandphoneUser string    `json:"no_hp"`
	NoHandphoneOrtu string    `json:"no_hp_ortu"`
	AlamatSekolah   string    `json:"alamat_sekolah"`
	AlamatUser      string    `json:"alamat_user"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
type UpdateProfileUserRequest struct {
	NamaUser        string `json:"nama_user" binding:"required"`
	NamaOrtu        string `json:"nama_ortu" binding:"required"`
	EmailUser       string `json:"email_user" binding:"required"`
	NoHandphoneUser string `json:"no_hp" binding:"required"`
	PasswordUser    string `json:"password_user" binding:"required"`
	NoHandphoneOrtu string `json:"no_hp_ortu" binding:"required"`
	AlamatSekolah   string `json:"alamat_sekolah" binding:"required"`
	AlamatUser      string `json:"alamat_user" binding:"required"`
}
