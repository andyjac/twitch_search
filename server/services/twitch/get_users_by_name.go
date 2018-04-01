package twitch

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"streaming_app/server/services/api"
)

func GetUsersByName(name string) (Users, api.WebError) {
	var response TwitchUserResponse
	var users Users

	if name == "" {
		return users, api.WebError{
			Status:  http.StatusBadRequest,
			Message: "You must provide at least one username",
			Error:   errors.New(api.BadRequest),
		}
	}

	res, err := Get("/users", QueryParams{"login": name})

	if err.Error != nil {
		return users, err
	}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Println(err.Error())
		return users, api.ErrInternalServerError
	}

	// if we got an error from twitch...
	if response.Error != "" {
		return users, api.WebError{
			Status:  response.Status,
			Message: response.Message,
			Error:   errors.New(response.Error),
		}
	}
	// or we did not find any users
	if response.Total == 0 {
		return users, api.WebError{
			Status:  http.StatusNotFound,
			Message: "Streamer not found: " + name,
			Error:   errors.New(api.NotFound),
		}
	}

	return response.Users, api.WebError{}
}
