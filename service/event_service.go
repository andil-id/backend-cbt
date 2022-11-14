package service

import (
	"context"

	"github.com/andil-id/api/model/web"
)

type EventService interface {
	GetAllEvents(ctx context.Context) ([]web.Event, error)
	GetEventById(ctx context.Context, id string) (web.Event, error)
	CreateEvent(ctx context.Context, data web.CreateEventRequest) (web.Event, error)
}
