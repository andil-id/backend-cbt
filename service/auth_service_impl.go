package service

import (
	"context"
	"database/sql"

	"github.com/andil-id/api/exception"
	"github.com/andil-id/api/helper"
	"github.com/andil-id/api/model/web"
	"github.com/andil-id/api/pkg"
	"github.com/andil-id/api/repository"
	"github.com/go-playground/validator/v10"
	e "github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	UserAuthRepository  repository.UserRepository
	AdminAuthRepository repository.AdminRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewAuthService(DB *sql.DB, validate *validator.Validate, userRepo repository.UserRepository, adminRepo repository.AdminRepository) AuthService {
	return &AuthServiceImpl{
		UserAuthRepository:  userRepo,
		AdminAuthRepository: adminRepo,
		DB:                  DB,
		Validate:            validate,
	}
}
func (service AuthServiceImpl) Login(ctx context.Context, user string, request web.LoginAuthRequest) (string, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return "", err
	}
	tx, err := service.DB.Begin()
	if err != nil {
		return "", err
	}
	defer helper.CommitOrRollback(tx)
	if user == "user" {
		user, err := service.UserAuthRepository.FindUserByEmail(ctx, tx, request.Username)
		if err != nil {
			return "", e.Wrap(exception.ErrBadRequest, "Email not registered")
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		if err != nil {
			return "", e.Wrap(exception.ErrBadRequest, "Password wrong")
		}
		signedToken, err := pkg.GenereateJwtToken(user.Id, user.Name, user.Email, "user")
		if err != nil {
			return "", err
		}
		return signedToken, nil
	} else if user == "admin" {
		admin, err := service.AdminAuthRepository.FindAdminByUsername(ctx, tx, request.Username)
		if err != nil {
			return "", e.Wrap(exception.ErrBadRequest, "Wrong username")
		}
		err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(request.Password))
		if err != nil {
			return "", e.Wrap(exception.ErrBadRequest, "Wrong password")
		}
		signedToken, err := pkg.GenereateJwtToken(admin.Id, admin.Name, admin.Username, "admin")
		if err != nil {
			return "", err
		}
		return signedToken, nil
	} else {
		return "", e.Wrap(exception.ErrBadRequest, "Request param not allowed")
	}
}
