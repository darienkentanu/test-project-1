package route

import (
	"database/sql"
	"test-project-1/controller"
	"test-project-1/lib/database"
	"test-project-1/middlewares"

	"github.com/labstack/echo/v4"
)

func transactionRoute(e *echo.Echo, db *sql.DB) {
	tm := database.NewTransactionModel(db)
	im := database.NewItemModel(db)
	tc := controller.NewTransactionController(tm, im)
	e.POST("/newtransaction", tc.NewTransaction, middlewares.IsLoggedIn)
}
