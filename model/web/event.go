package web

import (
	"mime/multipart"
	"time"
)

type Event struct {
	Id          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Banner      string    `json:"banner,omitempty"`
	StartAt     time.Time `json:"start_at,omitempty"`
	EndAt       time.Time `json:"end_at,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type CreateEventRequest struct {
	Title       string                `form:"title" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Banner      *multipart.FileHeader `form:"banner" binding:"required"`
	StartAt     time.Time             `form:"start_at" binding:"required"`
	EndAt       time.Time             `form:"end_at" binding:"required"`
}
