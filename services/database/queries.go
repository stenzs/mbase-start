package database

import (
	"time"

	"github.com/google/uuid"
)

func CreateTable() {
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
	query := `INSERT INTO "tasks"("uuid", "created_at") values($1, $2)`
	_, err := db.Exec(query, uuid, time.Now())
	return err
}
