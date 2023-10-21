package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Web() *echo.Echo {
	e := echo.New()

	content := "<h1>Welcome</h1>"

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, content)
	})

	return e
}
