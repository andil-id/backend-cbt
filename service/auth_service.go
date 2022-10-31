package service

import (
	"context"

	"github.com/andil-id/api/model/web"
)

type AuthService interface {
	Login(ctx context.Context, user string, request web.LoginAuthRequest) (string, error)
	ForgetPassword(ctx context.Context, user string, request web.ForgetPasswordAuthRequest) error
	ChangePassowrd(ctx context.Context, user string, email string, request web.ChangePasswordAuthRequest) error
}
