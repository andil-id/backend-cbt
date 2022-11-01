package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/andil-id/api/model/domain"
	"github.com/segmentio/ksuid"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) GetUserById(ctx context.Context, tx *sql.Tx, id string) (domain.User, error) {
	SQL := "SELECT * FROM table_user WHERE id_user = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var user domain.User
	if rows.Next() {
		err := rows.Scan(&user.IdUser, &user.NamaUser, &user.NamaOrtu, &user.EmailUser, &user.PasswordUser, &user.NoHandphoneUser, &user.NoHandphoneOrtu, &user.AlamatUser, &user.AlamatSekolah, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {
		return user, errors.New("data not found")
	}
}
func (repo *UserRepositoryImpl) SaveUser(ctx context.Context, tx *sql.Tx, user domain.User) error {
	SQL := "INSERT INTO table_user (id_user, nama_user, no_handphone, alamat_asal, alamat_sekolah, email_user, password_user, nama_ortu, no_hp_ortu, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	id := ksuid.New().String()
	_, err := tx.ExecContext(ctx, SQL, id, user.NamaUser, user.NoHandphoneUser, user.AlamatUser, user.AlamatSekolah, user.EmailUser, user.PasswordUser, user.NamaOrtu, user.NoHandphoneOrtu, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}
func (repo *UserRepositoryImpl) DeleteUser(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "DELETE FROM table_user WHERE id_user = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}
	return nil
}
func (repo *UserRepositoryImpl) GetAllUser(ctx context.Context, tx *sql.Tx) ([]domain.User, error) {
	SQL := "SELECT * FROM table_user LIMIT 10"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var user []domain.User
	for rows.Next() {
		var u domain.User
		err := rows.Scan(&u.IdUser, &u.NamaUser, &u.NamaOrtu, &u.EmailUser, &u.NoHandphoneUser, &u.NoHandphoneOrtu, &u.AlamatUser, &u.AlamatSekolah, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			panic(err)
		}
		user = append(user, u)
	}
	return user, nil
}
func (repo *UserRepositoryImpl) FindUserByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "SELECT * FROM table_user WHERE email_user = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, email)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var user domain.User
	if rows.Next() {
		err := rows.Scan(&user.IdUser, &user.NamaUser, &user.NamaOrtu, &user.EmailUser, &user.NoHandphoneUser, &user.NoHandphoneOrtu, &user.AlamatUser, &user.AlamatSekolah, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {
		return user, errors.New("data not found")
	}
}
func (repo *UserRepositoryImpl) UpdateProfileUser(ctx context.Context, tx *sql.Tx, id string, user domain.User) error {
	SQL := "UPDATE table_user SET nama_user=?,no_handphone=?,alamat_asal=?,alamat_sekolah=?,email_user=?,nama_ortu=?,no_hp_ortu=?,updated_at=? WHERE id_user = ?"
	_, err := tx.ExecContext(ctx, SQL, user.NamaUser, user.NoHandphoneUser, user.AlamatUser, user.AlamatSekolah, user.EmailUser, user.NamaOrtu, user.NoHandphoneOrtu, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}
func (repo *UserRepositoryImpl) UpdatePasswordUser(ctx context.Context, tx *sql.Tx, email string, pasword string) error {
	SQL := "UPDATE table_user SET password_user = ? WHERE email_user = ?"
	_, err := tx.ExecContext(ctx, SQL, pasword, email)
	if err != nil {
		return err
	}
	return nil
}
