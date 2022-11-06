package service

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/andil-id/api/exception"
	"github.com/andil-id/api/helper"
	"github.com/andil-id/api/model/domain"
	"github.com/andil-id/api/model/web"
	"github.com/andil-id/api/model/web/moodle"
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
		Id:                pengguna.Id,
		Name:              pengguna.Name,
		ParentName:        pengguna.ParentName,
		Email:             pengguna.Email,
		PhoneNumber:       pengguna.PhoneNumber,
		ParentPhoneNumber: pengguna.ParentPhoneNumber,
		SchoolAddress:     pengguna.SchoolAddress,
		Address:           pengguna.Address,
		CreatedAt:         pengguna.CreatedAt,
		UpdatedAt:         pengguna.UpdatedAt,
	}
	return response_pengguna, nil
}
func (s *UserServiceImpl) RegisterUser(ctx context.Context, user web.RegisterUserRequest) (web.UserResponse, error) {
	registeredUser := web.UserResponse{}

	err := s.Validate.Struct(user)
	if err != nil {
		return registeredUser, err
	}
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	_, err = s.UserRepository.FindUserByEmail(ctx, tx, user.Email)
	if err == nil {
		return registeredUser, e.Wrap(exception.ErrNotFound, "Email has alredy taken for other user")
	}

	sliceName := strings.Split(user.Name, " ")
	firstName := sliceName[0]
	lastName := strings.Join(sliceName[1:], " ")
	reqBody := moodle.UserCreateRequest{
		Users: []moodle.Users{
			{
				Username:  user.Username,
				Password:  user.Password,
				Firstname: firstName,
				Lastname:  lastName,
				Email:     user.Email,
			},
		},
	}
	jsonReq, err := json.Marshal(reqBody)
	if err != nil {
		return registeredUser, err
	}

	baseUrl := os.Getenv("BASE_URL")
	token := os.Getenv("MOODLE_TOKEN")
	client := &http.Client{}
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/core_user_create_users", baseUrl), bytes.NewBuffer(jsonReq))
	if err != nil {
		return registeredUser, err
	}
	r.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Accept":        {"application/json"},
		"Authorization": {token},
	}
	res, err := client.Do(r)
	if err != nil {
		return registeredUser, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println("http statuscode --", res.StatusCode)
		return registeredUser, e.Wrap(exception.ErrService, "error when create user in moodle")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return registeredUser, err
	}
	createdAt := time.Now()
	updatedAt := time.Now()
	passwordHash := string(bytes)
	id, err := s.UserRepository.SaveUser(ctx, tx, domain.Users{
		Name:              user.Name,
		Username:          user.Username,
		Email:             user.Email,
		Password:          passwordHash,
		PhoneNumber:       user.PhoneNumber,
		ParentPhoneNumber: user.ParentPhoneNumber,
		SchoolAddress:     user.SchoolAddress,
		Address:           user.Address,
		ParentName:        user.ParentName,
		CreatedAt:         createdAt,
		UpdatedAt:         updatedAt,
	})
	if err != nil {
		return registeredUser, err
	}
	registeredUser = web.UserResponse{
		Id:                id,
		Name:              user.Name,
		Username:          user.Username,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		Address:           user.Address,
		SchoolAddress:     user.SchoolAddress,
		ParentName:        user.ParentName,
		ParentPhoneNumber: user.ParentPhoneNumber,
		CreatedAt:         createdAt,
		UpdatedAt:         updatedAt,
	}
	return registeredUser, nil
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
			Id:                p.Id,
			Name:              p.Name,
			Email:             p.Email,
			PhoneNumber:       p.PhoneNumber,
			ParentPhoneNumber: p.ParentPhoneNumber,
			Address:           p.Address,
			SchoolAddress:     p.SchoolAddress,
			CreatedAt:         p.CreatedAt,
			UpdatedAt:         p.UpdatedAt,
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
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return err
	}
	passwordHash := string(bytes)
	defer helper.CommitOrRollback(tx)
	s.UserRepository.UpdateProfileUser(ctx, tx, id, domain.Users{
		Name:              request.Name,
		Email:             request.Email,
		Password:          passwordHash,
		PhoneNumber:       request.PhoneNumber,
		ParentPhoneNumber: request.ParentPhoneNumber,
		SchoolAddress:     request.SchoolAddress,
		Address:           request.Address,
		UpdatedAt:         time.Now(),
	})
	return nil
}
