package database

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

func CreateTable() {
	db := OpenConnection()
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	query := `CREATE TABLE IF NOT EXISTS tasks (
	   			uuid uuid NOT NULL,
	   			status varchar(50) NOT NULL DEFAULT 'CREATED',
	   			created_at timestamp NOT NULL,
	   			started_at timestamp,
	   			finished_at timestamp,
	   			PRIMARY KEY (uuid)
	)`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func InsertTask(uuid uuid.UUID) error {
	db := OpenConnection()
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	query := `INSERT INTO "tasks"("uuid", "created_at") values($1, $2)`
	_, err := db.Exec(query, uuid, time.Now())
	return err
}

func UpdateTaskByUuid(uuid uuid.UUID, status string) error {
	var err error
	db := OpenConnection()
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	switch status {
	case "PROCESSING":
		query := `UPDATE "tasks" SET status = $1, started_at = $2
               		WHERE uuid = $3`
		_, err = db.Exec(query, status, time.Now(), uuid)
	case "ERROR_IN_DOWNLOAD":
		query := `UPDATE "tasks" SET status = $1, finished_at = $2
               		WHERE uuid = $3`
		_, err = db.Exec(query, status, time.Now(), uuid)
	case "ERROR_IN_PROCESS":
		query := `UPDATE "tasks" SET status = $1, finished_at = $2
               		WHERE uuid = $3`
		_, err = db.Exec(query, status, time.Now(), uuid)
	case "SUCCESS":
		query := `UPDATE "tasks" SET status = $1, finished_at = $2
               		WHERE uuid = $3`
		_, err = db.Exec(query, status, time.Now(), uuid)
	default:
	}

	return err

}
