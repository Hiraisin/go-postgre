package handlers

import (
	"net/http"

	"github.com/hiraisin/go-postgre/config"
	"github.com/hiraisin/go-postgre/models"
	"github.com/labstack/echo"
)

func GetAll(c echo.Context) error {
	result, err := models.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func Input(c echo.Context) error {
	data := config.M{
		"title":   "Form Input",
		"message": "Hallo Rizky",
	}
	return c.Render(http.StatusOK, "index.html", data)
}

func StoreData(c echo.Context) error {
	Nama := c.FormValue("nama")
	Alamat := c.FormValue("alamat")
	Telepon := c.FormValue("telepon")
	if Nama == "" || Alamat == "" || Telepon == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Inputan tidak boleh kosong",
		})
	}

	result, err := models.StorePegawai(Nama, Alamat, Telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
