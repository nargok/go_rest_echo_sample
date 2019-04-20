package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type Comment struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}

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

	e.POST("/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, "Hello "+name)
	})

	e.GET("/comments", getComments)

	e.Logger.Fatal(e.Start(":1323"))
}

func getComments(c echo.Context) error {
	comments := []Comment{Comment{1, "Test comment", "This is Test"}}
	comment2 := Comment{Id: 2, Name: "Test comment2", Text: "This is Test2"}
	comments = append(comments, comment2)
	return c.JSON(http.StatusOK, comments)
}
