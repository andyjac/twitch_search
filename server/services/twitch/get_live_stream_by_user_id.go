package twitch

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"streaming_app/server/services/api"
)

func GetLiveStreamByUserId(id string) (Stream, api.WebError) {
	var response TwitchStreamResponse
	var stream Stream

	if id == "" {
		return stream, api.WebError{
			Status:  http.StatusBadRequest,
			Message: "You must provide a user id",
			Error:   errors.New(api.BadRequest),
		}
	}

	res, err := Get("/streams/"+id, QueryParams{})

	if err.Error != nil {
		return stream, err
	}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Println(err.Error())
		return stream, api.ErrInternalServerError
	}

	if response.Error != "" {
		return stream, api.WebError{
			Status:  response.Status,
			Message: response.Message,
			Error:   errors.New(response.Error),
		}
	}
	if response.Stream.Id == 0 {
		return stream, api.WebError{
			Status:  http.StatusNotFound,
			Message: "Live stream not found",
			Error:   errors.New(api.NotFound),
		}
	}

	return response.Stream, api.WebError{}
}
