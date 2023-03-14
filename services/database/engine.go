package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Engine() {
	var err error

	connStr := os.Getenv("DB_URL")

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("The database is connected")
}
