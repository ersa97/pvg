package models

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Responses struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseContext(code int, message interface{}, data interface{}, c echo.Context) error {
	if code == 200 { // Success
		return SuccessContext(message, data, c)
	} else if code == 400 { // Bad Request
		return ValidationContext(message, data, c)
	} else if code == 404 { // Notfound
		return NotFoundContext(message, data, c)
	} else if code == 504 { // Timeout
		return TimeoutContext(message, c)
	}
	return ErrorContext(message, c) // Internal Server Error
}

func SuccessContext(message interface{}, data interface{}, c echo.Context) (err error) {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"StatusCode": 200,
		"Status":     "success",
		"Message":    message,
		"Data":       data,
	})
}

func ValidationContext(message interface{}, data interface{}, c echo.Context) (err error) {
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"StatusCode": 400,
		"Status":     "validation",
		"Message":    message,
		"Data":       data,
	})
}

func NotFoundContext(message interface{}, data interface{}, c echo.Context) (err error) {
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"StatusCode": 404,
		"Status":     "not found",
		"Message":    message,
		"Data":       nil,
	})
}

func TimeoutContext(message interface{}, c echo.Context) (err error) {
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"StatusCode": 504,
		"Status":     "timeout",
		"Message":    message,
		"Data":       nil,
	})
}

func ErrorContext(message interface{}, c echo.Context) (err error) {
	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"StatusCode": 500,
		"Status":     "failed",
		"Message":    message,
		"Data":       nil,
	})
}
