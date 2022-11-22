package service

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/andil-id/api/config"
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
	var res web.UserResponse
	user, err := s.UserRepository.GetUserById(ctx, s.DB, id)
	if err != nil {
		return res, e.Wrap(exception.ErrNotFound, err.Error())
	}
	res = web.UserResponse{
		Id:                user.Id,
		Username:          user.Username,
		Name:              user.Name,
		ParentName:        user.ParentName,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		ParentPhoneNumber: user.ParentPhoneNumber,
		SchoolAddress:     user.SchoolAddress,
		Address:           user.Address,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
	}
	return res, nil
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

	byte, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return registeredUser, err
	}
	createdAt := time.Now()
	updatedAt := time.Now()
	passwordHash := string(byte)
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

	baseUrl := config.MoodleBaseUrl()
	token := config.MoodleToken()
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
		var failureResponse interface{}
		err := json.NewDecoder(res.Body).Decode(&failureResponse)
		if err != nil {
			return registeredUser, err
		}
		log.Println(failureResponse)
		log.Println("http statuscode --", res.StatusCode)
		err = s.UserRepository.DeleteUser(ctx, tx, id)
		if err != nil {
			return registeredUser, nil
		}
		return registeredUser, e.Wrap(exception.ErrService, "error when create user in moodle")
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
	err = s.UserRepository.DeleteUser(ctx, tx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) GetAllUser(ctx context.Context) ([]web.UserResponse, error) {
	var res []web.UserResponse
	users, err := s.UserRepository.GetAllUser(ctx, s.DB)
	if err != nil {
		return res, err
	}
	for _, user := range users {
		res = append(res, web.UserResponse{
			Id:                user.Id,
			Name:              user.Name,
			Email:             user.Email,
			PhoneNumber:       user.PhoneNumber,
			ParentPhoneNumber: user.ParentPhoneNumber,
			Address:           user.Address,
			SchoolAddress:     user.SchoolAddress,
			CreatedAt:         user.CreatedAt,
			UpdatedAt:         user.UpdatedAt,
		})
	}
	return res, nil
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
