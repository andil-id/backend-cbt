package repository

import (
	"context"
	"database/sql"

	"github.com/andil-id/api/model/domain"
)

type UserRepository interface {
	GetUserById(ctx context.Context, tx *sql.Tx, id string) (domain.Users, error)
	SaveUser(ctx context.Context, tx *sql.Tx, pengguna domain.Users) (string, error)
	DeleteUser(ctx context.Context, tx *sql.Tx, id string) error
	GetAllUser(ctx context.Context, tx *sql.Tx) ([]domain.Users, error)
	FindUserByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.Users, error)
	UpdatePasswordUser(ctx context.Context, tx *sql.Tx, email string, password string) error
	UpdateProfileUser(ctx context.Context, tx *sql.Tx, id string, pengguna domain.Users) error
}
