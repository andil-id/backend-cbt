package controller

import (
	"fmt"
	"net/http"

	"github.com/andil-id/api/exception"
	"github.com/andil-id/api/helper"
	"github.com/andil-id/api/model/web"
	"github.com/andil-id/api/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	e "github.com/pkg/errors"
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

func (cl *OrderControllerImpl) ConfirmOrder(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	id := c.Param("id")
	if token["role"].(string) != "admin" {
		c.Error(e.Wrap(exception.ErrUnAuth, ""))
		return
	}
	err := cl.OrderService.UpdateOrderStatus(c.Request.Context(), "confirm", id)
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, nil, helper.Meta{
		StatusCode: http.StatusOK,
	})
}

func (cl *OrderControllerImpl) RejectOrder(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	if token["role"].(string) != "admin" {
		c.Error(e.Wrap(exception.ErrUnAuth, ""))
		return
	}

	id := c.Param("id")
	err := cl.OrderService.UpdateOrderStatus(c.Request.Context(), "reject", id)
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, nil, helper.Meta{
		StatusCode: http.StatusOK,
	})
}

func (cl *OrderControllerImpl) GetOrderById(c *gin.Context) {
	id := c.Param("id")
	res, err := cl.OrderService.GetOrderById(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, res, helper.Meta{
		StatusCode: http.StatusOK,
	})
}

func (cl *OrderControllerImpl) GetOrderByEventId(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	if token["role"].(string) != "admin" {
		c.Error(e.Wrap(exception.ErrUnAuth, ""))
		return
	}

	id := c.Param("id")
	res, err := cl.OrderService.GetOrderByEventId(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, res, helper.Meta{
		TotalData:  len(res),
		StatusCode: http.StatusOK,
	})
}

func (cl *OrderControllerImpl) GetOrderEventByUser(c *gin.Context) {
	token := c.MustGet("token").(jwt.MapClaims)
	id := token["id"]
	res, err := cl.OrderService.GetOrderEventByUserId(c.Request.Context(), fmt.Sprintf("%v", id))
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, res, helper.Meta{
		TotalData:  len(res),
		StatusCode: http.StatusOK,
	})
}
