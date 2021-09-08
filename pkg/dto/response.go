package dto

import "net/http"

type CommonResponse interface {
	Code() int
	Response() Response
}

type Response struct {
	Message string `json:"message"`
	Data    interface{}
}

type BadRequestResponse string

func (b BadRequestResponse) Code() int {
	return http.StatusBadRequest
}

func (b BadRequestResponse) Response() Response {
	return Response{
		Message: string(b),
		Data:    nil,
	}
}

func BadRequest(msg string) CommonResponse {
	return BadRequestResponse(msg)
}

type InternalErrorResponse struct {
	error error
}

func (InternalErrorResponse) Code() int {
	return http.StatusInternalServerError
}
func (r InternalErrorResponse) Response() Response {
	return Response{
		Message: r.error.Error(),
	}
}

func InternalError(err error) CommonResponse {
	return InternalErrorResponse{error: err}
}

type SuccessResponse struct {
	data interface{}
}

func (s SuccessResponse) Code() int {
	return http.StatusOK
}

func (s SuccessResponse) Response() Response {
	return Response{
		Message: "success",
		Data:    s.data,
	}
}

func Success(data interface{}) CommonResponse {
	return SuccessResponse{
		data: data,
	}
}
