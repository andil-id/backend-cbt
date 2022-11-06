package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/andil-id/api/model/domain"
	"github.com/segmentio/ksuid"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) GetUserById(ctx context.Context, tx *sql.Tx, id string) (domain.Users, error) {
	SQL := "SELECT * FROM users WHERE id_user = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var user domain.Users
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.ParentName, &user.Email, &user.Password, &user.PhoneNumber, &user.ParentPhoneNumber, &user.Address, &user.SchoolAddress, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {
		return user, errors.New("data not found")
	}
}
func (repo *UserRepositoryImpl) SaveUser(ctx context.Context, tx *sql.Tx, user domain.Users) (string, error) {
	SQL := "INSERT INTO users (id, name, username, phone_number, address, school_address, email, password, parent_name, parent_phone_number, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
	id := ksuid.New().String()
	log.Println("ini parent name", user.ParentName)
	_, err := tx.ExecContext(ctx, SQL, id, user.Name, user.Username, user.PhoneNumber, user.Address, user.Address, user.Email, user.Password, user.ParentName, user.ParentPhoneNumber, time.Now(), time.Now())
	if err != nil {
		return "", err
	}
	return id, nil
}
func (repo *UserRepositoryImpl) DeleteUser(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}
	return nil
}
func (repo *UserRepositoryImpl) GetAllUser(ctx context.Context, tx *sql.Tx) ([]domain.Users, error) {
	SQL := "SELECT * FROM users LIMIT 10"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var user []domain.Users
	for rows.Next() {
		var u domain.Users
		err := rows.Scan(&u.Id, &u.Name, &u.ParentName, &u.Email, &u.PhoneNumber, &u.ParentPhoneNumber, &u.Address, &u.SchoolAddress, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			panic(err)
		}
		user = append(user, u)
	}
	return user, nil
}
func (repo *UserRepositoryImpl) FindUserByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.Users, error) {
	SQL := "SELECT * FROM users WHERE email = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, email)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var user domain.Users
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.ParentName, &user.Email, &user.PhoneNumber, &user.ParentPhoneNumber, &user.Address, &user.SchoolAddress, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {
		return user, errors.New("data not found")
	}
}
func (repo *UserRepositoryImpl) UpdateProfileUser(ctx context.Context, tx *sql.Tx, id string, user domain.Users) error {
	SQL := "UPDATE users SET name=?,phone_number=?,address=?,school_address=?,email=?,parent_name=?,parent_phone_number=?,updated_at=? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.PhoneNumber, user.Address, user.SchoolAddress, user.Email, user.ParentName, user.ParentPhoneNumber, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}
func (repo *UserRepositoryImpl) UpdatePasswordUser(ctx context.Context, tx *sql.Tx, email string, pasword string) error {
	SQL := "UPDATE users SET password = ? WHERE email = ?"
	_, err := tx.ExecContext(ctx, SQL, pasword, email)
	if err != nil {
		return err
	}
	return nil
}
