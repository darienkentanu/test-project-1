package controller

import (
	"net/http"
	"test-project-1/lib/database"
	"test-project-1/model"
	"time"

	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

type UserController struct {
	UserModel database.UserModel
}

func NewUserController(UserModel database.UserModel) *UserController {
	return &UserController{UserModel: UserModel}
}

func (uc *UserController) RegisterUser(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)
	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	u, err := uc.UserModel.AddUser(user)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, "error inserting to database")
	}
	return c.JSON(http.StatusAccepted, M{
		"status": "success",
		"data":   u,
	})
}
