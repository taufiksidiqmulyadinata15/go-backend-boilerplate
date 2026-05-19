package response

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/pkg/apperror"
)

type APIResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func Success(c echo.Context, status int, message string, data any) error {
	return c.JSON(status, APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func Error(c echo.Context, status int, message string, err any) error {
	return c.JSON(status, APIResponse{
		Status:  status,
		Message: message,
		Error:   err,
	})
}

func OK(c echo.Context, message string, data any) error {
	return Success(c, http.StatusOK, message, data)
}

func Created(c echo.Context, message string, data any) error {
	return Success(c, http.StatusCreated, message, data)
}

func BadRequest(c echo.Context, message string, err any) error {
	return Error(c, http.StatusBadRequest, message, err)
}

func NotFound(c echo.Context, message string, err any) error {
	return Error(c, http.StatusNotFound, message, err)
}

func InternalServerError(c echo.Context, message string, err any) error {
	return Error(c, http.StatusInternalServerError, message, err)
}
func FromError(c echo.Context, err error) error {
	var appErr *apperror.AppError

	if errors.As(err, &appErr) {
		return Error(c, appErr.Status, appErr.Message, nil)
	}

	return InternalServerError(c, "Internal server error", err.Error())
}
