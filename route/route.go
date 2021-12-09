package route

import (
	"database/sql"
	"test-project-1/controller"
	"test-project-1/lib/database"

	"github.com/labstack/echo/v4"
)

func NewRoutes(db *sql.DB) *echo.Echo {
	e := echo.New()
	um := database.NewUserModel(db)
	uc := controller.NewUserController(um)
	e.POST("/register", uc.RegisterUser)
	return e
}
