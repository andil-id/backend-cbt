package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/andil-id/api/exception"
	"github.com/andil-id/api/helper"
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
		if serviceError(c, err) {
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
		helper.ResponseError(c, http.StatusBadRequest, errorMessages)
		return true
	} else {
		return false
	}
}

func notFoundError(c *gin.Context, err error) bool {
	if e.Cause(err) == exception.ErrNotFound {
		errorMessage := strings.Split(err.Error(), ":")
		helper.ResponseError(c, http.StatusNotFound, errorMessage[0])
		return true
	} else {
		return false
	}
}

func badRequestError(c *gin.Context, err error) bool {
	if e.Cause(err) == exception.ErrBadRequest {
		errorMessage := strings.Split(err.Error(), ":")
		helper.ResponseError(c, http.StatusBadRequest, errorMessage[0])
		return true
	} else {
		return false
	}
}

func serviceError(c *gin.Context, err error) bool {
	if e.Cause(err) == exception.ErrService {
		errorMessage := strings.Split(err.Error(), ":")
		helper.ResponseError(c, http.StatusBadRequest, errorMessage[0])
		return true
	} else {
		return false
	}
}

func internalServerError(c *gin.Context, err error) {
	helper.ResponseError(c, http.StatusBadRequest, "INTERNAL SERVER ERROR")
}
