package service

import (
	"context"

	"github.com/andil-id/api/model/web"
)

type AuthService interface {
	LoginUser(ctx context.Context, request web.LoginUserRequest) (string, error)
	LoginAdmin(ctx context.Context, request web.LoginAdminRequest) (string, error)
}
