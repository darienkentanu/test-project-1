package route

import (
	"database/sql"
	"test-project-1/controller"
	"test-project-1/lib/database"

	"github.com/labstack/echo/v4"
)

func userRoute(e *echo.Echo, db *sql.DB) {
	um := database.NewUserModel(db)
	uc := controller.NewUserController(um)
	e.POST("/register", uc.RegisterUser)
	e.POST("/login", uc.LoginUser)
}
