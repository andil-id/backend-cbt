package controller

import (
	"net/http"

	"github.com/andil-id/api/exception"
	"github.com/andil-id/api/helper"
	"github.com/andil-id/api/model/web"
	"github.com/andil-id/api/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	e "github.com/pkg/errors"
)

type EventControllerImpl struct {
	EventService service.EventService
}

func NewEventController(eventService service.EventService) EventController {
	return &EventControllerImpl{
		EventService: eventService,
	}
}

func (cl *EventControllerImpl) AddEvent(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	if token["role"].(string) != "admin" {
		c.Error(e.Wrap(exception.ErrUnAuth, ""))
		return
	}
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
	helper.ResponseSuccess(c, res, helper.Meta{
		StatusCode: http.StatusOK,
	})
}

func (cl *EventControllerImpl) GetAllEvents(c *gin.Context) {
	res, err := cl.EventService.GetAllEvents(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, res, helper.Meta{
		StatusCode: http.StatusOK,
		TotalData:  len(res),
	})
}

func (cl *EventControllerImpl) GetEventById(c *gin.Context) {
	id := c.Param("id")
	res, err := cl.EventService.GetEventById(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, res, helper.Meta{
		StatusCode: http.StatusOK,
	})
}
