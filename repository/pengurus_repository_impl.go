package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/andil-id/api/model/domain"
	"github.com/segmentio/ksuid"
)

type AdminRepositoryImpl struct {
}

func NewAdminRepository() AdminRepository {
	return &AdminRepositoryImpl{}
}
func (repository *AdminRepositoryImpl) GetAdminById(ctx context.Context, tx *sql.Tx, id string) (domain.Admin, error) {
	SQL := "SELECT * FROM table_admin WHERE id_admin = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	admin := domain.Admin{}
	if rows.Next() {
		err := rows.Scan(&admin.IdAdmin, &admin.NamaAdmin, &admin.UsernameAdmin, &admin.PasswordAdmin, &admin.CreatedAt, &admin.UpdatedAt)
		if err != nil {
			panic(err)
		}
		return admin, nil
	} else {
		return admin, errors.New("data not found")
	}
}
func (repository *AdminRepositoryImpl) SaveAdmin(ctx context.Context, tx *sql.Tx, admin domain.Admin) error {
	SQL := "INSERT INTO `table_admin`(`id_admin`, `nama_admin`, `username`, `password`, `created_at`, `updated_at`) VALUES ('?','?','?','?','?','')"
	id := ksuid.New().String()
	_, err := tx.ExecContext(ctx, SQL, id, admin.NamaAdmin, admin.UsernameAdmin, admin.PasswordAdmin, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}
func (repository *AdminRepositoryImpl) DeleteAdmin(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "DELETE FROM `table_admin` WHERE id_admin = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}
	return nil
}
func (repository *AdminRepositoryImpl) GetAllAdmin(ctx context.Context, tx *sql.Tx) ([]domain.Admin, error) {
	SQL := "SELECT * FROM `table_admin` LIMIT 10"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	allAdmin := []domain.Admin{}
	for rows.Next() {
		admin := domain.Admin{}
		err := rows.Scan(&admin.IdAdmin, &admin.NamaAdmin, &admin.UsernameAdmin, &admin.PasswordAdmin, &admin.CreatedAt, &admin.UpdatedAt)
		if err != nil {
			panic(err)
		}
		allAdmin = append(allAdmin, admin)
	}
	return allAdmin, nil
}
func (repository *AdminRepositoryImpl) FindAdminByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.Admin, error) {
	SQL := "SELECT * FROM `table_admin` WHERE username = ?"
	rows, err := tx.QueryContext(ctx, SQL, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var admin domain.Admin
	if rows.Next() {
		err := rows.Scan(&admin.IdAdmin, &admin.NamaAdmin, &admin.UsernameAdmin, &admin.PasswordAdmin, &admin.CreatedAt, &admin.UpdatedAt)
		if err != nil {
			panic(err)
		}
		return admin, nil
	} else {
		return admin, errors.New("data not found")
	}
}
func (repository *AdminRepositoryImpl) UpdatePasswordAdmin(ctx context.Context, tx *sql.Tx, username string, password string) error {
	SQL := "UPDATE `table_admin` SET `password`='?', `updated_at`='?' WHERE username = ?"
	_, err := tx.ExecContext(ctx, SQL, password, time.Now(), username)
	if err != nil {
		return err
	}
	return nil
}
func (repository *AdminRepositoryImpl) UpdateProfileAdmin(ctx context.Context, tx *sql.Tx, id string, admin domain.Admin) error {
	SQL := "UPDATE `table_admin` SET `nama_admin`='?',`username`='?',`password`='?', `updated_at`='?' WHERE id_admin = ?"
	_, err := tx.ExecContext(ctx, SQL, admin.NamaAdmin, admin.UsernameAdmin, admin.PasswordAdmin, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}
