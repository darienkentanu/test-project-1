package helper

import (
	"database/sql"
	"test-project-1/config"

	"github.com/labstack/echo/v4"
)

func InitEcho() (*echo.Echo, *sql.DB) {
	e := echo.New()
	db, err := config.InitSQLTest()
	if err != nil {
		panic(err)
	}
	return e, db
}
