package http_err

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	StatusText string `json:"status_text"`
	Message    string `json:"message"`
}

var (
	ErrStatusOK         = &ErrorResponse{StatusCode: http.StatusOK, Message: "Status OK"}
	ErrBadRequest       = &ErrorResponse{StatusCode: http.StatusBadRequest, Message: "Bad request"}
	ErrUnauthorized     = &ErrorResponse{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"}
	ErrForbidden        = &ErrorResponse{StatusCode: http.StatusForbidden, Message: "Access forbidden"}
	ErrNotFound         = &ErrorResponse{StatusCode: http.StatusNotFound, Message: "Resource not found"}
	ErrMethodNotAllowed = &ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: "Method not allowed"}
	ErrInternalServer   = &ErrorResponse{StatusCode: http.StatusInternalServerError, Message: "Internal Server Error"}
	ErrStrAuthIncorrect = "Auth data incorrect"
)

func ErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: http.StatusBadRequest,
		StatusText: "Bad request",
		Message:    err.Error(),
	}
}

func ServerErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: http.StatusInternalServerError,
		StatusText: "Internal server error",
		Message:    err.Error(),
	}
}

type errorResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
