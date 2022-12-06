package service

import (
	"context"
	"database/sql"
	"path/filepath"
	"strings"
	"time"

	"github.com/andil-id/api/exception"
	"github.com/andil-id/api/helper"
	"github.com/andil-id/api/model/domain"
	"github.com/andil-id/api/model/web"
	"github.com/andil-id/api/repository"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	e "github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

type EventServiceImpl struct {
	DB              *sql.DB
	Validate        *validator.Validate
	EventRepository repository.EventRepository
	Cld             *cloudinary.Cloudinary
}

func NewEventService(db *sql.DB, validate *validator.Validate, eventRepository repository.EventRepository, cld *cloudinary.Cloudinary) EventService {
	return &EventServiceImpl{
		DB:              db,
		Cld:             cld,
		Validate:        validate,
		EventRepository: eventRepository,
	}
}

func (s *EventServiceImpl) CreateEvent(ctx context.Context, data web.CreateEventRequest) (web.Event, error) {
	now := time.Now()
	res := web.Event{}
	aceptedFiles := []string{"JPG", "JPEG", "PNG", "WEBP", "HEIF"}
	bannerExt := strings.ToUpper(strings.TrimLeft(filepath.Ext(data.Banner.Filename), "."))
	certificateExt := strings.ToUpper(strings.TrimLeft(filepath.Ext(data.Certificate.Filename), "."))

	err := s.Validate.Struct(data)
	if err != nil {
		return res, err
	}

	// * checking file type and file size
	if !slices.Contains(aceptedFiles, certificateExt) || !slices.Contains(aceptedFiles, bannerExt) {
		return res, e.Wrapf(exception.ErrBadRequest, "Format file is not suported: file must be %s format", strings.Join(aceptedFiles, ", "))
	}
	if data.Certificate.Size > 5242880 || data.Banner.Size > 5242880 {
		return res, e.Wrap(exception.ErrBadRequest, "File is to large: maximum file size is 5 mb")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return res, err
	}
	defer helper.CommitOrRollback(tx)

	banner, err := data.Banner.Open()
	if err != nil {
		return res, err
	}
	bannerPath, err := helper.UploadFileToFirebaseStorageAndGetURL(ctx, banner)
	if err != nil {
		return res, err
	}
	certificate, err := data.Certificate.Open()
	if err != nil {
		return res, err
	}
	certificatePath, err := helper.UploadFileToFirebaseStorageAndGetURL(ctx, certificate)
	if err != nil {
		return res, err
	}

	event := domain.Events{
		Title:           data.Title,
		Description:     data.Description,
		Banner:          bannerPath,
		Certificate:     certificatePath,
		Price:           data.Price,
		Type:            data.Type,
		BankAccountNum:  data.BankAccountNum,
		BackAccountName: data.BankAccountName,
		RecipientName:   data.RecipientName,
		Location:        data.Location,
		StartAt:         data.StartAt,
		EndAt:           data.EndAt,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	id, err := s.EventRepository.SaveEvent(ctx, tx, event)
	if err != nil {
		return res, err
	}

	res = web.Event{
		Id:              id,
		Title:           data.Title,
		Description:     data.Description,
		Banner:          bannerPath,
		Certificate:     certificatePath,
		Price:           data.Price,
		Type:            data.Type,
		BankAccountNum:  &data.BankAccountNum,
		BankAccountName: &data.BankAccountName,
		RecipientName:   &data.RecipientName,
		Location:        data.Location,
		StartAt:         data.StartAt,
		EndAt:           data.EndAt,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	return res, nil
}

func (s *EventServiceImpl) GetAllEvents(ctx context.Context) ([]web.Event, error) {
	var res []web.Event
	events, err := s.EventRepository.GetAllEvents(ctx, s.DB)
	if err != nil {
		return res, err
	}

	for _, event := range events {
		res = append(res, web.Event{
			Id:             event.Id,
			Title:          event.Title,
			Description:    event.Description,
			Banner:         event.Banner,
			Certificate:    event.Certificate,
			Price:          event.Price,
			Type:           event.Type,
			BankAccountNum: &event.BankAccountNum,
			Location:       event.Location,
			StartAt:        event.StartAt,
			EndAt:          event.EndAt,
			CreatedAt:      event.CreatedAt,
			UpdatedAt:      event.UpdatedAt,
		})
	}
	return res, nil
}

func (s *EventServiceImpl) GetEventById(ctx context.Context, id string) (web.Event, error) {
	var res web.Event
	event, err := s.EventRepository.GetEventById(ctx, s.DB, id)
	if err != nil {
		return res, e.Wrap(exception.ErrNotFound, err.Error())
	}
	res = web.Event{
		Id:             event.Id,
		Title:          event.Title,
		Description:    event.Description,
		Banner:         event.Banner,
		Certificate:    event.Certificate,
		Price:          event.Price,
		Type:           event.Type,
		BankAccountNum: &event.BankAccountNum,
		Location:       event.Location,
		StartAt:        event.StartAt,
		EndAt:          event.EndAt,
		CreatedAt:      event.CreatedAt,
		UpdatedAt:      event.UpdatedAt,
	}
	return res, nil
}
