package dto

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func BadRequest(format string, arg ...interface{}) error {
	msg := fmt.Sprintf(format, arg...)
	return echo.NewHTTPError(http.StatusBadRequest, msg)
}

func InternalError(err error) error {
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) *Response {
	return &Response{
		Message: "success",
		Data:    data,
	}
}
