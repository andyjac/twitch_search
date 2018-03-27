package models

import (
	"log"
	"streaming_app/server/database"
)

type Streamer struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Is_Streaming bool   `json:"isStreaming"`
}

type Streamers []Streamer

func GetStreamers() (Streamers, HTTPError) {
	var streamers Streamers
	var streamer Streamer
	var error HTTPError

	db := database.Connect()
	rows, err := db.Query("select id, name, is_streaming from streamer;")

	if err != nil {
		log.Println(err.Error())
		return streamers, ErrInternalServerError
	}

	for rows.Next() {
		err = rows.Scan(&streamer.Id, &streamer.Name, &streamer.Is_Streaming)
		streamers = append(streamers, streamer)

		if err != nil {
			log.Println(err.Error())
			return streamers, ErrInternalServerError
		}
	}

	defer rows.Close()
	defer db.Close()

	return streamers, error
}

func GetStreamerById(id string) (Streamer, HTTPError) {
	var streamer Streamer
	var error HTTPError

	db := database.Connect()
	row := db.QueryRow("select id, name, is_streaming from streamer where id = ?", id)
	err := row.Scan(&streamer.Id, &streamer.Name, &streamer.Is_Streaming)

	if err != nil {
		log.Println(err.Error())
		return streamer, ErrStreamerNotFound
	}

	defer db.Close()
	return streamer, error
}
