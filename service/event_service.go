package service

import (
	"context"

	"github.com/andil-id/api/model/web"
)

type EventService interface {
	CreateEvent(ctx context.Context, data web.CreateEventRequest) (web.Event, error)
}
