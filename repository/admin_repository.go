package repository

import (
	"context"
	"database/sql"

	"github.com/andil-id/api/model/domain"
)

type AdminRepository interface {
	GetAdminById(ctx context.Context, tx *sql.Tx, id string) (domain.Admins, error)
	SaveAdmin(ctx context.Context, tx *sql.Tx, admin domain.Admins) (string, error)
	DeleteAdmin(ctx context.Context, tx *sql.Tx, id string) error
	GetAllAdmin(ctx context.Context, tx *sql.Tx) ([]domain.Admins, error)
	FindAdminByUsername(ctx context.Context, tx *sql.Tx, email string) (domain.Admins, error)
	UpdatePasswordAdmin(ctx context.Context, tx *sql.Tx, email string, password string) error
	UpdateProfileAdmin(ctx context.Context, tx *sql.Tx, id string, admin domain.Admins) error
}
