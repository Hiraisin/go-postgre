package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hiraisin/go-postgre/config"
	"github.com/hiraisin/go-postgre/models"
	"github.com/labstack/echo"
)

func FormLogin(c echo.Context) error {
	data := config.M{
		"title": "Form Login",
	}
	return c.Render(http.StatusOK, "login.html", data)
}

func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Inputan tidak boleh kosong",
		})
	}
	err := models.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": t,
	})
}
