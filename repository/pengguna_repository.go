package repository

import (
	"context"
	"database/sql"

	"github.com/andil-id/api/model/domain"
)

type UserRepository interface {
	GetUserById(ctx context.Context, tx *sql.Tx, id string) (domain.User, error)
	SaveUser(ctx context.Context, tx *sql.Tx, pengguna domain.User) error
	DeleteUser(ctx context.Context, tx *sql.Tx, id string) error
	GetAllUser(ctx context.Context, tx *sql.Tx) ([]domain.User, error)
	FindUserByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	UpdatePasswordUser(ctx context.Context, tx *sql.Tx, email string, password string) error
	UpdateProfileUser(ctx context.Context, tx *sql.Tx, id string, pengguna domain.User) error
}
