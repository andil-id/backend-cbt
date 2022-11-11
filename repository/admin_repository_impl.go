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
func (repository *AdminRepositoryImpl) GetAdminById(ctx context.Context, tx *sql.Tx, id string) (domain.Admins, error) {
	SQL := "SELECT * FROM admins WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	admin := domain.Admins{}
	if rows.Next() {
		err := rows.Scan(&admin.Id, &admin.Name, &admin.Username, &admin.Password, &admin.CreatedAt, &admin.UpdatedAt)
		if err != nil {
			panic(err)
		}
		return admin, nil
	} else {
		return admin, errors.New("data not found")
	}
}
func (repository *AdminRepositoryImpl) SaveAdmin(ctx context.Context, tx *sql.Tx, admin domain.Admins) error {
	SQL := "INSERT INTO admins (id, name, username, password, created_at, updated_at) VALUES (?,?,?,?,?,?)"
	id := ksuid.New().String()
	_, err := tx.ExecContext(ctx, SQL, id, admin.Name, admin.Username, admin.Password, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}
func (repository *AdminRepositoryImpl) DeleteAdmin(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "DELETE FROM admins WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}
	return nil
}
func (repository *AdminRepositoryImpl) GetAllAdmin(ctx context.Context, tx *sql.Tx) ([]domain.Admins, error) {
	SQL := "SELECT * FROM admins LIMIT 10"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	allAdmin := []domain.Admins{}
	for rows.Next() {
		admin := domain.Admins{}
		err := rows.Scan(&admin.Id, &admin.Name, &admin.Username, &admin.Password, &admin.CreatedAt, &admin.UpdatedAt)
		if err != nil {
			panic(err)
		}
		allAdmin = append(allAdmin, admin)
	}
	return allAdmin, nil
}
func (repository *AdminRepositoryImpl) FindAdminByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.Admins, error) {
	SQL := "SELECT * FROM admins WHERE username = ?"
	rows, err := tx.QueryContext(ctx, SQL, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var admin domain.Admins
	if rows.Next() {
		err := rows.Scan(&admin.Id, &admin.Name, &admin.Username, &admin.Password, &admin.CreatedAt, &admin.UpdatedAt)
		if err != nil {
			panic(err)
		}
		return admin, nil
	} else {
		return admin, errors.New("data not found")
	}
}
func (repository *AdminRepositoryImpl) UpdatePasswordAdmin(ctx context.Context, tx *sql.Tx, username string, password string) error {
	SQL := "UPDATE admins SET password=?, updated_at=? WHERE username = ?"
	_, err := tx.ExecContext(ctx, SQL, password, time.Now(), username)
	if err != nil {
		return err
	}
	return nil
}
func (repository *AdminRepositoryImpl) UpdateProfileAdmin(ctx context.Context, tx *sql.Tx, id string, admin domain.Admins) error {
	SQL := "UPDATE admins SET name=?,username=?,password=?, updated_at=? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, admin.Name, admin.Username, admin.Password, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}
