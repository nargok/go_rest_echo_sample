package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.String(http.StatusOK, "Hello"+name)
	})

	e.POST("/", func(c echo.Context) error {
		name := c.FormValue("name")
		return c.String(http.StatusOK, "Hello "+name)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
