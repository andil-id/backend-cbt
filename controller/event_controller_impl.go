package controller

import (
	"net/http"

	"github.com/andil-id/api/helper"
	"github.com/andil-id/api/model/web"
	"github.com/andil-id/api/service"
	"github.com/gin-gonic/gin"
)

type EventControllerImpl struct {
	EventService service.EventService
}

func NewEventController(eventService service.EventService) EventController {
	return &EventControllerImpl{
		EventService: eventService,
	}
}

func (cl EventControllerImpl) AddEvent(c *gin.Context) {
	req := web.CreateEventRequest{}
	err := c.Bind(&req)
	if err != nil {
		c.Error(err)
		return
	}
	res, err := cl.EventService.CreateEvent(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}
	meta := helper.Meta{
		StatusCode: http.StatusOK,
		Message:    "Data was succesfully transmited!",
	}
	helper.ResponseSuccess(c, res, meta)
}
