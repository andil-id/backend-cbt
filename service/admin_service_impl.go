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
	_ "github.com/joho/godotenv/autoload"
	e "github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AdminServiceImpl struct {
	AdminRepository repository.AdminRepository
	UserRepository  repository.UserRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewAdminService(adminRepository repository.AdminRepository, db *sql.DB, validator *validator.Validate, penggunaRepository repository.UserRepository) AdminService {
	return &AdminServiceImpl{
		AdminRepository: adminRepository,
		DB:              db,
		Validate:        validator,
		UserRepository:  penggunaRepository,
	}
}

func (service *AdminServiceImpl) RegisterAdmin(ctx context.Context, request web.RegisterAdminRequest) (web.Admin, error) {
	var registeredAdmin web.Admin

	// * validate request body
	err := service.Validate.Struct(request)
	if err != nil {
		return registeredAdmin, err
	}
	// * rolback db transaction when error
	tx, err := service.DB.Begin()
	if err != nil {
		return registeredAdmin, err
	}
	defer helper.CommitOrRollback(tx)
	// * check avaliable email
	_, err = service.AdminRepository.FindAdminByUsername(ctx, tx, request.Username)
	if err == nil {
		return registeredAdmin, e.Wrap(exception.ErrNotFound, "Username has alredy taken for other user")
	}
	// * generate password hash
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return registeredAdmin, err
	}
	passwordHash := string(bytes)
	now := time.Now()
	// * save admin to db
	admin := domain.Admins{
		Name:      request.Name,
		Username:  request.Username,
		Password:  passwordHash,
		CreatedAt: now,
		UpdatedAt: now,
	}
	id, err := service.AdminRepository.SaveAdmin(ctx, tx, admin)
	if err != nil {
		return registeredAdmin, err
	}

	registeredAdmin = web.Admin{
		Id:        id,
		Username:  request.Username,
		Name:      request.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return registeredAdmin, nil
}
func (service *AdminServiceImpl) GetAdminById(ctx context.Context, idAdmin string) (web.GetAdminResponse, error) {
	var adminResponse web.GetAdminResponse
	tx, err := service.DB.Begin()
	if err != nil {
		return adminResponse, err
	}
	defer helper.CommitOrRollback(tx)
	admin, err := service.AdminRepository.GetAdminById(ctx, tx, idAdmin)
	if err != nil {
		return adminResponse, e.Wrap(exception.ErrNotFound, "Admin is not found")
	}
	adminResponse = web.GetAdminResponse{
		Id:        admin.Id,
		Name:      admin.Name,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}

	return adminResponse, nil
}
func (service *AdminServiceImpl) GetAllAdmin(ctx context.Context) ([]web.GetAdminResponse, error) {
	var adminResponse []web.GetAdminResponse
	tx, err := service.DB.Begin()
	if err != nil {
		return adminResponse, err
	}
	defer helper.CommitOrRollback(tx)
	admin, err := service.AdminRepository.GetAllAdmin(ctx, tx)
	if err != nil {
		return adminResponse, err
	}
	for _, data := range admin {
		toAdminResponse := web.GetAdminResponse{
			Id:        data.Id,
			Name:      data.Name,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		}
		adminResponse = append(adminResponse, toAdminResponse)
	}
	return adminResponse, nil
}
func (service *AdminServiceImpl) UpdateProfileAdmin(ctx context.Context, id string, request web.UpdateProfileAdminRequest) error {
	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)
	err = service.AdminRepository.UpdateProfileAdmin(ctx, tx, id, domain.Admins{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		return err
	}
	return nil
}
func (service *AdminServiceImpl) DeleteAdmin(ctx context.Context, id string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)
	service.AdminRepository.DeleteAdmin(ctx, tx, id)
	return nil
}
