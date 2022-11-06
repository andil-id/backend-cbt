package web

import (
	"time"
)

type RegisterUserRequest struct {
	Name              string `json:"name" binding:"required"`
	Username          string `json:"username" binding:"required"`
	ParentName        string `json:"parent_name" binding:"required"`
	Email             string `json:"email" binding:"required"`
	Password          string `json:"password" binding:"required"`
	PhoneNumber       string `json:"phone_number" binding:"required"`
	ParentPhoneNumber string `json:"parent_phone_number" binding:"required"`
	SchoolAddress     string `json:"school_address" binding:"required"`
	Address           string `json:"address" binding:"required"`
}

type UserResponse struct {
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	Username          string    `json:"username"`
	ParentName        string    `json:"parent_name"`
	Email             string    `json:"email"`
	PhoneNumber       string    `json:"phone_number"`
	ParentPhoneNumber string    `json:"parent_phone_number"`
	SchoolAddress     string    `json:"school_address"`
	Address           string    `json:"address"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type UpdateProfileUserRequest struct {
	Name              string `json:"name" binding:"required"`
	ParentName        string `json:"parent_name" binding:"required"`
	Email             string `json:"email" binding:"required"`
	PhoneNumber       string `json:"phone_number" binding:"required"`
	Password          string `json:"password" binding:"required"`
	ParentPhoneNumber string `json:"parent_phone_number" binding:"required"`
	SchoolAddress     string `json:"school_address" binding:"required"`
	Address           string `json:"address" binding:"required"`
}
