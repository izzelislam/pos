package main

import (
	"pos/config"
	"pos/routes"

	"github.com/labstack/echo/v4"
)

func main()  {
	config.InitDb()
	e := echo.New()
	e = 	routes.Route(e)

	e.Logger.Fatal(e.Start(":8000"))
}