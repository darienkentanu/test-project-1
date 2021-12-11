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
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expires after 1 hour

	claims["id"] = id

	tokenString, err := token.SignedString([]byte(constants.JWT_SECRET))
	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

// func CurrentLoginID(c echo.Context) int {
// 	token := c.Get("user").(*jwt.Token)
// 	if token != nil && token.Valid {
// 		claims := token.Claims.(jwt.MapClaims)
// 		id := claims["id"]
// 		switch id.(type) {
// 		case float64:
// 			return int(id.(float64))
// 		default:
// 			return id.(int)
// 		}
// 	}
// 	return -1 // invalid user
// }
