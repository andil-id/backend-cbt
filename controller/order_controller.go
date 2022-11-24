package controller

import "github.com/gin-gonic/gin"

type OrderController interface {
	CreateOrderEvent(c *gin.Context)
	ConfirmOrder(c *gin.Context)
	RejectOrder(c *gin.Context)
	GetOrderById(c *gin.Context)
	GetOrderByEventId(c *gin.Context)
	GetOrderEventByUser(c *gin.Context)
}
