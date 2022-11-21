package controller

import (
	"fmt"
	"net/http"

	"github.com/andil-id/api/helper"
	"github.com/andil-id/api/model/web"
	"github.com/andil-id/api/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (cl *OrderControllerImpl) CreateOrderEvent(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	id := token["id"]
	req := web.CreateOrderRequest{}
	err := c.Bind(&req)
	if err != nil {
		c.Error(err)
		return
	}
	res, err := cl.OrderService.CreateOrder(c.Request.Context(), req, fmt.Sprintf("%v", id))
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, res, helper.Meta{
		StatusCode: http.StatusOK,
	})
}
