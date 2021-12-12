package middlewares

import (
	"fmt"
	"test-project-1/constants"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(constants.JWT_SECRET),
})

func CreateToken(id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // Token expires after 2 hour

	claims["id"] = id

	tokenString, err := token.SignedString([]byte(constants.JWT_SECRET))
	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
