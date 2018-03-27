package database

import (
	"database/sql"
	"log"
	"streaming_app/server/config"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	Connect()
}

func Connect() *sql.DB {
	config := config.Read()

	host := config.Database.Host
	port := config.Database.Port
	user := config.Database.User
	password := config.Database.Password
	typ := config.Database.Type

	db, err := sql.Open(typ, user+":"+password+"@tcp("+host+":"+port+")/go_react_native_dev")

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
