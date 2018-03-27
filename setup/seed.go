package main

import (
	"log"
	"strconv"
	"streaming_app/server/database"
)

var NUM_USERS = 100

func main() {
	log.Println("Seeding database...")

	var isStreaming bool
	var name string

	db := database.Connect()

	for i := 1; i < NUM_USERS; i++ {
		name = "streamer_" + strconv.Itoa(i)

		if i%5 == 0 {
			isStreaming = false
		} else {
			isStreaming = true
		}

		stmt, err := db.Prepare("insert into streamer(name, is_streaming) values (?, ?)")

		if err != nil {
			log.Println(err.Error())
		}

		_, err = stmt.Exec(name, isStreaming)

		if err != nil {
			log.Println(err.Error())
		}
	}

	log.Println("Database seeded.")
}
