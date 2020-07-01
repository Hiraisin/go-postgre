package routes

import (
	"net/http"

	"github.com/hiraisin/go-postgre/config"
	"github.com/hiraisin/go-postgre/handlers"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Renderer = config.NewRenderer("./*.html", true)
	// e.Renderer = handlers.NewRenderer("./*.html", true)

	e.GET("/tes", func(c echo.Context) error {
		return c.String(http.StatusOK, "oke")
	})

	e.GET("/input", handlers.Input)
	e.GET("/pegawai", handlers.GetAll)
	e.POST("/post", handlers.StoreData)

	return e
}
