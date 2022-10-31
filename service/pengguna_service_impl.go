package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/andil-id/api/exception"
	"github.com/andil-id/api/helper"
	"github.com/andil-id/api/model/domain"
	"github.com/andil-id/api/model/web"
	"github.com/andil-id/api/repository"
	"github.com/go-playground/validator/v10"
	e "github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(repo repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: repo,
		DB:             DB,
		Validate:       validate,
	}
}

func (s *UserServiceImpl) GetUserById(ctx context.Context, id string) (web.UserResponse, error) {
	var response_pengguna web.UserResponse
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	pengguna, err := s.UserRepository.GetUserById(ctx, tx, id)
	if err != nil {
		return response_pengguna, e.Wrap(exception.ErrNotFound, err.Error())
	}
	response_pengguna = web.UserResponse{
		IdUser:          pengguna.IdUser,
		NamaUser:        pengguna.NamaUser,
		NamaOrtu:        pengguna.NamaOrtu,
		EmailUser:       pengguna.EmailUser,
		NoHandphoneUser: pengguna.NoHandphoneUser,
		NoHandphoneOrtu: pengguna.NoHandphoneOrtu,
		AlamatSekolah:   pengguna.AlamatSekolah,
		AlamatUser:      pengguna.AlamatUser,
		CreatedAt:       pengguna.CreatedAt,
		UpdatedAt:       pengguna.UpdatedAt,
	}
	return response_pengguna, nil
}
func (s *UserServiceImpl) RegisterUser(ctx context.Context, pengguna web.RegisterUserRequest) error {
	err := s.Validate.Struct(pengguna)
	if err != nil {
		return err
	}
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	_, err = s.UserRepository.FindUserByEmail(ctx, tx, pengguna.EmailUser)
	if err == nil {
		return e.Wrap(exception.ErrNotFound, "Email has alredy taken for other user")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(pengguna.PasswordUser), 12)
	if err != nil {
		return err
	}
	passwordHash := string(bytes)
	err = s.UserRepository.SaveUser(ctx, tx, domain.User{
		NamaUser:        pengguna.NamaUser,
		EmailUser:       pengguna.EmailUser,
		PasswordUser:    passwordHash,
		NoHandphoneUser: pengguna.NoHandphoneUser,
		NoHandphoneOrtu: pengguna.NoHandphoneOrtu,
		AlamatSekolah:   pengguna.AlamatSekolah,
		AlamatUser:      pengguna.AlamatUser,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *UserServiceImpl) DeleteUser(ctx context.Context, id string) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)
	s.UserRepository.DeleteUser(ctx, tx, id)
	return nil
}
func (s *UserServiceImpl) GetAllUser(ctx context.Context) ([]web.UserResponse, error) {
	var response_pengguna []web.UserResponse
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	pengguna, err := s.UserRepository.GetAllUser(ctx, tx)
	if err != nil {
		return response_pengguna, err
	}
	for _, p := range pengguna {
		response_pengguna = append(response_pengguna, web.UserResponse{
			IdUser:          p.IdUser,
			NamaUser:        p.NamaUser,
			EmailUser:       p.EmailUser,
			NoHandphoneUser: p.NoHandphoneUser,
			NoHandphoneOrtu: p.NoHandphoneOrtu,
			AlamatUser:      p.AlamatUser,
			AlamatSekolah:   p.AlamatSekolah,
			CreatedAt:       p.CreatedAt,
			UpdatedAt:       p.UpdatedAt,
		})
	}
	return response_pengguna, nil
}
func (s *UserServiceImpl) UpdateProfileUser(ctx context.Context, id string, request web.UpdateProfileUserRequest) error {
	err := s.Validate.Struct(request)
	if err != nil {
		return err
	}
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.PasswordUser), 12)
	if err != nil {
		return err
	}
	passwordHash := string(bytes)
	defer helper.CommitOrRollback(tx)
	s.UserRepository.UpdateProfileUser(ctx, tx, id, domain.User{
		NamaUser:        request.NamaUser,
		EmailUser:       request.EmailUser,
		PasswordUser:    passwordHash,
		NoHandphoneUser: request.NoHandphoneUser,
		NoHandphoneOrtu: request.NoHandphoneOrtu,
		AlamatSekolah:   request.AlamatSekolah,
		AlamatUser:      request.AlamatUser,
		UpdatedAt:       time.Now(),
	})
	return nil
}
