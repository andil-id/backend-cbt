package exception

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	e "github.com/pkg/errors"
)

func ErrorAppHandler() gin.HandlerFunc {
	return jsonErrorReporter(gin.ErrorTypeAny)
}

func jsonErrorReporter(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectError := c.Errors.ByType(errType)
		if len(detectError) == 0 {
			return
		}
		err := detectError[0].Err

		if validationErrors(c, err) {
			return
		}
		if notFoundError(c, err) {
			return
		}
		if badRequestError(c, err) {
			return
		}
		internalServerError(c, err)
	}
}

func validationErrors(c *gin.Context, err error) bool {
	_, errorValidation := err.(validator.ValidationErrors)
	if errorValidation {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error in field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": errorMessages,
			"data":    nil,
		})
		c.Abort()
		return true
	} else {
		return false
	}
}

func notFoundError(c *gin.Context, err error) bool {
	if e.Cause(err) == ErrNotFound {
		errorMessage := strings.Split(err.Error(), ":")
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": errorMessage[0],
			"data":    nil,
		})
		c.Abort()
		return true
	} else {
		return false
	}
}

func badRequestError(c *gin.Context, err error) bool {
	if e.Cause(err) == ErrBadRequest {
		errorMessage := strings.Split(err.Error(), ":")
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": errorMessage[0],
			"data":    nil,
		})
		c.Abort()
		return true
	} else {
		return false
	}
}

func internalServerError(c *gin.Context, err error) {
	c.IndentedJSON(500, gin.H{
		"code":          500,
		"error_message": err.Error(),
		"message":       "INTERNAL SERVER ERROR",
		"data":          nil,
	})
	c.Abort()
}
