package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitSQL() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "KlikA2C:KlikA2C123!@tcp(127.0.0.1:3306)/KlikA2C?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
