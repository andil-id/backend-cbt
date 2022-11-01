package repository

import (
	"context"
	"database/sql"

	"github.com/andil-id/api/model/domain"
)

type AdminRepository interface {
	GetAdminById(ctx context.Context, tx *sql.Tx, id string) (domain.Admin, error)
	SaveAdmin(ctx context.Context, tx *sql.Tx, admin domain.Admin) error
	DeleteAdmin(ctx context.Context, tx *sql.Tx, id string) error
	GetAllAdmin(ctx context.Context, tx *sql.Tx) ([]domain.Admin, error)
	FindAdminByUsername(ctx context.Context, tx *sql.Tx, email string) (domain.Admin, error)
	UpdatePasswordAdmin(ctx context.Context, tx *sql.Tx, email string, password string) error
	UpdateProfileAdmin(ctx context.Context, tx *sql.Tx, id string, admin domain.Admin) error
}
