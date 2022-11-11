package controller

import "github.com/gin-gonic/gin"

type EventController interface {
	AddEvent(c *gin.Context)
}
