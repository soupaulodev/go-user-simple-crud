package infra

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB_DRIVER = "postgres"
var DB_URI = ""

func InitDB() *sql.DB {
	db, err := sql.Open(DB_DRIVER, DB_URI)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
