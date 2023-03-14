package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func CheckEngine() {
	connStr := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("The database is connected")
}

func OpenConnection() *sql.DB {
	connStr := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return db
}
