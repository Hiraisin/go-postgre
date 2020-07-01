package main

import (
	"github.com/hiraisin/go-postgre/config"
	"github.com/hiraisin/go-postgre/routes"
)

func main() {
	conf := config.GetConfig()
	config.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(conf.APP_PORT))
}
