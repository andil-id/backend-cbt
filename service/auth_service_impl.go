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
		err = bcrypt.CompareHashAndPassword([]byte(admin.PasswordAdmin), []byte(request.Password))
		if err != nil {
			return "", e.Wrap(exception.ErrBadRequest, "Wrong password")
		}
		signedToken, err := pkg.GenereateJwtToken(admin.IdAdmin, admin.NamaAdmin, admin.UsernameAdmin, "admin")
		if err != nil {
			return "", err
		}
		return signedToken, nil
	} else {
		return "", e.Wrap(exception.ErrBadRequest, "Request param not allowed")
	}
}
func (service AuthServiceImpl) ForgetPassword(ctx context.Context, user string, request web.ForgetPasswordAuthRequest) error {
	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)
	if user == "user" {
		user, err := service.UserAuthRepository.FindUserByEmail(ctx, tx, request.Username)
		if err != nil {
			return e.Wrap(exception.ErrBadRequest, "Email not registered")
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.NewPassword))
		if err == nil {
			return e.Wrap(exception.ErrBadRequest, "New password can't be same as the old password")
		}
		bytes, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), 12)
		if err != nil {
			return err
		}
		passwordHash := string(bytes)
		service.UserAuthRepository.UpdatePasswordUser(ctx, tx, request.Username, passwordHash)
		return nil
	} else if user == "admin" {
		admin, err := service.AdminAuthRepository.FindAdminByUsername(ctx, tx, request.Username)
		if err != nil {
			return e.Wrap(exception.ErrBadRequest, "Username not registered")
		}
		err = bcrypt.CompareHashAndPassword([]byte(admin.PasswordAdmin), []byte(request.NewPassword))
		if err == nil {
			return e.Wrap(exception.ErrBadRequest, "New password can't be same as the old password")
		}
		bytes, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), 12)
		if err != nil {
			return err
		}
		passwordHash := string(bytes)
		err = service.AdminAuthRepository.UpdatePasswordAdmin(ctx, tx, request.Username, passwordHash)
		if err != nil {
			return err
		}
		return nil
	} else {
		return e.Wrap(exception.ErrBadRequest, "Request param not allowed")
	}
}
func (service AuthServiceImpl) ChangePassowrd(ctx context.Context, user string, username string, request web.ChangePasswordAuthRequest) error {
	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	if user == "user" {
		user, err := service.UserAuthRepository.FindUserByEmail(ctx, tx, username)
		if err != nil {
			return err
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword))
		if err != nil {
			return e.Wrap(exception.ErrBadRequest, "Wrong old password")
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.NewPassword))
		if err == nil {
			return e.Wrap(exception.ErrBadRequest, "New password can't be same as the old password")
		}
		bytes, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), 12)
		if err != nil {
			return err
		}
		passwordHash := string(bytes)
		err = service.UserAuthRepository.UpdatePasswordUser(ctx, tx, username, passwordHash)
		if err != nil {
			return err
		}
		return nil
	} else if user == "admin" {
		admin, err := service.AdminAuthRepository.FindAdminByUsername(ctx, tx, username)
		if err != nil {
			return err
		}
		// * check old password
		err = bcrypt.CompareHashAndPassword([]byte(admin.PasswordAdmin), []byte(request.OldPassword))
		if err != nil {
			return e.Wrap(exception.ErrBadRequest, "Wrong old password")
		}
		// * check new password
		err = bcrypt.CompareHashAndPassword([]byte(admin.PasswordAdmin), []byte(request.NewPassword))
		if err == nil {
			return e.Wrap(exception.ErrBadRequest, "New password can't be same as the old password")
		}
		// * generate new password
		bytes, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), 12)
		if err != nil {
			return err
		}
		passwordHash := string(bytes)
		// * update password
		err = service.AdminAuthRepository.UpdatePasswordAdmin(ctx, tx, username, passwordHash)
		if err != nil {
			return err
		}
		return nil
	} else {
		return e.Wrap(exception.ErrBadRequest, "Request param not allowed")
	}
}
