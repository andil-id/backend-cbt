package controller

import "github.com/gin-gonic/gin"

type EventController interface {
	AddEvent(c *gin.Context)
	GetAllEvents(c *gin.Context)
	GetEventById(c *gin.Context)
}
