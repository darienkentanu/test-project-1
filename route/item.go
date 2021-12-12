package route

import (
	"database/sql"
	"test-project-1/controller"
	"test-project-1/lib/database"
	"test-project-1/middlewares"

	"github.com/labstack/echo/v4"
)

func itemRoute(e *echo.Echo, db *sql.DB) {
	im := database.NewItemModel(db)
	ic := controller.NewItemController(im)
	e.GET("/getitems", ic.GetItems, middlewares.IsLoggedIn)
	e.POST("/additem", ic.AddItem, middlewares.IsLoggedIn)
	e.DELETE("/deleteitem/:id", ic.DeleteItem, middlewares.IsLoggedIn)
	e.PUT("/edititem/:id", ic.EditItem, middlewares.IsLoggedIn)
}
