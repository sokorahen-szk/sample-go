package main

import (
	router "sokorahen-szk/sample-go/web-router"

	"github.com/labstack/echo/v4"
)

func main() {
	app := router.NewRouter(echo.New())
	app.Logger.Fatal(app.Start("localhost:1323"))
}
