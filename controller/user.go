package controller

import (
	"net/http"
	"test-project-1/helper"
	"test-project-1/lib/database"
	"test-project-1/middlewares"
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
	hash, err := helper.GenerateHashPassword(user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error hashing the password")
	}
	user.Password = hash
	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	u, err := uc.UserModel.AddUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error inserting to database")
	}
	return c.JSON(http.StatusAccepted, M{
		"status": "success",
		"data":   u,
	})
}

func (uc *UserController) LoginUser(c echo.Context) error {
	login := model.User_Login{}
	c.Bind(&login)
	account, err := uc.UserModel.GetAccountByEmailNullToken(login)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error getting user data")
	}
	pwdCorrect := helper.CheckPasswordHash(login.Password, account.Password)
	if !pwdCorrect {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid password")
	}
	token, err := middlewares.CreateToken(account.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating the token")
	}
	err = uc.UserModel.UpdateToken(token, account.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating the token")
	}
	user, err := uc.UserModel.GetAccountByEmailWithToken(login)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error getting user data", err.Error())
	}
	return c.JSON(http.StatusOK, M{
		"status": "login success",
		"data":   user,
	})
}
