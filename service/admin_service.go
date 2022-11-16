package service

import (
	"context"

	"github.com/andil-id/api/model/web"
)

type AdminService interface {
	RegisterAdmin(ctx context.Context, request web.RegisterAdminRequest) (web.Admin, error)
	GetAdminById(ctx context.Context, id string) (web.GetAdminResponse, error)
	GetAllAdmin(ctx context.Context) ([]web.GetAdminResponse, error)
	UpdateProfileAdmin(ctx context.Context, id string, request web.UpdateProfileAdminRequest) error
	DeleteAdmin(ctx context.Context, id string) error
}
