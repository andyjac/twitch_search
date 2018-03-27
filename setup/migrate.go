package main

import (
	"log"
	"streaming_app/server/database"
)

func main() {
	db := database.Connect()

	log.Println("Migrating database...")
	log.Println("Dropping table: streamer...")

	stmt, err := db.Prepare("DROP TABLE IF EXISTS streamer;")

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Creating table: streamer...")

	createStreamerTable := `
CREATE TABLE streamer (
id int NOT NULL AUTO_INCREMENT,
name varchar(40) NOT NULL,
is_streaming boolean DEFAULT false,
PRIMARY KEY (id)
);`

	stmt, err = db.Prepare(createStreamerTable)

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Database migrated.")
}
