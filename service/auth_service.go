package service

import (
	"context"

	"github.com/andil-id/api/model/web"
)

type AuthService interface {
	Login(ctx context.Context, user string, request web.LoginAuthRequest) (string, error)
}
