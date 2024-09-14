package infra

import (
	"database/sql"
	"log"
)

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/users")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
