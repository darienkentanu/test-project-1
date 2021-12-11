package route

import (
	"database/sql"
	"test-project-1/middlewares"

	"github.com/labstack/echo/v4"
)

func NewRoutes(db *sql.DB) *echo.Echo {
	e := echo.New()
	middlewares.LogMiddlewares(e)
	userRoute(e, db)
	itemRoute(e, db)
	transactionRoute(e, db)
	return e
}
