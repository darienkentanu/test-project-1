package helper

import (
	"database/sql"
	"os"
	"test-project-1/config"
)

func InitDB() (db *sql.DB, err error) {
	if os.Getenv("CONN_STRING") == "" {
		db, err = config.InitSQL()
	} else {
		db, err = config.InitSQLDocker()
	}
	return db, err
}
