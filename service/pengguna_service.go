package service

import (
	"context"

	"github.com/andil-id/api/model/web"
)

type UserService interface {
	RegisterUser(ctx context.Context, user web.RegisterUserRequest) error
	GetUserById(ctx context.Context, id string) (web.UserResponse, error)
	GetAllUser(ctx context.Context) ([]web.UserResponse, error)
	UpdateProfileUser(ctx context.Context, id string, user web.UpdateProfileUserRequest) error
	DeleteUser(ctx context.Context, id string) error
}
