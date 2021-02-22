package shared

import (
	"net/http"
)

// ErrorResponse is the response that represents an error.
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Error is required by the error interface.
func (e Response) Error() string {
	return e.Message
}

// StatusCode is required by routing.HTTPError interface.
func (e Response) StatusCode() int {
	return e.Status
}

// StatusInternalServerError creates a new error response representing an internal server error (HTTP 500)
func StatusInternalServerError(msg string) Response {
	if msg == "" {
		msg = "We encountered an error while processing your request."
	}
	return Response{
		Status:  http.StatusInternalServerError,
		Message: msg,
	}
}

// StatusNotFound creates a new error response representing a resource-not-found error (HTTP 404)
func StatusNotFound(msg string) Response {
	if msg == "" {
		msg = "The requested resource was not found."
	}
	return Response{
		Status:  http.StatusNotFound,
		Message: msg,
	}
}

// StatusUnauthorized creates a new error response representing an authentication/authorization failure (HTTP 401)
func StatusUnauthorized(msg string) Response {
	if msg == "" {
		msg = "You are not authenticated to perform the requested action."
	}
	return Response{
		Status:  http.StatusUnauthorized,
		Message: msg,
	}
}

// StatusForbidden creates a new error response representing an authorization failure (HTTP 403)
func StatusForbidden(msg string) Response {
	if msg == "" {
		msg = "You are not authorized to perform the requested action."
	}
	return Response{
		Status:  http.StatusForbidden,
		Message: msg,
	}
}

// StatusBadRequest creates a new error response representing a bad request (HTTP 400)
func StatusBadRequest(msg string) Response {
	if msg == "" {
		msg = "Your request is in a bad format."
	}
	return Response{
		Status:  http.StatusBadRequest,
		Message: msg,
	}
}

// StatusUnprocessableEntity creates a new error response representing a bad request (HTTP 400)
func StatusUnprocessableEntity(msg string) Response {
	if msg == "" {
		msg = "Unprocessable Entity"
	}
	return Response{
		Status:  http.StatusUnprocessableEntity,
		Message: msg,
	}
}

func Success(resp interface{}) Response {
	return Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    resp,
	}
}

func SuccessWithCode(resp interface{}, status int) Response {
	return Response{
		Status:  status,
		Message: "Success",
		Data:    resp,
	}
}
