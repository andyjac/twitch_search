package twitch

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"streaming_app/server/services/api"
)

func GetUserByLogin(name string) (User, api.WebError) {
	var response TwitchUserResponse
	var user User

	if name == "" {
		return user, api.WebError{
			Status:  http.StatusBadRequest,
			Message: "You must provide a user login name",
			Error:   errors.New(api.BadRequest),
		}
	}

	res, err := Get("/users", QueryParams{"login": name})

	if err.Error != nil {
		return user, err
	}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Println(err.Error())
		return user, api.ErrInternalServerError
	}

	// if we got an error from twitch...
	if response.Error != "" {
		return user, api.WebError{
			Status:  response.Status,
			Message: response.Message,
			Error:   errors.New(response.Error),
		}
	}
	// or we did not find a user
	if response.Total == 0 {
		return user, api.WebError{
			Status:  http.StatusNotFound,
			Message: "Streamer not found: " + name,
			Error:   errors.New(api.NotFound),
		}
	}

	user = response.Users[0]
	return user, api.WebError{}
}
