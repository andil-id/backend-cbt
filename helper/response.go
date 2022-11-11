package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	TotalPage  int `json:"total_page,omitempty"`
	TotalData  int `json:"total_data,omitempty"`
	StatusCode int `json:"status_code,omitempty"`
	Message    any `json:"message,omitempty"`
}

func ResponseSuccess(c *gin.Context, data any, meta Meta) {
	metaData := Meta{
		StatusCode: meta.StatusCode,
	}

	switch c.Request.Method {
	case "GET":
		metaData = Meta{
			Page:      meta.Page,
			PerPage:   meta.PerPage,
			TotalPage: meta.TotalPage,
			TotalData: meta.TotalData,
			Message:   "Data was successfully retrieved!",
		}

	case "DELETE":
		metaData = Meta{
			Message: "Data was successfully deleted!",
		}
	case "POST", "PUT", "PATCH":
		metaData = Meta{
			Message: "Data was sunccesfully transmited!",
		}
	default:
		err := errors.New("htpp method not recognized")
		c.Error(err)
	}
	c.JSON(meta.StatusCode, gin.H{
		"data": data,
		"meta": metaData,
	})

}

func ResponseError(c *gin.Context, code int, message any) {
	meta := Meta{
		StatusCode: code,
		Message:    message,
	}
	c.JSON(code, gin.H{
		"data": nil,
		"meta": meta,
	})
	c.Abort()
}
