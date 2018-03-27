package models

import (
	"errors"
	"net/http"
)

type HTTPError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

var (
	streamerNotFound    = "Streamer not found."
	internalServerError = "Internal server error."
)

var (
	ErrStreamerNotFound = HTTPError{
		Status:  http.StatusNotFound,
		Message: streamerNotFound,
		Error:   errors.New(streamerNotFound),
	}
	ErrInternalServerError = HTTPError{
		Status:  http.StatusInternalServerError,
		Message: internalServerError,
		Error:   errors.New(internalServerError),
	}
)
