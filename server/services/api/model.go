package api

import (
	"errors"
	"net/http"
)

type WebError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

var (
	ServerError = "Internal Server Error"
	BadRequest  = "Bad Request"
	NotFound    = "Not Found"
)

var ErrInternalServerError = WebError{
	Status:  http.StatusInternalServerError,
	Message: "Uh oh! Looks like something went wrong with your request.",
	Error:   errors.New(ServerError),
}
