package model

import "net/http"

const (
	StatusOK = http.StatusOK // 200
	StatusCreated = http.StatusCreated // 201
	StatusNoContent = http.StatusNoContent // 204
	StatusBadRequest = http.StatusBadRequest   // 400
	StatusUnauthorized = http.StatusUnauthorized // 401
	StatusForbidden = http.StatusForbidden // 403
	StatusMethodNotAllowed = http.StatusMethodNotAllowed // 405
	StatusNotAcceptable = http.StatusNotAcceptable // 406
	StatusNotFound   = http.StatusNotFound     // 404
	StatusConflict = http.StatusConflict // 409
	StatusTooManyRequests = http.StatusTooManyRequests // 429
	StatusRequestTimeout = http.StatusRequestTimeout // 408
	StatusUnsupportedMediaType = http.StatusUnsupportedMediaType // 415
	StatusUnprocessableEntity = http.StatusUnprocessableEntity // 422
	StatusInternalServerError = http.StatusInternalServerError // 500
	StatusServiceUnavailable = http.StatusServiceUnavailable // 503
	StatusBadGateway = http.StatusBadGateway // 502
)

// StatusCodeToText converts HTTP status code to human readable text
func StatusCodeToText(code int) string {
	switch code {
	case StatusOK:
		return "OK"
	case StatusCreated:
		return "Created"
	case StatusNoContent:
		return "No Content"
	case StatusBadRequest:
		return "validation"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusForbidden:
		return "Forbidden"
	case StatusMethodNotAllowed:
		return "Method Not Allowed"
	case StatusNotAcceptable:
		return "Not Acceptable"
	case StatusNotFound:
		return "Not Found"
	case StatusConflict:
		return "Conflict"
	case StatusTooManyRequests:
		return "Too Many Requests"
	case StatusRequestTimeout:
		return "Request Timeout"
	case StatusUnsupportedMediaType:
		return "Unsupported Media Type"
	case StatusInternalServerError:
		return "Internal Server Error"
	case StatusServiceUnavailable:
		return "Service Unavailable"
	case StatusBadGateway:
		return "Bad Gateway"
	case StatusUnprocessableEntity:
		return "constraint"
	default:
		return "Unknown Status"
	}
}


type Response[T any] struct {
	Data T `json:"data,omitempty"`
	Message string `json:"message"`
	Status int `json:"-"`
	Code string `json:"code"`
}

func (r *Response[T]) Ok() bool {
	return r.Status == StatusOK
}

func (r *Response[T]) Error() bool {
	return r.Status != StatusOK
}

func (r *Response[T]) StatusCode() int {
	return r.Status
}

func (r *Response[T]) StatusText() string {
	return StatusCodeToText(r.Status)
}

func NotFoundResponse[T any]() Response[T] {
	var zero T
	return Response[T]{
		Data: zero,
		Message: "Not Found",
		Status: StatusNotFound,
	}
}

func SuccessResponse[T any](data T) Response[T] {
	return Response[T]{
		Data: data,
		Message: "Success",
		Status: StatusOK,
		Code: "success",
	}
}

func ErrorResponse[T any](response Response[T]) Response[T] {
	return Response[T]{
		Data: response.Data,
		Message: response.Message,
		Status: response.Status,
		Code: StatusCodeToText(response.Status),
	}
}
