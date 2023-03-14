package database

import (
	"database/sql"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateTable() {
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	query := `CREATE TABLE IF NOT EXISTS tasks (
    			uuid uuid NOT NULL,
    			status varchar(50) NOT NULL DEFAULT 'CREATED',
    			created_at timestamp NOT NULL,
    			started_at timestamp,
    			finished_at timestamp,
    			PRIMARY KEY (uuid)
)`
	_, e := db.Exec(query)
	//_, e := db.Exec(insertDynStmt, "Jack", 21)
	CheckError(e)
}
