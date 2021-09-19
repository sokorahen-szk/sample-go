package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo) *echo.Echo {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e
}
