package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/andil-id/api/exception"
	"github.com/andil-id/api/helper"
	"github.com/andil-id/api/model/domain"
	"github.com/andil-id/api/model/web"
	"github.com/andil-id/api/pkg"
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

func (service *AdminServiceImpl) RegisterAdmin(ctx context.Context, request web.RegisterAdminRequest) error {
	// * validate request body
	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}
	// * rolback db transaction when error
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)
	// * check avaliable email
	_, err = service.AdminRepository.FindAdminByUsername(ctx, tx, request.UsernameAdmin)
	if err == nil {
		return e.Wrap(exception.ErrNotFound, "Email has alredy taken for other user")
	}
	// * generate password hash
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.PasswordAdmin), 12)
	if err != nil {
		return err
	}
	passwordHash := string(bytes)
	// * generate special code
	uuid := pkg.SpecialKode()
	// * save admin to db
	admin := domain.Admin{
		IdAdmin:       uuid,
		NamaAdmin:     request.NamaAdmin,
		UsernameAdmin: request.UsernameAdmin,
		PasswordAdmin: passwordHash,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	err = service.AdminRepository.SaveAdmin(ctx, tx, admin)
	if err != nil {
		return err
	}
	return nil
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
		IdAdmin:   admin.IdAdmin,
		NamaAdmin: admin.NamaAdmin,
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
			IdAdmin:   data.IdAdmin,
			NamaAdmin: data.NamaAdmin,
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
	err = service.AdminRepository.UpdateProfileAdmin(ctx, tx, id, domain.Admin{
		NamaAdmin:     request.NamaAdmin,
		UsernameAdmin: request.UsernameAdmin,
		PasswordAdmin: request.PasswordAdmin,
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
