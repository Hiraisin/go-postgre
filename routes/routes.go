package routes

import (
	"net/http"

	"github.com/hiraisin/go-postgre/config"
	"github.com/hiraisin/go-postgre/handlers"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Renderer = config.NewRenderer("./templates/*.html", true)

	e.GET("/tes", func(c echo.Context) error {
		return c.String(http.StatusOK, "oke")
	})

	e.GET("/login", handlers.FormLogin)
	e.POST("/login", handlers.CheckLogin)

	e.GET("/input", handlers.Input)
	e.GET("/pegawai", handlers.GetAll)
	e.POST("/post", handlers.StoreData)

	return e
}
